package bookingkamaroperasi

import (
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/jasonrichdarmawan/cipta_integrasi_nusantara/helper"
)

// URL path: /bookingkamaroperasi/dd-mm-yyyy hour:minute+offset/minutes
// example: /bookingkamaroperasi/10-10-1999 10:00+offset/60
func POST(w http.ResponseWriter, r *http.Request) {
	regPath := regexp.MustCompile(`\/bookingkamaroperasi\/([\d-+ :]+)\/(\d{1,4})`)
	match := regPath.FindStringSubmatch(r.URL.Path)
	if match == nil {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// dd-mm-yyyy hour:minute+OffSet
	bookingDate, err := time.Parse("02-01-2006 15:04Z07:00", match[1])
	if err != nil {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		helper.Print(err, r)
		return
	}
	duration, err := strconv.Atoi(match[2])
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		helper.Print(err, r)
		return
	}

	switch r.Method {
	case "POST":
		if TryAppend(bookingDate, time.Duration(duration)*time.Minute) {
			w.Write([]byte("true"))
		} else {
			w.Write([]byte("false"))
		}
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}

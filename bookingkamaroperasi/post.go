package bookingkamaroperasi

import (
	"log"
	"net/http"
	"regexp"
	"time"
)

// URL path: /bookingkamaroperasi/dd-mm-yyyy hour:minute+offset/minutes
// example: /bookingkamaroperasi/10-10-1999 10:00+offset/60
func Post(w http.ResponseWriter, r *http.Request) {
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
		log.Println(r.URL.Path + " " + r.RemoteAddr + " " + err.Error())
		return
	}
	duration := match[2]

	switch r.Method {
	case "POST":
		w.Write([]byte(bookingDate.String() + " " + duration))
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}

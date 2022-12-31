package calculateage

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/jasonrichdarmawan/cipta_integrasi_nusantara_backend/helper"
)

type umur struct {
	Year  int     `json:"year"`
	Month int     `json:"month"`
	Day   float64 `json:"day"`
}

type calculateAgeResponseBody struct {
	Umur umur `json:"umur"`
}

// Duly noted: the spec requirement is to
// update the data automatically and scheduled.
// However, the provided scenario only give example for calculating age.
// To avoid making out of spec feature, we will only calculating age.
func GET(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	regPath := regexp.MustCompile(`\/calculateage\/([\d-]+)`)
	match := regPath.FindStringSubmatch(r.URL.Path)
	if match == nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// dd-mm-yyyy
	dateOfBirth, err := time.Parse("02-01-2006Z07:00", match[1]+"+07:00")
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		helper.Print(err, r)
		return
	}

	today := helper.RemoveFromHour(time.Now())

	log.Println(dateOfBirth, today)

	yearsDiff, monthsDiff, daysDiff, err := CalculateYearsMonthsDaysBetweenDates(dateOfBirth, today)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	p := calculateAgeResponseBody{}
	p.Umur = umur{Year: yearsDiff, Month: monthsDiff, Day: daysDiff}

	responseBody, err := json.Marshal(p)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(responseBody)
}

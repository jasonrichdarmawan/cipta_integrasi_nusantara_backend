package hitunggaji

import (
	"encoding/json"
	"net/http"

	"github.com/jasonrichdarmawan/cipta_integrasi_nusantara_backend/helper"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type salaryStructure struct {
	// for simplicity, we will use float64
	// see playground/money_test.go for uint64 implementation
	Value    float64  `json:"value"`
	Country  country  `json:"country"`
	DatePaid datePaid `json:"datePaid"`
}

type hitungGajiRequestBody struct {
	Employee     employee          `json:"employee"`
	KomponenGaji []salaryStructure `json:"komponengaji"`
}

func calculateYearlyNetIncomePerCountryPerYear(komponenGaji []salaryStructure) map[int]map[int]float64 {
	nestedMap := map[int]map[int]float64{}

	for _, salary := range komponenGaji {
		_, ok := nestedMap[int(salary.Country)]
		if !ok {
			nestedMap[int(salary.Country)] = map[int]float64{}
		}

		thisYear := salary.DatePaid.Time.Year()
		nestedMap[int(salary.Country)][thisYear] += salary.Value
	}

	return nestedMap
}

var p = message.NewPrinter(language.Indonesian)

func POST(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var body hitungGajiRequestBody
	err := decoder.Decode(&body)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		helper.Print(err, r)
		return
	}

	yearlyNetIncomePerCountryPerYear := calculateYearlyNetIncomePerCountryPerYear(body.KomponenGaji)

	// out of scope requirements
	// komponengaji key includes salaries from 2 different countries.
	if len(yearlyNetIncomePerCountryPerYear) != 1 {
		http.Error(w, "Not Implemented. komponengaji key includes salaries from 2 different countries.", http.StatusNotImplemented)
		return
	}

	// out of scope requirements
	// the komponengaji key includes salaries from 2 different years.
	for _, value := range yearlyNetIncomePerCountryPerYear {
		if len(value) != 1 {
			http.Error(w, "Not Implemented. the komponengaji key includes salaries from 2 different years.", http.StatusNotImplemented)
			return
		}
	}

	for countryInt, years := range yearlyNetIncomePerCountryPerYear {
		for _, year := range years {
			yearlyTaxableIncome, err := CalculateYearlyTaxableIncome(country(countryInt), body.Employee.MarriageStatus, year)
			if err != nil {
				http.Error(w, "Not Implemented. "+err.Error(), http.StatusNotImplemented)
				return
			}
			yearlyTax, err := CalculateYearlyTax(country(countryInt), yearlyTaxableIncome)
			if err != nil {
				http.Error(w, "Not Implemented. "+err.Error(), http.StatusNotImplemented)
				return
			}
			monthlyTax := yearlyTax / 12
			w.Write([]byte(p.Sprintf("%.0f", monthlyTax)))
			return
		}
	}
}

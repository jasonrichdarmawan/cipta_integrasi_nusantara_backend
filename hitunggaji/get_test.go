package hitunggaji_test

import (
	"net/http"
	"testing"

	"github.com/jasonrichdarmawan/cipta_integrasi_nusantara_backend/hitunggaji"
	"github.com/steinfletcher/apitest"
)

func TestGET(t *testing.T) {
	// Scenario 1
	// An employee in Indonesia, married and have children
	// with monthly net income 30 millions Rupiah
	apitest.New().
		HandlerFunc(hitunggaji.GET).
		Post("/hitunggaji").
		JSON(`{
			"employee": {
				"marriageStatus": 3
			},
			"komponengaji": [
                {"value": 30000000, "datePaid": "25-01-2006", "country": 0},
				{"value": 30000000, "datePaid": "25-02-2006", "country": 0},
				{"value": 30000000, "datePaid": "25-03-2006", "country": 0},
				{"value": 30000000, "datePaid": "25-04-2006", "country": 0},
				{"value": 30000000, "datePaid": "25-05-2006", "country": 0},
				{"value": 30000000, "datePaid": "25-06-2006", "country": 0},
				{"value": 30000000, "datePaid": "25-07-2006", "country": 0},
				{"value": 30000000, "datePaid": "25-08-2006", "country": 0},
				{"value": 30000000, "datePaid": "25-09-2006", "country": 0},
				{"value": 30000000, "datePaid": "25-10-2006", "country": 0},
				{"value": 30000000, "datePaid": "25-11-2006", "country": 0},
				{"value": 30000000, "datePaid": "25-12-2006", "country": 0}
			]
		}`).
		Expect(t).
		Body("2.166.667").
		Status(http.StatusOK).
		End()

	// Scenario 2
	// An employee in Vietnam, married
	// with monthly net income 30 million VND
	// monthly insurance cost 1 million VND
	apitest.New().
		HandlerFunc(hitunggaji.GET).
		Post("/hitunggaji").
		JSON(`{
			"employee": {
				"marriageStatus": 1
			},
			"komponengaji": [
                {"value": 30000000, "datePaid": "25-01-2006", "country": 1},
				{"value": -1000000, "datePaid": "25-01-2006", "country": 1},
				{"value": 30000000, "datePaid": "25-02-2006", "country": 1},
				{"value": -1000000, "datePaid": "25-02-2006", "country": 1},
				{"value": 30000000, "datePaid": "25-03-2006", "country": 1},
				{"value": -1000000, "datePaid": "25-03-2006", "country": 1},
				{"value": 30000000, "datePaid": "25-04-2006", "country": 1},
				{"value": -1000000, "datePaid": "25-04-2006", "country": 1},
				{"value": 30000000, "datePaid": "25-05-2006", "country": 1},
				{"value": -1000000, "datePaid": "25-05-2006", "country": 1},
				{"value": 30000000, "datePaid": "25-06-2006", "country": 1},
				{"value": -1000000, "datePaid": "25-06-2006", "country": 1},
				{"value": 30000000, "datePaid": "25-07-2006", "country": 1},
				{"value": -1000000, "datePaid": "25-07-2006", "country": 1},
				{"value": 30000000, "datePaid": "25-08-2006", "country": 1},
				{"value": -1000000, "datePaid": "25-08-2006", "country": 1},
				{"value": 30000000, "datePaid": "25-09-2006", "country": 1},
				{"value": -1000000, "datePaid": "25-09-2006", "country": 1},
				{"value": 30000000, "datePaid": "25-10-2006", "country": 1},
				{"value": -1000000, "datePaid": "25-10-2006", "country": 1},
				{"value": 30000000, "datePaid": "25-11-2006", "country": 1},
				{"value": -1000000, "datePaid": "25-11-2006", "country": 1},
				{"value": 30000000, "datePaid": "25-12-2006", "country": 1},
				{"value": -1000000, "datePaid": "25-12-2006", "country": 1}
			]
		}`).
		Expect(t).
		Body("1.779.167").
		Status(http.StatusOK).
		End()
}

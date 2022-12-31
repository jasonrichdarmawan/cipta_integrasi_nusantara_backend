package validasialergiobat_test

import (
	"net/http"
	"testing"

	"github.com/jasonrichdarmawan/cipta_integrasi_nusantara_backend/validasialergiobat"
	"github.com/steinfletcher/apitest"
)

func TestGET(t *testing.T) {
	validasialergiobat.InitializeDB()

	// Scenario 1
	// A patient with ibuprofen allergy was prescribed by a doctor with
	// Proris sirup 60ml, Paratusin sirup 60ml
	apitest.New().
		HandlerFunc(validasialergiobat.GET).
		Get("/validasialergiobat").
		JSON(`{
			"pasien": {
				"name": "Bejo",
				"allergies": ["ibuprofen"]
			},
			"resep": [
				{"obat": "Proris sirup 60ml"},
				{"obat": "Paratusin sirup 60ml"}
			]
		}`).
		Expect(t).
		Body(`{
			"resep": [
				{
					"obat": "Proris sirup 60ml"
				}
			]
		}`).
		Status(http.StatusOK).
		End()

	// Scenario 2 out of scope requirements
	// return "unknown" key if the medicine is not in the database.
	apitest.New().
		HandlerFunc(validasialergiobat.GET).
		Get("/validasialergiobat").
		JSON(`{
			"pasien": {
				"name": "Bejo",
				"allergies": ["ibuprofen"]
			},
			"resep": [
				{"obat": "Proris sirup 60ml"},
				{"obat": "Paratusin sirup 60ml"},
				{"obat": "Amoxicillin ER 775 Mg Tablet"}
			]
		}`).
		Expect(t).
		Body(`{
			"resep": [
				{
					"obat": "Proris sirup 60ml"
				},
				{
					"unknown": "Amoxicillin ER 775 Mg Tablet"
				}
			]
		}`).
		Status(http.StatusOK).
		End()
}

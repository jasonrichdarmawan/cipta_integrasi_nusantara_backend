package validasialergiobat

import (
	"encoding/json"
	"net/http"

	"github.com/jasonrichdarmawan/cipta_integrasi_nusantara_backend/helper"
)

type patient struct {
	Name      string   `json:"name"`
	Allergies []string `json:"allergies"`
}

type resep struct {
	Obat string `json:"obat"`
}

type validasiAlergiObatRequestBody struct {
	Patient patient `json:"pasien"`
	Resep   []resep `json:"resep"`
}

type validasiAlergiObatResponseBody struct {
	Resep []map[string]string `json:"resep"`
}

func GET(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var requestBody validasiAlergiObatRequestBody
	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		helper.Print(err, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	p := validasiAlergiObatResponseBody{}

	for _, resep := range requestBody.Resep {
		if !IsMedicineExistsInMedicineDB(resep.Obat) {
			p.Resep = append(p.Resep, map[string]string{"unknown": resep.Obat})
			continue
		}

		for _, allergy := range requestBody.Patient.Allergies {
			if IsProductHaveCommonCompositionThatCauseAllergies(allergy, resep.Obat) {
				p.Resep = append(p.Resep, map[string]string{"obat": resep.Obat})
				break
			}
		}
	}

	responseBody, err := json.Marshal(p)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	w.Write(responseBody)
}

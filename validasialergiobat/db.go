package validasialergiobat

import (
	"strings"
)

// {"Paratusin 10 Tablet": ["Noscapine 10 mg", ...]}
var medicineDB map[string][]string

// {"ibuprofen": ["proris"]}
var allergyDB map[string][]string

func CheckAndInsertIntoAllergyDBIfProductHaveCommonCompositionThatCauseAllergies(productName string, productCompositions []string) {
	for _, productComposition := range productCompositions {
		for commonCompositionThatCauseAllergies := range allergyDB {

			// Does that product composition exists in the allergyDB?
			if strings.Contains(strings.ToLower(productComposition), commonCompositionThatCauseAllergies) {

				// Does that product name exists in the allergyDB[commonComposition]?
				for _, commonProductThatCauseAllergies := range allergyDB[commonCompositionThatCauseAllergies] {
					if commonProductThatCauseAllergies == productName {
						continue
					}
					allergyDB[commonCompositionThatCauseAllergies] = append(allergyDB[commonCompositionThatCauseAllergies], productName)
				}
			}
		}
	}
}

func IsProductHaveCommonCompositionThatCauseAllergies(compositionThatCauseAllergies string, productName string) bool {
	for _, product := range allergyDB[compositionThatCauseAllergies] {
		if product == productName {
			return true
		}
	}

	return false
}

func InsertToMedicineDB(productName string, productCompositions []string) {
	medicineDB[productName] = productCompositions
	CheckAndInsertIntoAllergyDBIfProductHaveCommonCompositionThatCauseAllergies(productName, productCompositions)
}

func IsMedicineExistsInMedicineDB(productName string) bool {
	for product := range medicineDB {
		if strings.EqualFold(product, productName) {
			return true
		}
	}

	return false
}

func InsertIntoAllergyDB(compositionThatCauseAllergies string, productName string) error {
	if IsProductHaveCommonCompositionThatCauseAllergies(compositionThatCauseAllergies, productName) {
		return nil
	}
	allergyDB[compositionThatCauseAllergies] = append(allergyDB[compositionThatCauseAllergies], productName)
	return nil
}

// reference: https://www.halodoc.com/obat-dan-vitamin/proris-suspensi-60-ml
func InitializeDB() {
	medicineDB = map[string][]string{}
	allergyDB = map[string][]string{}

	// Scenario 1:
	// A doctor input a new product and believe there are no common compositions in the product that cause allergies.
	// In fact, there are not.
	paratusinName := "Paratusin 10 Tablet"
	paratusinCompositions := []string{
		"Noscapine 10 mg", "chlorpheniramine maleate 2 mg",
		"glyceryl guaiacolate 50 mg", "paracetamol 500 mg",
		"phenylpropanolamine HCl 15 mg",
	}
	InsertToMedicineDB(paratusinName, paratusinCompositions)

	paratusin2Name := "Paratusin sirup 60ml"
	paratusin2Compositions := []string{
		"kandungan utama paracetamol 125 mg",
		"disamping kandungan lainnya terdiri dari pseudoepedrid 7.5 mg",
		"noscapine 10 mg", "ctm 0.5 mg", "guafenisin 25 mg",
		"succus liquiritae 125 ethanol 10 %.",
	}
	InsertToMedicineDB(paratusin2Name, paratusin2Compositions)

	// Scenario 2:
	// A doctor input a new product and believe the product have common compositions that cause allergies.
	prorisName := "Proris Suspensi 60ml"
	prorisCompositions := []string{"Tiap 5 ml mengandung : Ibuprofen 100 mg"}
	InsertToMedicineDB(prorisName, prorisCompositions)
	InsertIntoAllergyDB("ibuprofen", prorisName)

	proris2Name := "Proris sirup 60ml"
	proris2Compositions := []string{"Ibuprofen"}
	InsertToMedicineDB(proris2Name, proris2Compositions)
	InsertIntoAllergyDB("ibuprofen", proris2Name)

	aspirinName := "Cardio Aspirin 100 mg 10 Tablet"
	aspirinCompositions := []string{"Tiap tablet mengandung: Acetylsalicylic acid 100 mg"}
	InsertToMedicineDB(aspirinName, aspirinCompositions)
	InsertIntoAllergyDB("acetylsalicylic acid", aspirinName)

	// Scenario 3:
	// A doctor input a new product and believe there are no common compositions in the product that cause allergies.
	// In fact, there are.
	aspirin2Name := "Ascardia 80 mg 10 Tablet"
	aspirin2Compositions := []string{"Tiap tablet mengandung: Acetylsalicylic acid 80 mg"}
	InsertToMedicineDB(aspirin2Name, aspirin2Compositions)
}

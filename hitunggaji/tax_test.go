package hitunggaji_test

import (
	"testing"

	"github.com/jasonrichdarmawan/cipta_integrasi_nusantara_backend/hitunggaji"
)

func TestCalculateTaxableIncome(t *testing.T) {
	// Scenario 1:
	// An employee in Indonesia, married and have children,
	// monthly salary 30 million Rupiah
	netIncomeIDN, err := hitunggaji.CalculateYearlyTaxableIncome(hitunggaji.CountryIDN, hitunggaji.MARRIED_HAVE_CHILDREN, 30*1e6*12)
	if err != nil {
		t.Error(err)
	}
	if netIncomeIDN != 285e6 {
		t.Errorf("netIncomeIDN = %f; want 285e6", netIncomeIDN)
	}

	// Scenario 2:
	// An employee in Vietnam, married,
	// monthly salary 30 million VND, monthly insurance cost 1 million VND
	netIncomeVNM, err := hitunggaji.CalculateYearlyTaxableIncome(hitunggaji.CountryVNM, hitunggaji.MARRIED, 30*1e6*12-1*1e6*12)

	if err != nil {
		t.Error(err)
	}

	if netIncomeVNM != 318e6 {
		t.Errorf("netIncomeVNM = %f; want 318e6", netIncomeVNM)
	}
}

func TestCalculateTax(t *testing.T) {
	// Scenario 1
	// An employee in Indonesia with 285*1e6 taxable income
	// Note: the scope requirements use 15% instead of 10% for layer 50-250
	// The correct one is 10%.
	taxIDN, err := hitunggaji.CalculateYearlyTax(hitunggaji.CountryIDN, 285*1e6)
	if err != nil {
		t.Error(err)
	}
	if taxIDN != 26*1e6 {
		t.Errorf("taxIDN = %f; want 26*1e6", taxIDN)
	}

	// Scenario 2
	// An employee in Vietnam with 318*1e6 taxable income
	taxVNM, err := hitunggaji.CalculateYearlyTax(hitunggaji.CountryVNM, 318*1e6)
	if err != nil {
		t.Error(err)
	}
	if taxVNM != 21.35*1e6 {
		t.Errorf("taxVNM = %f; want 21.35*1e6", taxVNM)
	}

}

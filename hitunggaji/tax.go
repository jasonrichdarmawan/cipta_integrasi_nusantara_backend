package hitunggaji

import (
	"errors"
	"strconv"
	"time"
)

type marriageStatus uint16

const (
	NOT_MARRIED marriageStatus = iota
	MARRIED
	MARRIED_NO_CHILDREN
	MARRIED_HAVE_CHILDREN
)

type employee struct {
	MarriageStatus marriageStatus `json:"marriageStatus"`
}

type country int

const (
	CountryIDN country = iota // Indonesia
	CountryVNM                // Vietnam
)

type datePaid struct {
	time.Time
}

func (t *datePaid) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	var err error
	tt, err := time.Parse(`"02-01-2006"`, string(data))
	*t = datePaid{tt}
	return err
}

func calculateYearlyTaxableIncomeIDN(marriageStatus marriageStatus, yearlyNetIncome float64) (float64, error) {
	var nonTaxableIncome float64
	switch marriageStatus {
	case NOT_MARRIED:
		nonTaxableIncome = 25 * 1e6
	case MARRIED_NO_CHILDREN:
		nonTaxableIncome = 50 * 1e6
	case MARRIED_HAVE_CHILDREN:
		nonTaxableIncome = 75 * 1e6
	default:
		return 0, errors.New("marriage status " + strconv.FormatUint(uint64(marriageStatus), 10) + " is not supported")
	}

	return yearlyNetIncome - nonTaxableIncome, nil
}

func calculateYearlyTaxableIncomeVNM(marriageStatus marriageStatus, yearlyNetIncome float64) (float64, error) {
	var nonTaxableIncome float64
	switch marriageStatus {
	case NOT_MARRIED:
		nonTaxableIncome = 15 * 1e6
	case MARRIED:
		nonTaxableIncome = 30 * 1e6
	default:
		return 0, errors.New("marriage status " + strconv.FormatUint(uint64(marriageStatus), 10) + " is not supported")
	}

	return yearlyNetIncome - nonTaxableIncome, nil
}

func CalculateYearlyTaxableIncome(country country, marriageStatus marriageStatus, yearlyNetIncome float64) (float64, error) {
	switch country {
	case CountryIDN:
		return calculateYearlyTaxableIncomeIDN(marriageStatus, yearlyNetIncome)
	case CountryVNM:
		return calculateYearlyTaxableIncomeVNM(marriageStatus, yearlyNetIncome)
	default:
		return 0, errors.New("country " + strconv.FormatUint(uint64(country), 10) + " not supported.")
	}
}

func calculateYearlyTaxIDN(yearlyTaxableIncome float64) float64 {
	var tax float64

	// layer 0-50 millions Rupiah
	if yearlyTaxableIncome < 50*1e6 {
		tax += yearlyTaxableIncome * 5 / 100
		return tax
	} else {
		tax += 50 * 1e6 * 5 / 100
		yearlyTaxableIncome -= 50 * 1e6
	}

	// layer 50-250 millions Rupiah
	if yearlyTaxableIncome < 250*1e6 {
		tax += yearlyTaxableIncome * 10 / 100
		return tax
	} else {
		tax += 250 * 1e6 * 10 / 100
		yearlyTaxableIncome -= 250 * 1e6
	}

	// layer >250 millions Rupiah
	tax += yearlyTaxableIncome * 15 / 100
	return tax
}

func calculateYearlyTaxVNM(yearlyTaxableIncome float64) float64 {
	var tax float64

	// layer 0 - 50 milllions VND
	if yearlyTaxableIncome < 50*1e6 {
		tax += yearlyTaxableIncome * 2.5 / 100
		return tax
	} else {
		tax += 50 * 1e6 * 2.5 / 100
		yearlyTaxableIncome -= 50 * 1e6
	}

	// layer >50 millions VND
	tax += yearlyTaxableIncome * 7.5 / 100
	return tax
}

func CalculateYearlyTax(country country, yearlyTaxableIncome float64) (float64, error) {
	if yearlyTaxableIncome < 0 {
		return 0, nil
	}

	switch country {
	case CountryIDN:
		return calculateYearlyTaxIDN(yearlyTaxableIncome), nil
	case CountryVNM:
		return calculateYearlyTaxVNM(yearlyTaxableIncome), nil
	default:
		return 0, errors.New("country " + strconv.FormatUint(uint64(country), 10) + " not supported.")
	}
}

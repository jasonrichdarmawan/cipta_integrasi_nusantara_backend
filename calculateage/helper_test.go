package calculateage_test

import (
	"testing"
	"time"

	"github.com/jasonrichdarmawan/cipta_integrasi_nusantara_backend/calculateage"
)

func TestLeapYear(t *testing.T) {
	years := []int{1900, 1904, 2100, 2104, 2200, 2204, 2300, 2304}
	for _, year := range years {
		// If leap year, then true.
		if year%4 == 0 && year%100 != 0 || year%400 == 0 {

			// check year against know not leap years.
			for notLeapYear := range []int{1900, 2100, 2200, 2300} {
				if year == notLeapYear {
					t.Errorf("%d = leap year; want not leap year", notLeapYear)
					break
				}
			}
		}
	}
}

func TestCalculateDaysBetweenDates(t *testing.T) {
	// Scenario 1: leap year && a.YearDay() <= b.YearDay()
	a, err := time.Parse("02-01-2006", "01-01-1924")
	if err != nil {
		t.Error(err)
		return
	}
	b, err := time.Parse("02-01-2006", "01-01-1925")
	if err != nil {
		t.Error(err)
		return
	}

	yearsDiff, monthsDiff, daysDiff, err := calculateage.CalculateYearsMonthsDaysBetweenDates(a, b)
	if err != nil {
		t.Error(err)
		return
	}
	if yearsDiff != 1 || monthsDiff != 0 || daysDiff != 0 {
		t.Errorf("yearsDiff = %d, monthsDiff = %d, daysDiff = %f; want 1, 0, 0", yearsDiff, monthsDiff, daysDiff)
	}

	// Scenario 2: leap year && a.YearDay() > b.YearDay()
	c, err := time.Parse("02-01-2006", "02-01-1924")
	if err != nil {
		t.Error(err)
		return
	}
	d, err := time.Parse("02-01-2006", "01-01-1925")
	if err != nil {
		t.Error(err)
		return
	}

	yearsDiff, monthsDiff, daysDiff, err = calculateage.CalculateYearsMonthsDaysBetweenDates(c, d)
	if err != nil {
		t.Error(err)
		return
	}
	if yearsDiff != 0 || monthsDiff != 11 || daysDiff != 30 {
		t.Errorf("yearsDiff = %d, monthsDiff = %d, daysDiff = %f; want 0, 11, 30", yearsDiff, monthsDiff, daysDiff)
	}

	// Scenario 3: leap year + not leap year && a.YearDay() <= b.YearDay()
	e, err := time.Parse("02-01-2006", "01-01-1924")
	if err != nil {
		t.Error(err)
		return
	}
	f, err := time.Parse("02-01-2006", "01-01-1926")
	if err != nil {
		t.Error(err)
		return
	}

	yearsDiff, monthsDiff, daysDiff, err = calculateage.CalculateYearsMonthsDaysBetweenDates(e, f)
	if err != nil {
		t.Error(err)
		return
	}
	if yearsDiff != 2 || monthsDiff != 0 || daysDiff != 0 {
		t.Errorf("yearsDiff = %d, monthsDiff = %d, daysDiff = %f; want 2, 0, 0", yearsDiff, monthsDiff, daysDiff)
	}

	// Scenario 4: leap year + not leap year && a.YearDay() > b.YearDay()
	g, err := time.Parse("02-01-2006", "02-01-1924")
	if err != nil {
		t.Error(err)
		return
	}
	h, err := time.Parse("02-01-2006", "01-01-1926")
	if err != nil {
		t.Error(err)
		return
	}

	yearsDiff, monthsDiff, daysDiff, err = calculateage.CalculateYearsMonthsDaysBetweenDates(g, h)
	if err != nil {
		t.Error(err)
		return
	}
	if yearsDiff != 1 || monthsDiff != 11 || daysDiff != 30 {
		t.Errorf("yearsDiff = %d, monthsDiff = %d, daysDiff = %f; want 1, 11, 30", yearsDiff, monthsDiff, daysDiff)
	}

	// Scenario 5: leap year + not leap year + not leap year + not leap year + leap year && a.YearDay() <= b.YearDay()
	i, err := time.Parse("02-01-2006", "01-01-1924")
	if err != nil {
		t.Error(err)
		return
	}
	j, err := time.Parse("02-01-2006", "01-01-1929")
	if err != nil {
		t.Error(err)
		return
	}

	yearsDiff, monthsDiff, daysDiff, err = calculateage.CalculateYearsMonthsDaysBetweenDates(i, j)
	if err != nil {
		t.Error(err)
		return
	}
	if yearsDiff != 5 || monthsDiff != 0 || daysDiff != 0 {
		t.Errorf("yearsDiff = %d, monthsDiff = %d, daysDiff = %f; want 5, 0, 0", yearsDiff, monthsDiff, daysDiff)
	}

	// Scenario 6: leap year + not leap year + not leap year + not leap year + leap year && a.YearDay() <= b.YearDay()
	k, err := time.Parse("02-01-2006", "02-01-1924")
	if err != nil {
		t.Error(err)
		return
	}
	l, err := time.Parse("02-01-2006", "01-01-1929")
	if err != nil {
		t.Error(err)
		return
	}

	yearsDiff, monthsDiff, daysDiff, err = calculateage.CalculateYearsMonthsDaysBetweenDates(k, l)
	if err != nil {
		t.Error(err)
		return
	}
	if yearsDiff != 4 || monthsDiff != 11 || daysDiff != 30 {
		t.Errorf("yearsDiff = %d, monthsDiff = %d, daysDiff = %f; want 4, 11, 30", yearsDiff, monthsDiff, daysDiff)
	}
}

func TestCalculateDaysBetweenDates2(t *testing.T) {
	// Scenario 7: start from month 10 instead of 1
	m, err := time.Parse("02-01-2006", "28-10-1983")
	if err != nil {
		t.Error(err)
		return
	}
	n, err := time.Parse("02-01-2006", "01-01-2023")
	if err != nil {
		t.Error(err)
		return
	}

	yearsDiff, monthsDiff, daysDiff, err := calculateage.CalculateYearsMonthsDaysBetweenDates(m, n)
	if err != nil {
		t.Error(err)
		return
	}
	if yearsDiff != 39 || monthsDiff != 2 || daysDiff != 4 {
		t.Errorf("yearsDiff = %d, monthsDiff = %d, daysDiff = %f; want 39, 2, 4", yearsDiff, monthsDiff, daysDiff)
	}
}

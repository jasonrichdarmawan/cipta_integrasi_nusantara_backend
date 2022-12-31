package calculateage

import (
	"errors"
	"time"
)

func CalculateYearsMonthsDaysBetweenDates(a, b time.Time) (yearsDiff int, monthsDiff int, daysDiff float64, err error) {
	if !a.Before(b) {
		return 0, 0, 0, errors.New("a.Before(b) should be true")
	}

	yearsDiff = b.Year() - a.Year()

	if a.YearDay() > b.YearDay() {
		yearsDiff -= 1
	}

	a = a.AddDate(yearsDiff, 0, 0)

	daysDiff = b.Sub(a).Hours() / 24

	year := a.Year()

	// reference: https://www.calendarr.com/united-states/how-many-days-in-a-month/
	daysInMonths := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	// reference: https://www.calendar.best/leap-years.html
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		daysInMonths[1] = 29
	}

	monthsDiff = 0
	for _, daysInMonth := range daysInMonths[a.Month()-1:] {
		daysDiff -= float64(daysInMonth)
		if daysDiff < 0 {
			daysDiff += float64(daysInMonth)
			break
		}

		monthsDiff += 1
	}

	return
}

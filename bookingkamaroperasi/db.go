package bookingkamaroperasi

import "time"

type BookingKamarOperasi struct {
	startDate time.Time
	endDate   time.Time
}

var DB = map[time.Time][]BookingKamarOperasi{}

func truncateDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func IsThereNoScheduleIn2Hours(key time.Time, bookingDate time.Time) bool {
	document := DB[key]

	for _, tuple := range document {
		diff := bookingDate.Sub(tuple.endDate)

		if diff < 2*time.Hour {
			return false
		}
	}

	return true
}

func TryAppend(bookingDate time.Time, duration time.Duration) bool {
	yesterday := truncateDate(bookingDate.AddDate(0, 0, -1))
	if !IsThereNoScheduleIn2Hours(yesterday, bookingDate) {
		return false
	}

	today := truncateDate(bookingDate)
	if !IsThereNoScheduleIn2Hours(today, bookingDate) {
		return false
	}

	DB[today] = append(DB[today],
		BookingKamarOperasi{
			startDate: bookingDate,
			endDate:   bookingDate.Add(duration),
		})

	return true
}

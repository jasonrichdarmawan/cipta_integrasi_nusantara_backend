package bookingkamaroperasi

import (
	"sync"
	"time"
)

type BookingKamarOperasi struct {
	startDate time.Time
	endDate   time.Time
}

type SafeDB struct {
	mu sync.RWMutex
	v  map[time.Time][]BookingKamarOperasi
}

func (c *SafeDB) Append(key time.Time, bookingKamarOperasi BookingKamarOperasi) {
	c.mu.Lock()
	c.v[key] = append(c.v[key], bookingKamarOperasi)
	c.mu.Unlock()
}

func (c *SafeDB) Value(key time.Time) []BookingKamarOperasi {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.v[key]
}

var c = SafeDB{v: make(map[time.Time][]BookingKamarOperasi)}

func truncateDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func isThereNoScheduleIn2Hours(key time.Time, bookingDate time.Time) bool {
	document := c.Value(key)

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
	if !isThereNoScheduleIn2Hours(yesterday, bookingDate) {
		return false
	}

	today := truncateDate(bookingDate)
	if !isThereNoScheduleIn2Hours(today, bookingDate) {
		return false
	}

	c.Append(today,
		BookingKamarOperasi{
			startDate: bookingDate,
			endDate:   bookingDate.Add(duration),
		})

	return true
}

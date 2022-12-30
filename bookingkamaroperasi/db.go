package bookingkamaroperasi

import (
	"log"
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

func isThereNoScheduleIn2Hours(key time.Time, bookingDate time.Time, duration time.Duration) bool {
	document := c.Value(key)

	log.Println(len(document), bookingDate.Format("02-01-2006 15:04"))

	endDate := bookingDate.Add(duration)
	startDate := bookingDate

	for _, tuple := range document {
		var diff time.Duration
		if endDate.Before(tuple.startDate) {
			// scenario 1: endDate tuple.startDate

			diff = tuple.startDate.Sub(endDate)
			log.Println("endDate", endDate.Format("02-01-2006 15:04"), "tuple.startDate", tuple.startDate.Format("02-01-2006 15:04"), "diff", diff)
		} else {
			// scenario 2: tuple.startDate endDate

			if startDate.Before(tuple.endDate) {
				// tuple.startDate startDate tuple.endDate

				return false
			} else {

				// tuple.startDate tuple.endDate startDate
				diff = startDate.Sub(tuple.endDate)
				log.Println("endDate", startDate.Format("02-01-2006 15:04"), "tuple.endDate", tuple.endDate.Format("02-01-2006 15:04"), "diff", diff)
			}
		}

		// scenario 1: endDate <-diff-> tuple.startDate
		// scenario 2: tuple.startDate tuple.endDate <-diff-> startDate
		if diff < 2*time.Hour {
			return false
		}
	}

	return true
}

func TryAppend(bookingDate time.Time, duration time.Duration) bool {
	yesterday := truncateDate(bookingDate.AddDate(0, 0, -1))
	if !isThereNoScheduleIn2Hours(yesterday, bookingDate, duration) {
		return false
	}

	today := truncateDate(bookingDate)
	if !isThereNoScheduleIn2Hours(today, bookingDate, duration) {
		return false
	}

	tomorrow := truncateDate(bookingDate.AddDate(0, 0, 1))
	if !isThereNoScheduleIn2Hours(tomorrow, bookingDate, duration) {
		return false
	}

	c.Append(today,
		BookingKamarOperasi{
			startDate: bookingDate,
			endDate:   bookingDate.Add(duration),
		})

	return true
}

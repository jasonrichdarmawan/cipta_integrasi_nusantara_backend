package bookingkamaroperasi_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/jasonrichdarmawan/cipta_integrasi_nusantara/bookingkamaroperasi"
	"github.com/steinfletcher/apitest"
)

func TestTryAppend(t *testing.T) {
	now, err := time.Parse("02-01-2006 15:04Z07:00", "01-01-2022 10:00+07:00")
	if err != nil {
		t.Error(err)
	}

	// A doctor books on 01/01/2022 10:00 for 2 hours
	if got := bookingkamaroperasi.TryAppend(now, 2*time.Hour); !got {
		t.Errorf("%v-12:00 = %t; want true", now.Format("02-01-2006 15:04"), got)
	}

	// Another doctor books on 01/01/2022 15:00 for 2 hours
	today15 := now.Add(5 * time.Hour)
	if got := bookingkamaroperasi.TryAppend(today15, 2*time.Hour); !got {
		t.Errorf("%v-17:00 = %t; want true", today15.Format("02-01-2006 15:04"), got)
	}

	// Another doctor books on 01/01/2022 18:00 for 2 hours
	today18 := now.Add(8 * time.Hour)
	if got := bookingkamaroperasi.TryAppend(today18, 2*time.Hour); got {
		t.Errorf("%v-20:00 = %t; want false", today18.Format("02-01-2006 15:04"), got)
	}

	// Another doctor books on 01/01/2022 07:00 for 2 hours
	today07 := now.Add(-3 * time.Hour)
	if got := bookingkamaroperasi.TryAppend(today07, 2*time.Hour); got {
		t.Errorf("%v-09:00 = %t; want false", today07.Format("02-01-2006 15:04"), got)
	}

	// out of scope
	// Another doctor books on 01/01/2022 02:00 for 2 hours
	today02 := now.Add(-8 * time.Hour)
	if got := bookingkamaroperasi.TryAppend(today02, 2*time.Hour); !got {
		t.Errorf("%v-04:00 = %t; want false", today02.Format("02-01-2006 15:04"), got)
	}
	// Another doctor books on 31/12/2022 23:00 for 2 hours
	yesterday23_today01 := now.Add(-11 * time.Hour)
	if got := bookingkamaroperasi.TryAppend(yesterday23_today01, 2*time.Hour); got {
		t.Errorf("%v - 01-01-2022 01:00 = %t; want false", yesterday23_today01.Format("02-01-2006 15:04"), got)
	}
}

func TestPOST(t *testing.T) {
	apitest.New().
		HandlerFunc(bookingkamaroperasi.POST).
		Post("/bookingkamaroperasi/01-01-2022 10:00+07:00/120").
		Expect(t).
		Body("true").
		Status(http.StatusOK).
		End()

	apitest.New().
		HandlerFunc(bookingkamaroperasi.POST).
		Post("/bookingkamaroperasi/01-01-2022 15:00+07:00/120").
		Expect(t).
		Body("true").
		Status(http.StatusOK).
		End()

	apitest.New().
		HandlerFunc(bookingkamaroperasi.POST).
		Post("/bookingkamaroperasi/01-01-2022 18:00+07:00/120").
		Expect(t).
		Body("false").
		Status(http.StatusOK).
		End()

	apitest.New().
		HandlerFunc(bookingkamaroperasi.POST).
		Post("/bookingkamaroperasi/01-01-2022 07:00+07:00/120").
		Expect(t).
		Body("false").
		Status(http.StatusOK).
		End()
}

package bookingkamaroperasi_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/jasonrichdarmawan/cipta_integrasi_nusantara/bookingkamaroperasi"
	"github.com/steinfletcher/apitest"
)

func TestTryAppend(t *testing.T) {
	now := time.Now()

	// A doctor books on 01/01/2022 10:00 for 2 hours
	if got := bookingkamaroperasi.TryAppend(now, 2*time.Hour); !got {
		t.Errorf("TryAppend(10:00-12:00) = %t; want true", got)
	}

	// Another doctor books on 01/01/2022 15:00 for 2 hours
	if got := bookingkamaroperasi.TryAppend(now.Add(5*time.Hour), 2*time.Hour); !got {
		t.Errorf("TryAppend(15:00-17:00) = %t; want true", got)
	}

	// Another doctor books on 01/01/2022 18:00 for 2 hours
	if got := bookingkamaroperasi.TryAppend(now.Add(8*time.Hour), 2*time.Hour); got {
		t.Errorf("TryAppend(18:00-20:00) = %t; want false", got)
	}

	// Another doctor books on 01/01/2022 07:00 for 2 hours
	if got := bookingkamaroperasi.TryAppend(now.Add(-3*time.Hour), 2*time.Hour); got {
		t.Errorf("TryAppend(07:00-09:00) = %t; want false", got)
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

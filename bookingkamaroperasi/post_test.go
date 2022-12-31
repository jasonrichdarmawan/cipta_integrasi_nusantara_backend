package bookingkamaroperasi_test

import (
	"net/http"
	"testing"

	"github.com/jasonrichdarmawan/cipta_integrasi_nusantara_backend/bookingkamaroperasi"
	"github.com/steinfletcher/apitest"
)

func TestPOST(t *testing.T) {
	bookingkamaroperasi.InitializeDB()
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

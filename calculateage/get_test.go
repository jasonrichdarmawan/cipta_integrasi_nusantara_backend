package calculateage_test

import (
	"net/http"
	"testing"

	"github.com/jasonrichdarmawan/cipta_integrasi_nusantara_backend/calculateage"
	"github.com/steinfletcher/apitest"
)

func TestGET(t *testing.T) {
	// Scenario 1
	apitest.New().
		HandlerFunc(calculateage.GET).
		Get("/calculateage/28-10-1983").
		Expect(t).
		Body(`{
			"umur": {
				"year": 39,
				"month": 2,
				"day": 4
			}
		}`).
		Status(http.StatusOK).
		End()
}

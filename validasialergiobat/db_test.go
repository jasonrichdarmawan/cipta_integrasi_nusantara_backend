package validasialergiobat_test

import (
	"testing"

	"github.com/jasonrichdarmawan/cipta_integrasi_nusantara_backend/validasialergiobat"
)

func TestInitializeDB(t *testing.T) {
	validasialergiobat.InitializeDB()

	if got := validasialergiobat.IsMedicineExistsInMedicineDB("Paratusin 10 Tablet"); !got {
		t.Errorf("Paratusin 10 Tablet = %t; want true", got)
	}

	if got := validasialergiobat.IsMedicineExistsInMedicineDB("paratusin 10 tablet"); !got {
		t.Errorf("paratusin 10 tablet = %t; want true", got)
	}
}

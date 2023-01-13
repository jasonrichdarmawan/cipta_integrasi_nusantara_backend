package main

import (
	"log"
	"net/http"

	"github.com/jasonrichdarmawan/cipta_integrasi_nusantara_backend/bookingkamaroperasi"
	"github.com/jasonrichdarmawan/cipta_integrasi_nusantara_backend/calculateage"
	"github.com/jasonrichdarmawan/cipta_integrasi_nusantara_backend/hitunggaji"
	"github.com/jasonrichdarmawan/cipta_integrasi_nusantara_backend/validasialergiobat"
)

func main() {
	bookingkamaroperasi.InitializeDB_v2()
	validasialergiobat.InitializeDB()

	mux := http.NewServeMux()

	mux.HandleFunc("/bookingkamaroperasi/", bookingkamaroperasi.POST)
	mux.HandleFunc("/hitunggaji", hitunggaji.GET)
	mux.HandleFunc("/validasialergiobat", validasialergiobat.GET)
	mux.HandleFunc("/calculateage/", calculateage.GET)

	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		log.Panic(err)
	}
}

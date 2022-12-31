package main

import (
	"log"
	"net/http"

	"github.com/jasonrichdarmawan/cipta_integrasi_nusantara_backend/bookingkamaroperasi"
	"github.com/jasonrichdarmawan/cipta_integrasi_nusantara_backend/hitunggaji"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/bookingkamaroperasi/", bookingkamaroperasi.POST)
	mux.HandleFunc("/hitunggaji", hitunggaji.GET)

	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		log.Panic(err)
	}
}

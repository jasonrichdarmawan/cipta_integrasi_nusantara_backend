package main

import (
	"log"
	"net/http"

	"github.com/jasonrichdarmawan/cipta_integrasi_nusantara/bookingkamaroperasi"
	"github.com/jasonrichdarmawan/cipta_integrasi_nusantara/hitunggaji"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/bookingkamaroperasi/", bookingkamaroperasi.POST)
	mux.HandleFunc("/hitunggaji", hitunggaji.POST)

	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		log.Panic(err)
	}
}

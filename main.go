package main

import (
	"log"
	"net/http"

	"github.com/jasonrichdarmawan/cipta_integrasi_nusantara/bookingkamaroperasi"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/helloworld", bookingkamaroperasi.Post)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Panic(err)
	}
}

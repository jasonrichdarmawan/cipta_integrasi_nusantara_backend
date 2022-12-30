package helper

import (
	"log"
	"net/http"
)

func Print(err error, r *http.Request) {
	log.Println("ERROR" + " " + r.URL.Path + " " + r.RemoteAddr + " " + err.Error())
}

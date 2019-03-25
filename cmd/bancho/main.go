package main

import (
	"bancho"
	"log"
	"net/http"
	"time"
)

func main() {
	router := bancho.Router()

	server := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}

package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {

	// ========================================
	// start server with middleware
	// ========================================

	r := httpRouter()

	server := &http.Server{
		Handler: r,
		Addr:    "localhost:8080", // env var
	}
	log.Info("server up and running at ", server.Addr)
	log.Fatal(server.ListenAndServe())

}

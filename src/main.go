package main

import (
	"go-web-api/src/handler"
	"go-web-api/src/pg"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	// ========================================
	// connect to db and run migrations
	// ========================================
	// connect
	pgdb := pg.Connect()
	// migrate
	pg.HandleMigrations(pgdb)
	// inject

	// ========================================
	// start server with middleware
	// ========================================

	h := handler.NewHandler() // eventually connect the db here
	r := httpRouter(h)

	server := &http.Server{
		Handler: r,
		Addr:    "localhost:8080", // env var
	}
	log.Info("server up and running at ", server.Addr)
	log.Fatal(server.ListenAndServe())

}

package main

import (
	"fmt"
	"go-web-api/src/handler"
	"go-web-api/src/pg"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	// ========================================
	// connect to db and run migrations
	// ========================================
	if os.Getenv("ENV") != "production" && os.Getenv("DOCKER") != "true" {
		wd, err := os.Getwd()
		if err != nil {
			log.Error("error getting working directory: ", err)
		}
		err = godotenv.Load(fmt.Sprint(wd, "/../.env"))
		if err != nil {
			log.Error("an error occurred loading the env file: ", err)
		}
	}
	// connect
	pgdb := pg.Connect()
	// migrate
	pg.HandleMigrations(pgdb)
	// inject
	repo := pg.NewRepository(pgdb)
	// ========================================
	// start server with middleware
	// ========================================

	h := handler.NewHandler(repo) // eventually connect the db here
	r := httpRouter(h)

	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	server := &http.Server{
		Handler: r,
		Addr:    addr, // env var
	}
	log.Info("server up and running at ", server.Addr)
	log.Fatal(server.ListenAndServe())

}

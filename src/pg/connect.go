package pg

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

func Connect() *sql.DB {
	pgChan := make(chan *sql.DB, 1)
	var pgDSN string
	if os.Getenv("ENV") == "local" {
		log.Info("I'm a local!")
		pgHost := os.Getenv("POSTGRES_HOST")
		pgUser := os.Getenv("POSTGRES_USER")
		pgDBName := os.Getenv("POSTGRES_DB")
		pgPass := os.Getenv("POSTGRES_PASSWORD")
		// this string normally comes from the config (environment var)
		pgDSN = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s", pgUser, pgPass, pgDBName, pgHost)
	} else {
		pgDSN = os.Getenv("DATABASE_URL")
	}

	go func() {

		db, err := sql.Open("postgres", pgDSN)
		if err != nil {
			log.Fatal(fmt.Errorf("error connecting to postgres %+v", err))
		}

		// Ping the db to make sure we connected properly
		log.Info("Pinging the db")
		for {
			if err := db.Ping(); err != nil {
				log.Errorf("PostgresDB ERROR trying again in 30 seconds: %v\n", err)
				time.Sleep(time.Second * 30)
			} else {
				log.Info("connected to db")
				break
			}
		}

		pgChan <- db
	}()

	pgDB := <-pgChan
	return pgDB
}

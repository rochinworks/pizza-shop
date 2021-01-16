package pg

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	log "github.com/sirupsen/logrus"
)

func HandleMigrations(pg *sql.DB) error {
	if pg == nil {
		log.Fatal("Database is nil")
	}
	driver, err := postgres.WithInstance(pg, &postgres.Config{})
	if err != nil {
		log.Error("an error occurred with migrations1: ", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://pg/migrations", "postgres", driver)
	if err != nil {
		log.Error("error when migrating: ", err)
	}

	version, _, err := m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		log.WithField("version", version).Error(err)
		return err
	}

	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Error("an error occurred with migrations: ", err)
	} else if err == migrate.ErrNoChange {
		log.Info("no migration needed, moving on")
	} else {
		nversion, _, err := m.Version()
		if err != nil {
			log.Error(err)
			return err
		}
		log.Infof("migrated PG DB from version %d to version %d", version, nversion)
	}
	//	if err := m.Down(); err != nil {
	//		log.Error("an error occurred with migrations3: ", err)
	//	}

	//// when stupid, we can reset to specific versions
	//if err = m.Force(0); err != nil {
	//	log.Error(err)
	//	return err
	//}

	return nil
}

package pg

import (
	"context"
	"database/sql"

	log "github.com/sirupsen/logrus"
)

//NewRepository connects to the sql db
func NewRepository(db *sql.DB) Repository {
	return Repository{
		db: db,
	}
}

// project from event stream
type Repository struct {
	db *sql.DB
}

func (r Repository) store(ctx context.Context, user User) error {
	query := `INSERT INTO user (id, name) VALUES (?, ?) ON DUPLICATE KEY UPDATE id = ?, name = ?`
	_, err := r.db.ExecContext(
		ctx,
		query,
		user.ID,
		user.Name,
		user.ID, // start of upsert
		user.Name,
	)
	if err != nil {
		log.Error("Error while storing model")
		return err
	}

	return nil
}

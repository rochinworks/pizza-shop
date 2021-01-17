package pg

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

//NewRepository connects to the sql db
func NewRepository(db *sql.DB) repository {
	return repository{
		db: db,
	}
}

// project from event stream
type repository struct {
	db *sql.DB
}

func (r repository) Store(ctx context.Context, user User) (*uuid.UUID, error) {
	id := uuid.New()
	user.ID = id

	query := `INSERT INTO users (id, name) VALUES ($1, $2) ON CONFLICT ON CONSTRAINT users_pkey DO UPDATE SET name = $2 WHERE users.id = $1::varchar(36)`
	_, err := r.db.ExecContext(
		ctx,
		query,
		user.ID,
		user.Name,
	)
	if err != nil {
		log.Error("Error while storing user")
		return nil, err
	}

	return &id, nil
}

func (r repository) StoreOrder(ctx context.Context, pizza Pizza) (*uuid.UUID, error) {
	id := uuid.New()

	query := `INSERT INTO orders (id, style, status, userId) VALUES ($1, $2, $3, $4)`
	_, err := r.db.ExecContext(
		ctx,
		query,
		id,
		pizza.Style,
		pizza.Status,
		pizza.UserID,
	)
	if err != nil {
		log.Error("Error while storing order")
		return nil, err
	}

	return &id, nil
}

func (r repository) GetStatus(ctx context.Context, userID uuid.UUID, orderID uuid.UUID) (string, error) {
	query := `SELECT * FROM orders WHERE id = $1 AND userid = $2`
	rows, err := r.db.QueryContext(
		ctx,
		query,
		orderID,
		userID,
	)
	if err != nil {
		log.Error("Error while getting status")
		return "", err
	}
	var order Pizza
	for rows.Next() {
		err = rows.Scan(&order.OrderID, &order.Style, &order.Status, &order.UserID)
		if err != nil {
			return "", err
		}
	}

	return order.Status, nil
}

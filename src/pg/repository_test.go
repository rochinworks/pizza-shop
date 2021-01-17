package pg

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
)

func orderID() uuid.UUID {
	orderID := "5711a81a-163a-4fce-a019-b2f88f61562b"
	oID, _ := uuid.Parse(orderID)
	return oID
}

func userID() uuid.UUID {
	userID := "9d97a60e-586e-450a-8b2e-210a148db7c5"
	uID, _ := uuid.Parse(userID)
	return uID
}

func setupTest(query string) (context.Context, repository) {
	db, mock, _ := sqlmock.New()

	mock.ExpectQuery(query).
		WillReturnRows(sqlmock.NewRows([]string{"id", "style", "status", "userId"}).
			AddRow(orderID(), "hawaiian", "starting", userID()))
	ctx := context.Background()

	repo := NewRepository(db)

	return ctx, repo
}

func TestStore(t *testing.T) {
	db, mock, _ := sqlmock.New()
	query := `INSERT INTO users \(id, name\) VALUES \(\$1, \$2\) ON CONFLICT ON CONSTRAINT users_pkey DO UPDATE SET name = \$2 WHERE users.id = \$1\:\:varchar\(36\)`
	mock.ExpectExec(query).
		WithArgs(sqlmock.AnyArg(), "testing123").
		WillReturnResult(sqlmock.NewResult(1, 1))

	ctx := context.Background()
	user := User{
		ID:   userID(),
		Name: "testing123",
	}

	repo := NewRepository(db)
	_, err := repo.Store(ctx, user)
	if err != nil {
		t.Error("Could not store user correctly: ", err)
	}
}

func TestStoreOrder(t *testing.T) {
	db, mock, _ := sqlmock.New()
	query := `INSERT INTO orders \(id, style, status, userId\) VALUES \(\$1, \$2, \$3, \$4\)`
	mock.ExpectExec(query).
		WithArgs(sqlmock.AnyArg(), "hawaiian", "starting", sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	ctx := context.Background()
	pizza := Pizza{
		OrderID: orderID(),
		Style:   "hawaiian",
		UserID:  userID(),
		Status:  "starting",
	}

	repo := NewRepository(db)
	_, err := repo.StoreOrder(ctx, pizza)
	if err != nil {
		t.Error("Could not store pizza order correctly: ", err)
	}
}

func TestGetStatus(t *testing.T) {
	ctx, repo := setupTest(`SELECT \* FROM orders WHERE id = \$1 AND userid = \$2`)
	status, err := repo.GetStatus(ctx, userID(), orderID())
	if err != nil {
		t.Error("Could not get status from repo: ", err)
	}

	if status == "" {
		t.Error("no status returned")
	}

	expected := "starting"
	if status != expected {
		t.Error("wrong status")
	}
}

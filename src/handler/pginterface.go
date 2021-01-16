package handler

import (
	"context"
	"go-web-api/src/pg"

	"github.com/google/uuid"
)

//go:generate mockgen -source=pginterface.go -destination=mocks/pginterface.go -package=mocks
type Repository interface {
	Store(context.Context, pg.User) (*uuid.UUID, error)
	StoreOrder(context.Context, pg.Pizza) (*uuid.UUID, error)
	GetStatus(ctx context.Context, userID uuid.UUID, orderID uuid.UUID) (string, error)
}

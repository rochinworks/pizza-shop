package pg

import "github.com/google/uuid"

type User struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"username"`
}

type Pizza struct {
	OrderID uuid.UUID `json:"orderId"`
	Style   string    `json:"style"`
	UserID  uuid.UUID `json:"userId"`
	Status  string    `json:"status"`
}

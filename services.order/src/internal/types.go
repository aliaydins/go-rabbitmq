package order

import "github.com/google/uuid"

type CreateOrderRequest struct {
	Name string `json:"name"`
}

type UpdateOrderRequest struct {
	ID     uuid.UUID `json:"id"`
	Status int       `json:"status""`
}

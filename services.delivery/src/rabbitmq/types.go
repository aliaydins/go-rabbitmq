package rabbitmq

import "github.com/google/uuid"

type OrderEventType struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Status int       `json:"status"`
}

type DeliveryEventType struct {
	ID     uuid.UUID
	Done   bool
	Status int
}

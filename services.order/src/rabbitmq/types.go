package rabbitmq

import "github.com/google/uuid"

type DeliveryEventType struct {
	ID     uuid.UUID `json:"id"`
	Done   bool      `json:"done"`
	Status int       `json:"status"`
}

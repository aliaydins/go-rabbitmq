package event

import (
	"github.com/aliaydins/go-rabbitmq-example/services.order/src/entity"
	"github.com/google/uuid"
)

type OrderCreated struct {
	ID     uuid.UUID
	Name   string
	Status entity.Status
}

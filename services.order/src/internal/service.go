package order

import (
	"encoding/json"
	"github.com/aliaydins/go-rabbitmq-example/services.order/src/entity"
	"github.com/aliaydins/go-rabbitmq-example/services.order/src/event"
	"github.com/aliaydins/go-rabbitmq-example/services.order/src/rabbitmq"
	"github.com/google/uuid"
)

type Service struct {
	repo     *Repository
	rabbitmq *rabbitmq.RabbitMQ
}

func NewService(repo *Repository, rabbitmq *rabbitmq.RabbitMQ) *Service {
	return &Service{
		repo:     repo,
		rabbitmq: rabbitmq,
	}
}

func (s *Service) CreateOrder(req *CreateOrderRequest) (entity.Order, error) {

	order := entity.Order{
		ID:     uuid.New(),
		Name:   req.Name,
		Status: entity.Created,
	}

	createdOrder, err := s.repo.CreateOrder(order)
	if err != nil {
		return createdOrder, err
	}

	PublishOrderCreatedEvent(createdOrder, s.rabbitmq)

	return createdOrder, err
}

func (s *Service) UpdateOrder(req *UpdateOrderRequest) (*entity.Order, error) {

	order, err := s.repo.Update(req.ID, req.Status)
	if err != nil {
		return order, err
	}
	return order, err

}

func PublishOrderCreatedEvent(createdOrder entity.Order, r *rabbitmq.RabbitMQ) {
	orderCreatedEvent := event.OrderCreated{
		ID:     createdOrder.ID,
		Name:   createdOrder.Name,
		Status: createdOrder.Status,
	}

	payload, _ := json.Marshal(orderCreatedEvent)
	r.Publish(payload)
}

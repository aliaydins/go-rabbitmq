package delivery

import (
	"github.com/aliaydins/go-rabbitmq-example/services.delivery/src/entity"
	"github.com/google/uuid"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateDelivery(req *CreateDeliveryRequest) (entity.Delivery, error) {

	delivery := entity.Delivery{
		ID:   uuid.New(),
		Name: req.Name,
		Done: true,
	}

	createdOrder, err := s.repo.CreateDelivery(delivery)
	if err != nil {
		return createdOrder, err
	}

	return createdOrder, err
}

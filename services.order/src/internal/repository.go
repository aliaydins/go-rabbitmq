package order

import (
	"github.com/aliaydins/go-rabbitmq-example/services.order/src/entity"
	"github.com/aliaydins/go-rabbitmq-example/services.order/src/rabbitmq"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repository struct {
	db       *gorm.DB
	rabbitmq *rabbitmq.RabbitMQ
}

func NewRepository(db *gorm.DB, rabbitmq *rabbitmq.RabbitMQ) *Repository {
	db.Logger.LogMode(logger.Info)
	return &Repository{
		db:       db,
		rabbitmq: rabbitmq}
}

func (r *Repository) CreateOrder(order entity.Order) (entity.Order, error) {

	err := r.db.Model(&entity.Order{}).Create(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

func (r *Repository) FindById(id uuid.UUID) (*entity.Order, error) {
	order := new(entity.Order)
	err := r.db.Where("id = ?", id).First(&order).Error

	if err != nil {
		return nil, err
	}
	return order, nil
}

func (r *Repository) Update(id uuid.UUID, status int) (*entity.Order, error) {

	order, err := r.FindById(id)
	if err != nil {
		return order, err
	}

	err = r.db.Model(&order).Update("status", status).Error
	return order, err
}

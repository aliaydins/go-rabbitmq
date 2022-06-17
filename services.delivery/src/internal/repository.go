package delivery

import (
	"github.com/aliaydins/go-rabbitmq-example/services.delivery/src/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	db.Logger.LogMode(logger.Info)
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateDelivery(delivery entity.Delivery) (entity.Delivery, error) {

	err := r.db.Model(&entity.Delivery{}).Create(&delivery).Error
	if err != nil {
		return delivery, err
	}
	return delivery, nil
}

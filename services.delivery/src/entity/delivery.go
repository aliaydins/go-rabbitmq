package entity

import (
	"github.com/google/uuid"
	"time"
)

type Delivery struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name      string    `gorm:"not null"`
	Done      bool      `gorm:"default:false;"`
	CreatedAt time.Time `gorm:"not null default CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}

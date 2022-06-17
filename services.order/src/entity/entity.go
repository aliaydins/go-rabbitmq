package entity

import (
	"time"

	"github.com/google/uuid"
)

type Status int

const (
	Created   Status = 0
	Completed Status = 1
	Failed    Status = 2
)

type Order struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name      string    `gorm:"not null"`
	Status    Status    `gorm:"not null;"`
	CreatedAt time.Time `gorm:"not null default CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}

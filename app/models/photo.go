package models

import (
	"time"

	"github.com/google/uuid"
)

type Photo struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Title     string    `gorm:"not null"`
	PhotoUrl  string    `gorm:"not null"`
	UserID    uuid.UUID `gorm:"not null;type:uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

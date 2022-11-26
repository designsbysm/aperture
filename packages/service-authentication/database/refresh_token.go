package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshToken struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *gorm.DeletedAt
	UserID    uuid.UUID
}

func (t *RefreshToken) Create() error {
	return DB.Create(&t).Error
}

func (t *RefreshToken) Delete() error {
	return DB.Delete(&t, t).Error
}

func (t *RefreshToken) Read() error {
	return DB.First(&t, t).Error
}

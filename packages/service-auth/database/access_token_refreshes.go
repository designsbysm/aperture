package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccessTokenRefreshes struct {
	ID           uint64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *gorm.DeletedAt
	RefreshToken uuid.UUID
}

func (t *AccessTokenRefreshes) Create() error {
	return DB.Create(&t).Error
}

func (t *AccessTokenRefreshes) Delete() error {
	return DB.Delete(&t, t).Error
}

func (t *AccessTokenRefreshes) List() ([]RefreshToken, error) {
	list := []RefreshToken{}
	err := DB.Find(&list, t).Error

	return list, err
}

func (t *AccessTokenRefreshes) Read() error {
	return DB.First(&t, t).Error
}

func (t *AccessTokenRefreshes) Update() error {
	return DB.Save(&t).Error
}

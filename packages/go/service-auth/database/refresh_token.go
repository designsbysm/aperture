package database

import (
	"time"

	"github.com/gofrs/uuid"
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

func (t *RefreshToken) DeleteAll(userID uuid.UUID) error {
	return DB.Where("user_id = ?", userID.String()).Delete(&t).Error
}

func (t *RefreshToken) Delete() error {
	return DB.Delete(&t, t).Error
}

func (t *RefreshToken) List() ([]RefreshToken, error) {
	list := []RefreshToken{}
	err := DB.Find(&list, t).Error

	return list, err
}

func (t *RefreshToken) Read() error {
	return DB.First(&t, t).Error
}

func (t *RefreshToken) Update() error {
	return DB.Save(&t).Error
}

package database

import (
	"time"

	"aperture/go-types/userrole"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Role struct {
	ID        uint64          `json:"id"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	DeletedAt *gorm.DeletedAt `json:"deletedAt,omitempty"`
	UserID    uuid.UUID       `json:"userID"`
	Role      userrole.T      `json:"role"`
}

func (t *Role) Create() error {
	return DB.Create(&t).Error
}

func (t *Role) Delete() error {
	return DB.Delete(&t, t).Error
}

func (t *Role) List() ([]Role, error) {
	list := []Role{}
	err := DB.Find(&list, t).Error

	return list, err
}

func (t *Role) Read() error {
	return DB.First(&t, t).Error
}

func (t *Role) Update() error {
	return DB.Save(&t).Error
}

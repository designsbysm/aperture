package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	ID        uint            `json:"id"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	DeletedAt *gorm.DeletedAt `json:"deletedAt,omitempty"`
	UserID    uuid.UUID       `json:"userID"`
	Role      RoleEnum        `json:"role"`
}

func (r *Role) Create() error {
	return DB.Create(&r).Error
}

func (r *Role) Delete() error {
	return DB.Delete(&r, r).Error
}

func (r *Role) List() ([]Role, error) {
	list := []Role{}
	err := DB.Find(&list, r).Error

	return list, err
}

func (r *Role) Read() error {
	return DB.First(&r, r).Error
}

func (r *Role) Update() error {
	return DB.Save(&r).Error
}

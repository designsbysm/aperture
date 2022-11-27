package database

import (
	"time"

	"gorm.io/gorm"
)

type Log struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *gorm.DeletedAt
	Service   string
	Level     string
	Message   string
}

func (t *Log) Create() error {
	return DB.Create(&t).Error
}

func (t *Log) Delete() error {
	return DB.Delete(&t, t).Error
}

func (t *Log) List() ([]Log, error) {
	list := []Log{}
	err := DB.Find(&list, t).Error

	return list, err
}

func (t *Log) Read() error {
	return DB.First(&t, t).Error
}

func (t *Log) Update() error {
	return DB.Save(&t).Error
}

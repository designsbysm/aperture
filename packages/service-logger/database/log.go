package database

import (
	"time"

	"github.com/smaperture/go-types/loggerlevel"
	"github.com/smaperture/go-types/loggerservice"
	"gorm.io/gorm"
)

type Log struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *gorm.DeletedAt
	Service   loggerservice.T
	Level     loggerlevel.T
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

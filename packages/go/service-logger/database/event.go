package database

import (
	"time"

	"aperture/types/loggerlevel"
	"aperture/types/loggerservice"

	"gorm.io/gorm"
)

type Event struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *gorm.DeletedAt
	Service   loggerservice.T
	Level     loggerlevel.T
	Message   string
}

func (e *Event) Create() error {
	return DB.Create(&e).Error
}

func (e *Event) Delete() error {
	return DB.Delete(&e, e).Error
}

func (e *Event) List() ([]Event, error) {
	list := []Event{}
	err := DB.Find(&list, e).Error

	return list, err
}

func (e *Event) Read() error {
	return DB.First(&e, e).Error
}

func (e *Event) Update() error {
	return DB.Save(&e).Error
}

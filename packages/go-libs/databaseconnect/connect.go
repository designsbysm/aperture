package databaseconnect

import (
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Go(connection string, log int) (db *gorm.DB, err error) {
	if connection == "" {
		return nil, errors.New("connection string is missing")
	}

	return gorm.Open(
		postgres.Open(connection),
		&gorm.Config{
			Logger: logger(log),
		})
}

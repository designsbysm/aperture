package databaseconnect

import (
	"errors"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Go(connection string, schemaName string, logLevel int) (db *gorm.DB, err error) {
	if connection == "" {
		return nil, errors.New("connection string is missing")
	}

	return gorm.Open(
		postgres.Open(connection),
		&gorm.Config{
			Logger: logger(logLevel),
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: fmt.Sprintf("%s.", schemaName),
			},
		})
}

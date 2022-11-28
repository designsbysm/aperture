package databaseconnect

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Go(schemaName string) (db *gorm.DB, err error) {
	connection := viper.GetString("DATABASE_URL")
	logLevel := viper.GetInt("gorm.level")

	return gorm.Open(
		postgres.Open(connection),
		&gorm.Config{
			Logger: logger(logLevel),
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: fmt.Sprintf("%s.", schemaName),
			},
		})
}

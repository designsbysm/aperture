package databaseconnect

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Go(schemaName string) (db *gorm.DB, err error) {
	connection := viper.GetString("URL_DATABASE")
	if connection == "" {
		user := viper.GetString("DATABASE_USER")
		password := viper.GetString("DATABASE_PASSWORD")
		host := viper.GetString("URL_DOCKER")
		port := viper.GetString("PORT_DATABASE")
		name := viper.GetString("DATABASE_NAME")

		connection = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, name)
	}

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

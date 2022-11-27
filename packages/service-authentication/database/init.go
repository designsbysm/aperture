package database

import (
	"fmt"

	"github.com/smaperture/go-libs/databaseconnect"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() (err error) {
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

	DB, err = databaseconnect.Go(connection, "authentication", logLevel)
	if err != nil {
		return err
	}

	return nil
}

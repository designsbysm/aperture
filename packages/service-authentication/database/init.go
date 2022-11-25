package database

import (
	"fmt"

	"github.com/smaperture/golibs/databaseconnect"

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

	level := viper.GetInt("gorm.level")

	DB, err = databaseconnect.Go(connection, level)
	if err != nil {
		return err
	}

	return nil
}

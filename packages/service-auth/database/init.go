package database

import (
	"aperture/go-libs/databaseconnect"

	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() (err error) {
	DB, err = databaseconnect.Go("authentication")

	return err
}

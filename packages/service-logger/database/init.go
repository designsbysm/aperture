package database

import (
	"github.com/smaperture/go-libs/databaseconnect"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() (err error) {
	DB, err = databaseconnect.Go("logger")

	return err
}

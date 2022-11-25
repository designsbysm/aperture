package databaseconnect

import (
	"log"
	"os"
	"time"

	glog "gorm.io/gorm/logger"
)

func logger(level int) glog.Interface {
	return glog.New(
		log.New(os.Stdout, "\r\n", 0),
		glog.Config{
			Colorful:                  true,
			IgnoreRecordNotFoundError: true,
			LogLevel:                  glog.LogLevel(level),
			SlowThreshold:             200 * time.Millisecond,
		},
	)
}

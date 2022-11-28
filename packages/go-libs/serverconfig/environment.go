package serverconfig

import (
	"errors"
	"io/fs"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Environment() error {
	// load files
	if err := loadConfig("./config.yaml", true); err != nil {
		return err
	}

	if err := loadConfig("../../.env", false); err != nil {
		return err
	}

	if err := loadConfig(".env", false); err != nil {
		return err
	}

	// load env variables
	err := viper.BindEnv("DATABASE_URL")
	if err != nil {
		return err
	}

	err = viper.BindEnv("DOMAIN")
	if err != nil {
		return err
	}

	// setup stuff
	if viper.GetBool("gin.release") {
		gin.SetMode(gin.ReleaseMode)
	}

	return nil
}

func loadConfig(file string, required bool) error {
	if file == "" {
		return errors.New("config file path required")
	}

	viper.SetConfigFile(file)

	if err := viper.MergeInConfig(); err != nil {
		_, missing := err.(*fs.PathError)

		if (missing && required) || !missing {
			return err
		}
	}

	return nil
}

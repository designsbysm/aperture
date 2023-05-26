package api

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/designsbysm/timber/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Server() {
	router := gin.New()
	addCORS(router)
	AddRoute(router)

	port := viper.GetString("PORT_API")

	go func() {
		if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
			timber.Error("API:", err)
		}
	}()

	if viper.GetBool("gin.release") {
		timber.Info(fmt.Sprintf("API: listening on :%s", port))
	}

	// wait for ^c
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	fmt.Println("")
	timber.Info("API: closed")
}

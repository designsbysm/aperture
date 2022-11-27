package rpc

import (
	"context"

	"github.com/smaperture/go-libs/rpcclient"
	"github.com/smaperture/go-types/loggerlevel"
	"github.com/smaperture/go-types/loggerservice"
	"github.com/smaperture/service-authentication/loggerpb"
	"github.com/spf13/viper"
)

func LogEvent(level loggerlevel.T, message string) {
	host := viper.GetString("URL_DOCKER")
	port := viper.GetString("PORT_SERVICE_LOGGER")

	connection, err := rpcclient.Run(host, port)
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	client := loggerpb.NewLoggerServiceClient(connection)

	req := loggerpb.LogRequest{
		Service: string(loggerservice.Authentication),
		Level:   string(level),
		Message: message,
	}

	if _, err := client.Log(context.Background(), &req); err != nil {
		panic(err)
	}
}

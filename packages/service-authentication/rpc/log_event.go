package rpc

import (
	"context"

	"github.com/smaperture/golibs/rpcclient"
	"github.com/smaperture/service-authentication/loggerpb"
	"github.com/spf13/viper"
)

func LogEvent(level string, message string) {
	// host := viper.GetString("rpc.host")
	host := viper.GetString("URL_DOCKER")
	port := viper.GetString("PORT_SERVICE_LOGGER")

	connection, err := rpcclient.Run(host, port)
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	client := loggerpb.NewLoggerServiceClient(connection)

	req := loggerpb.LogRequest{
		Service: "authentication",
		Level:   level,
		Message: message,
	}

	if _, err := client.Log(context.Background(), &req); err != nil {
		panic(err)
	}
}

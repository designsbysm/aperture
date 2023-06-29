package rpc

import (
	"context"

	"aperture/libs/rpcclient"
	"aperture/service-auth/loggerpb"
	"aperture/types/loggerlevel"
	"aperture/types/loggerservice"

	"github.com/spf13/viper"
)

func LogEvent(level loggerlevel.T, message string) {
	host := viper.GetString("DOMAIN")
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

	if _, err := client.LogEvent(context.Background(), &req); err != nil {
		panic(err)
	}
}

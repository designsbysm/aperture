package rpc

import (
	"context"

	"aperture/go-types/loggerlevel"
	"aperture/go-types/loggerservice"
	"aperture/service-logger/database"
	"aperture/service-logger/loggerpb"
)

func (*server) Log(ctx context.Context, in *loggerpb.LogRequest) (*loggerpb.LogResponse, error) {
	service := in.GetService()
	level := in.GetLevel()
	message := in.GetMessage()

	log := database.Log{
		Service: loggerservice.FromString(service),
		Level:   loggerlevel.FromString(level),
		Message: message,
	}

	err := log.Create()
	res := loggerpb.LogResponse{}

	return &res, err
}

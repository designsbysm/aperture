package rpc

import (
	"context"

	"github.com/smaperture/go-types/loggerlevel"
	"github.com/smaperture/go-types/loggerservice"
	"github.com/smaperture/service-logger/database"
	"github.com/smaperture/service-logger/loggerpb"
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

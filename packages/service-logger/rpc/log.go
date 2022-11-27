package rpc

import (
	"context"

	"github.com/designsbysm/timber/v2"
	"github.com/smaperture/service-logger/database"
	"github.com/smaperture/service-logger/loggerpb"
)

func (*server) Log(ctx context.Context, in *loggerpb.LogRequest) (*loggerpb.LogResponse, error) {
	service := in.GetService()
	level := in.GetLevel()
	message := in.GetMessage()

	timber.Debug(service, level, message)

	log := database.Log{
		Service: service,
		Level:   level,
		Message: message,
	}

	err := log.Create()
	res := loggerpb.LogResponse{}

	return &res, err
}

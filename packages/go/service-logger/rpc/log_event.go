package rpc

import (
	"context"

	"aperture/libs/loggerpb"
	"aperture/service-logger/database"
	"aperture/types/loggerlevel"
	"aperture/types/loggerservice"
)

func (*server) LogEvent(ctx context.Context, in *loggerpb.LogRequest) (*loggerpb.LogResponse, error) {
	service := in.GetService()
	level := in.GetLevel()
	message := in.GetMessage()

	event := database.Event{
		Service: loggerservice.FromString(service),
		Level:   loggerlevel.FromString(level),
		Message: message,
	}

	err := event.Create()
	res := loggerpb.LogResponse{}

	return &res, err
}

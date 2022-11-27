package rpc

import (
	"github.com/smaperture/go-libs/rpcserver"
	"github.com/smaperture/service-logger/loggerpb"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type server struct {
	loggerpb.UnimplementedLoggerServiceServer
}

func Server() error {
	port := viper.GetString("PORT_SERVICE_LOGGER")

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	loggerpb.RegisterLoggerServiceServer(s, &server{})

	return rpcserver.Run(s, "", port)
}
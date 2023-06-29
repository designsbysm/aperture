package rpc

import (
	"aperture/libs/authenticationpb"
	"aperture/libs/rpcserver"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type server struct {
	authenticationpb.UnimplementedAuthenticationServiceServer
}

func Server() error {
	port := viper.GetString("PORT_SERVICE_AUTHENTICATION")

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	authenticationpb.RegisterAuthenticationServiceServer(s, &server{})

	return rpcserver.Run(s, "", port)
}

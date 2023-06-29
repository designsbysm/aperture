package rpc

import (
	"context"

	"aperture/gateway/authenticationpb"
	"aperture/go-libs/rpcclient"
	"aperture/go-types/emailaddress"

	"github.com/spf13/viper"
)

func Login(username emailaddress.T, password string, longLivedToken string) (*authenticationpb.LoginResponse, error) {
	host := viper.GetString("DOMAIN")
	port := viper.GetString("PORT_SERVICE_AUTHENTICATION")

	connection, err := rpcclient.Run(host, port)
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	client := authenticationpb.NewAuthenticationServiceClient(connection)

	req := authenticationpb.LoginRequest{
		Username:       string(username),
		Password:       password,
		LongLivedToken: longLivedToken,
	}

	return client.Login(context.Background(), &req)
}

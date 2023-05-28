package rpc

import (
	"context"

	"aperture/go-libs/rpcclient"
	"aperture/go-types/emailaddress"
	"aperture/rest-api/authenticationpb"

	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type authenticationResponse struct {
	// [ ] make AccessToken type
	AccessToken uuid.UUID `json:"accessToken"`
	// [ ] make RefreshToken type
	RefreshToken uuid.UUID `json:"refreshToken"`
	Expiration   int32     `json:"expiration"`
	// TODO: add user info
}

func Login(username emailaddress.T, password string) (authenticationResponse, error) {
	host := viper.GetString("DOMAIN")
	port := viper.GetString("PORT_SERVICE_AUTHENTICATION")

	connection, err := rpcclient.Run(host, port)
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	client := authenticationpb.NewAuthenticationServiceClient(connection)

	req := authenticationpb.LoginRequest{
		Username: string(username),
		Password: string(password),
	}

	result := authenticationResponse{}

	if res, err := client.Login(context.Background(), &req); err != nil {
		return result, err
	} else {
		accessToken, err := uuid.Parse(res.AccessToken)
		if err != nil {
			return result, err
		}

		refreshToken, err := uuid.Parse(res.RefreshToken)
		if err != nil {
			return result, err
		}

		result.AccessToken = accessToken
		result.RefreshToken = refreshToken
		result.Expiration = res.Expiration

		return result, nil
	}
}

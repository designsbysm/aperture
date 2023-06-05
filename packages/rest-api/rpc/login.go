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
	// [ ] make UserID type
	// [ ] make AccessToken type
	// [ ] make RefreshToken type
	AccessToken  string    `json:"accessToken"`
	RefreshToken uuid.UUID `json:"refreshToken"`
}

func Login(username emailaddress.T, password string) (*authenticationpb.LoginResponse, error) {
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

}

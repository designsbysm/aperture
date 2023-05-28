package rpc

import (
	"context"
	"errors"

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
	UserID       uuid.UUID `json:"-"`
	AccessToken  string    `json:"accessToken"`
	RefreshToken uuid.UUID `json:"refreshToken"`
	Expiration   int32     `json:"expiration"`
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

	if res, err := client.Login(context.Background(), &req); err != nil {
		return authenticationResponse{}, err
	} else {
		userID, err := uuid.Parse(res.UserID)
		if err != nil {
			return authenticationResponse{}, err
		}

		if res.AccessToken == "" {
			return authenticationResponse{}, errors.New("missing accessToken")
		}
		// TODO: check jwt validity

		refreshToken, err := uuid.Parse(res.RefreshToken)
		if err != nil {
			return authenticationResponse{}, err
		}

		result := authenticationResponse{
			UserID:       userID,
			AccessToken:  res.AccessToken,
			RefreshToken: refreshToken,
			Expiration:   res.Expiration,
		}

		return result, nil
	}
}

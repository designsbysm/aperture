package rpc

import (
	"context"
	"errors"

	"aperture/service-authentication/authenticationpb"

	"github.com/google/uuid"
)

func (*server) Login(ctx context.Context, in *authenticationpb.LoginRequest) (*authenticationpb.LoginResponse, error) {
	username := in.GetUsername()
	password := in.GetPassword()

	res := authenticationpb.LoginResponse{}

	if username == "" || password == "" {
		return &res, errors.New("invalid login")
	}

	//TODO: complete login

	//TODO: change access token to JWT
	res.AccessToken = uuid.New().String()
	res.RefreshToken = uuid.New().String()
	res.Expiration = 1800
	// TODO: add user info

	return &res, nil
}

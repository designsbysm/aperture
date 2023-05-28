package rpc

import (
	"context"
	"errors"

	"aperture/service-auth/authenticationpb"

	"github.com/google/uuid"
)

//TODO: add logging

func (*server) Login(ctx context.Context, in *authenticationpb.LoginRequest) (*authenticationpb.LoginResponse, error) {
	username := in.GetUsername()
	password := in.GetPassword()

	if username == "" || password == "" {
		return &authenticationpb.LoginResponse{}, errors.New("invalid login")
	}

	//TODO: complete login

	res := authenticationpb.LoginResponse{
		UserID:       uuid.New().String(),
		AccessToken:  "jwt",
		RefreshToken: uuid.New().String(),
		Expiration:   1800,
	}

	return &res, nil
}

package rpc

import (
	"context"
	"errors"

	"aperture/go-types/emailaddress"
	"aperture/go-types/jwt"
	"aperture/service-auth/authenticationpb"
	"aperture/service-auth/database"

	"github.com/gofrs/uuid"
)

//TODO: add logging

func (*server) Login(ctx context.Context, in *authenticationpb.LoginRequest) (*authenticationpb.LoginResponse, error) {
	username := in.GetUsername()
	password := in.GetPassword()

	if username == "" || password == "" {
		return &authenticationpb.LoginResponse{}, errors.New("invalid login")
	}

	user := database.User{
		Email: emailaddress.T(username),
	}
	if err := user.Read(); err != nil {
		return &authenticationpb.LoginResponse{}, errors.New("user not found")
	}

	if err := user.PasswordValidate(password); err != nil {
		return &authenticationpb.LoginResponse{}, errors.New("invalid password")
	}

	role := database.Role{
		UserID: user.ID,
	}
	if err := role.Read(); err != nil {
		return &authenticationpb.LoginResponse{}, errors.New("role not found")
	}

	// TODO: long live tokens
	accessToken, err := jwt.Encode(user.ID, role.Role, user.FirstName, user.LastName, false)
	if err != nil {
		return &authenticationpb.LoginResponse{}, errors.New("unable to create accessToken")
	}

	token, err := uuid.NewV7()
	if err != nil {
		return &authenticationpb.LoginResponse{}, errors.New("unable to create uuid")
	}

	refreshToken := database.RefreshToken{
		ID:     token,
		UserID: user.ID,
	}
	if err := refreshToken.Create(); err != nil {
		return &authenticationpb.LoginResponse{}, errors.New("unable to create refreshToken")
	}

	res := authenticationpb.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken.ID.String(),
	}

	return &res, nil
}

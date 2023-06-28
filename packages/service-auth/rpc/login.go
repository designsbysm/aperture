package rpc

import (
	"context"
	"errors"
	"fmt"

	"aperture/go-types/emailaddress"
	"aperture/go-types/jwt"
	"aperture/go-types/loggerlevel"
	"aperture/service-auth/authenticationpb"
	"aperture/service-auth/database"

	"github.com/gofrs/uuid"
	"github.com/spf13/viper"
)

func (*server) Login(ctx context.Context, in *authenticationpb.LoginRequest) (*authenticationpb.LoginResponse, error) {
	username := in.GetUsername()
	password := in.GetPassword()

	if username == "" || password == "" {
		msg := "invalid login"
		LogEvent(loggerlevel.Error, msg)
		return &authenticationpb.LoginResponse{}, errors.New(msg)
	}

	user := database.User{
		Email: emailaddress.T(username),
	}
	if err := user.Read(); err != nil {
		msg := fmt.Sprintf("user not found: %s", username)
		LogEvent(loggerlevel.Error, msg)
		return &authenticationpb.LoginResponse{}, errors.New(msg)
	}

	if err := user.PasswordValidate(password); err != nil {
		msg := "invalid password"
		LogEvent(loggerlevel.Error, msg)
		return &authenticationpb.LoginResponse{}, errors.New(msg)
	}

	role := database.Role{
		UserID: user.ID,
	}
	if err := role.Read(); err != nil {
		msg := fmt.Sprintf("role not found for %s:", user.ID)
		LogEvent(loggerlevel.Error, msg)
		return &authenticationpb.LoginResponse{}, errors.New(msg)
	}

	userToken := in.GetLongLivedToken()
	secretToken := viper.GetString("jwt.longLived.secret")
	isLongLived := userToken != "" && secretToken != "" && userToken == secretToken
	accessToken, err := jwt.Encode(user.ID, role.Role, user.FirstName, user.LastName, isLongLived)
	if err != nil {
		msg := "unable to create accessToken"
		LogEvent(loggerlevel.Error, msg)
		return &authenticationpb.LoginResponse{}, errors.New(msg)
	} else if isLongLived {
		msg := fmt.Sprintf("long lived jwt created for %s:", user.ID)
		LogEvent(loggerlevel.Info, msg)
	}

	token, err := uuid.NewV7()
	if err != nil {
		msg := "unable to create uuid"
		LogEvent(loggerlevel.Error, msg)
		return &authenticationpb.LoginResponse{}, errors.New(msg)
	}

	refreshToken := database.RefreshToken{
		ID:     token,
		UserID: user.ID,
	}
	if err := refreshToken.Create(); err != nil {
		msg := fmt.Sprintf("unable to create refreshToken for %s:", user.ID)
		LogEvent(loggerlevel.Error, msg)
		return &authenticationpb.LoginResponse{}, errors.New(msg)
	}

	res := authenticationpb.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken.ID.String(),
	}

	msg := fmt.Sprintf("login successful: %s", user.ID)
	LogEvent(loggerlevel.Info, msg)
	return &res, nil
}

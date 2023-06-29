package rpc

import (
	"context"
	"errors"
	"fmt"

	"aperture/service-auth/authenticationpb"
	"aperture/service-auth/database"
	"aperture/types/emailaddress"
	"aperture/types/jwt"

	"github.com/designsbysm/timber/v2"
	"github.com/gofrs/uuid"
	"github.com/spf13/viper"
)

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
		msg := fmt.Sprintf("user not found: %s", username)
		return &authenticationpb.LoginResponse{}, errors.New(msg)
	}

	if err := user.PasswordValidate(password); err != nil {
		return &authenticationpb.LoginResponse{}, errors.New("invalid password")
	}

	role := database.Role{
		UserID: user.ID,
	}
	if err := role.Read(); err != nil {
		msg := fmt.Sprintf("role not found for: %s", user.ID)
		return &authenticationpb.LoginResponse{}, errors.New(msg)
	}

	userToken := in.GetLongLivedToken()
	secretToken := viper.GetString("jwt.longLived.secret")
	isLongLived := userToken != "" && secretToken != "" && userToken == secretToken
	accessToken, err := jwt.Encode(user.ID, role.Role, user.FirstName, user.LastName, isLongLived)
	if err != nil {
		return &authenticationpb.LoginResponse{}, errors.New("unable to create accessToken")
	} else if isLongLived {
		msg := fmt.Sprintf("long lived jwt created for: %s", user.ID)
		timber.Info(msg)
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
		msg := fmt.Sprintf("unable to create refreshToken for: %s", user.ID)
		return &authenticationpb.LoginResponse{}, errors.New(msg)
	}

	res := authenticationpb.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken.ID.String(),
	}

	msg := fmt.Sprintf("complete login rpc: %s", user.ID)
	timber.Info(msg)
	return &res, nil
}

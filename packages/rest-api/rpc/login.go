package rpc

import (
	"context"

	"aperture/rest-api/authenticationpb"

	"github.com/designsbysm/timber/v2"
)

func (*server) Login(ctx context.Context, in *authenticationpb.LoginRequest) (*authenticationpb.LoginResponse, error) {
	username := in.GetUsername()
	password := in.GetPassword()

	// fmt.Println(username, password)

	timber.Debug(username, password)

	// return collatz.Hailstone(seed)
	res := authenticationpb.LoginResponse{
		Token:      "test",
		Expiration: 1800,
	}

	return &res, nil
}

// syntax = "proto3";

// package authentication;
// option go_package = "/authenticationpb";

// message AuthorizeRequest {
//   string token = 1;
//   string role = 2;
// }

// message AuthorizeResponse {
//   bool allow = 1;
// }

// message LogoutRequest {
//   string token = 1;
// }

// message LogoutResponse {}

// service AuthenticationService {
//   rpc Authorize(AuthorizeRequest) returns (AuthorizeResponse) {};
//   rpc Login(LoginRequest) returns (LoginResponse) {};
//   rpc Logout(LogoutRequest) returns (LogoutResponse) {};
// }

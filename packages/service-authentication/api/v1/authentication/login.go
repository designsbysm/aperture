package authentication

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smaperture/go-types/emailaddress"
	"github.com/smaperture/go-types/jwt"
	"github.com/smaperture/go-types/loggerlevel"
	"github.com/smaperture/service-authentication/database"
	"github.com/smaperture/service-authentication/rpc"
	"github.com/spf13/viper"
)

type LoginRequest struct {
	Email    emailaddress.T
	Password string
}

func login(c *gin.Context) {
	req := LoginRequest{}
	if err := c.BindJSON(&req); err != nil {
		//nolint:errcheck
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if req.Email == "" || req.Password == "" {
		c.Status(http.StatusUnauthorized)
		return
	}

	user := database.User{
		Email: req.Email,
	}
	if err := user.Read(); err != nil {
		rpc.LogEvent(loggerlevel.Warn, fmt.Sprintf("unknown email: %s", req.Email))
		c.Status(http.StatusUnauthorized)
		return
	}

	if err := user.PasswordValidate(req.Password); err != nil {
		rpc.LogEvent(loggerlevel.Warn, fmt.Sprintf("wrong password for user: %s", user.ID))
		c.Status(http.StatusUnauthorized)
		return
	}

	role := database.Role{
		UserID: user.ID,
	}
	if err := user.Read(); err != nil {
		//nolint:errcheck
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	header := c.GetHeader("LongLived")
	secretKey := viper.GetString("jwt.longLived.secret")
	accessToken, err := jwt.Encode(user.ID, role.Role, header == secretKey)
	if err != nil {
		//nolint:errcheck
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	refreshToken := database.RefreshToken{
		UserID: user.ID,
	}
	if err := refreshToken.Create(); err != nil {
		//nolint:errcheck
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	res := AuthenticationResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken.ID,
	}

	rpc.LogEvent(loggerlevel.Info, fmt.Sprintf("user logged in: %s", user.ID))
	c.JSON(http.StatusOK, res)
}

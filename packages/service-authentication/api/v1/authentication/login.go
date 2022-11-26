package authentication

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smaperture/service-authentication/database"
	"github.com/smaperture/service-authentication/jwt"
	"github.com/spf13/viper"
)

type LoginRequest struct {
	Email    string
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
		//nolint:errcheck
		c.AbortWithError(http.StatusInternalServerError, err)
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

	c.JSON(http.StatusOK, res)
}

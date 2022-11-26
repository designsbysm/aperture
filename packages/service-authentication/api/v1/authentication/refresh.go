package authentication

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/smaperture/service-authentication/database"
	"github.com/smaperture/service-authentication/types/jwt"
	"github.com/spf13/viper"
)

func refreshAccess(c *gin.Context) {
	token, err := uuid.Parse(c.Param("token"))
	if err != nil {
		//nolint:errcheck
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	refreshToken := database.RefreshToken{
		ID: token,
	}
	if err := refreshToken.Read(); err != nil {
		//nolint:errcheck
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user := database.User{
		ID: refreshToken.UserID,
	}
	if err := user.Read(); err != nil {
		//nolint:errcheck
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	role := database.Role{
		UserID: user.ID,
	}
	if err := role.Read(); err != nil {
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

	res := AuthenticationResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken.ID,
	}

	c.JSON(http.StatusOK, res)
}

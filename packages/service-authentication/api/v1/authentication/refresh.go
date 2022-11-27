package authentication

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/smaperture/go-types/jwt"
	"github.com/smaperture/go-types/loggerlevel"
	"github.com/smaperture/go-types/userrole"
	"github.com/smaperture/service-authentication/database"
	"github.com/smaperture/service-authentication/rpc"
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

	userRole := database.Role{
		UserID: user.ID,
	}
	if err := userRole.Read(); err != nil {
		//nolint:errcheck
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if userRole.Role == userrole.Diabled {
		//nolint:errcheck
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	header := c.GetHeader("LongLived")
	secretKey := viper.GetString("jwt.longLived.secret")
	accessToken, err := jwt.Encode(user.ID, userRole.Role, header == secretKey)
	if err != nil {
		//nolint:errcheck
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	accessTokenRefreshes := database.AccessTokenRefreshes{
		RefreshToken: refreshToken.ID,
	}
	if err := accessTokenRefreshes.Create(); err != nil {
		//nolint:errcheck
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res := AuthenticationResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken.ID,
	}

	rpc.LogEvent(loggerlevel.Info, fmt.Sprintf("access token refeshed: %s", refreshToken.ID))
	c.JSON(http.StatusOK, res)
}

package auth

// import (
// 	"fmt"
// 	"net/http"

// 	"aperture/go-types/jwt"
// 	"aperture/go-types/loggerlevel"
// 	"aperture/go-types/userrole"
// 	"aperture/service-auth/database"
// 	"aperture/service-auth/rpc"

// 	"github.com/gin-gonic/gin"
// 	"github.com/google/uuid"
// 	"github.com/spf13/viper"
// )

// func refreshAccess(c *gin.Context) {
// 	token, err := uuid.Parse(c.Param("token"))
// 	if err != nil {
// 		//nolint:errcheck
// 		c.AbortWithStatus(http.StatusUnauthorized)
// 		return
// 	}

// 	refreshToken := database.RefreshToken{
// 		ID: token,
// 	}
// 	if err := refreshToken.Read(); err != nil {
// 		//nolint:errcheck
// 		c.AbortWithStatus(http.StatusUnauthorized)
// 		return
// 	}

// 	user := database.User{
// 		ID: refreshToken.UserID,
// 	}
// 	if err := user.Read(); err != nil {
// 		//nolint:errcheck
// 		c.AbortWithStatus(http.StatusUnauthorized)
// 		return
// 	}

// 	userRole := database.Role{
// 		UserID: user.ID,
// 	}
// 	if err := userRole.Read(); err != nil {
// 		//nolint:errcheck
// 		c.AbortWithError(http.StatusInternalServerError, err)
// 		return
// 	}
// 	if userRole.Role == userrole.Diabled {
// 		//nolint:errcheck
// 		c.AbortWithStatus(http.StatusUnauthorized)
// 		return
// 	}

// 	header := c.GetHeader("LongLived")
// 	secretKey := viper.GetString("jwt.longLived.secret")
// 	accessToken, err := jwt.Encode(user.ID, userRole.Role, header == secretKey)
// 	if err != nil {
// 		//nolint:errcheck
// 		c.AbortWithError(http.StatusInternalServerError, err)
// 		return
// 	}

// 	accessTokenRefreshes := database.AccessTokenRefreshes{
// 		RefreshToken: refreshToken.ID,
// 	}
// 	if err := accessTokenRefreshes.Create(); err != nil {
// 		//nolint:errcheck
// 		c.AbortWithStatus(http.StatusUnauthorized)
// 		return
// 	}

// 	res := AuthenticationResponse{
// 		AccessToken:  accessToken,
// 		RefreshToken: refreshToken.ID,
// 	}

// 	rpc.LogEvent(loggerlevel.Info, fmt.Sprintf("access token refeshed: %s", refreshToken.ID))
// 	c.JSON(http.StatusOK, res)
// }

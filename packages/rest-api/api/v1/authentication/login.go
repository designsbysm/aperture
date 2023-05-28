package authentication

import (
	"fmt"
	"net/http"

	"aperture/go-types/emailaddress"
	"aperture/go-types/loggerlevel"

	// "aperture/go-types/jwt"

	// "aperture/service-authentication/database"
	"aperture/rest-api/rpc"

	"github.com/gin-gonic/gin"
	// "github.com/spf13/viper"
)

type LoginRequest struct {
	Username emailaddress.T
	Password string
}

func login(c *gin.Context) {
	req := LoginRequest{}
	if err := c.BindJSON(&req); err != nil {
		//nolint:errcheck
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	auth, err := rpc.Login(req.Username, req.Password)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		rpc.LogEvent(loggerlevel.Warn, fmt.Sprintf("invalid login attempt: %s", req.Username))
		return
	}

	rpc.LogEvent(loggerlevel.Info, fmt.Sprintf("user logged in: %s", req.Username))
	c.JSON(http.StatusOK, auth)
}

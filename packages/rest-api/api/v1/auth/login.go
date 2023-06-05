package auth

import (
	"fmt"
	"net/http"

	"aperture/go-types/emailaddress"
	"aperture/go-types/loggerlevel"
	"aperture/rest-api/rpc"

	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Username emailaddress.T
	Password string
}

func login(c *gin.Context) {
	req := loginRequest{}
	if err := c.BindJSON(&req); err != nil {
		//nolint:errcheck
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	auth, err := rpc.Login(req.Username, req.Password)
	if err != nil {
		username := req.Username
		if username == "" {
			username = "<empty>"
		}

		rpc.LogEvent(loggerlevel.Warn, fmt.Sprintf("invalid login: %s", username))
		c.Status(http.StatusUnauthorized)
		return
	}

	// TODO: get user id from jwt
	userID := "get from jwt"
	rpc.LogEvent(loggerlevel.Info, fmt.Sprintf("user login: %s", userID))
	c.JSON(http.StatusOK, auth)
}

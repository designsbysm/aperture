package auth

import (
	"fmt"
	"net/http"

	"aperture/gateway/rpc"
	"aperture/go-types/emailaddress"
	"aperture/go-types/loggerlevel"

	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Username emailaddress.T
	Password string
}

func login(c *gin.Context) {
	rpc.LogEvent(loggerlevel.Info, "login request")

	req := loginRequest{}
	if err := c.BindJSON(&req); err != nil {
		//nolint:errcheck
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	key := http.CanonicalHeaderKey("long-lived-token")
	userToken := c.GetHeader(key)
	auth, err := rpc.Login(req.Username, req.Password, userToken)
	if err != nil {
		username := req.Username
		if username == "" {
			username = "<empty>"
		}

		rpc.LogEvent(loggerlevel.Warn, fmt.Sprintf("invalid login: %s", username))
		c.Status(http.StatusUnauthorized)
		return
	}

	rpc.LogEvent(loggerlevel.Info, fmt.Sprintf("user login: %s", req.Username))
	c.JSON(http.StatusOK, auth)
}

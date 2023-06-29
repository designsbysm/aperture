package auth

import (
	"net/http"

	"aperture/gateway/rpc"
	"aperture/types/emailaddress"
	"aperture/types/loggerlevel"

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

	key := http.CanonicalHeaderKey("long-lived-token")
	userToken := c.GetHeader(key)
	auth, err := rpc.Login(req.Username, req.Password, userToken)
	if err != nil {
		rpc.LogEvent(loggerlevel.Warn, err.Error())
		c.Status(http.StatusUnauthorized)
		return
	}

	c.JSON(http.StatusOK, auth)
}

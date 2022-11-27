package users

import (
	"net/http"

	"aperture/go-types/loggerlevel"
	"aperture/service-authentication/database"
	"aperture/service-authentication/rpc"

	"github.com/gin-gonic/gin"
)

func list(c *gin.Context) {
	user := database.User{}

	list, err := user.List()
	if err != nil {
		//nolint:errcheck
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	rpc.LogEvent(loggerlevel.Info, "users listed")
	c.JSON(http.StatusOK, list)
}

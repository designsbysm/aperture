package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smaperture/go-types/loggerlevel"
	"github.com/smaperture/service-authentication/database"
	"github.com/smaperture/service-authentication/rpc"
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

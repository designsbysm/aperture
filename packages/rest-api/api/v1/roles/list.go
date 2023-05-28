package roles

import (
	"net/http"

	"aperture/go-types/loggerlevel"
	"aperture/service-authentication/database"
	"aperture/service-authentication/rpc"

	"github.com/gin-gonic/gin"
)

func list(c *gin.Context) {
	role := database.Role{}

	list, err := role.List()
	if err != nil {
		//nolint:errcheck
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	rpc.LogEvent(loggerlevel.Info, "roles listed")
	c.JSON(http.StatusOK, list)
}

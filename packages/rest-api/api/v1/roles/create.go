package roles

import (
	"aperture/go-types/loggerlevel"
	"aperture/service-auth/database"
	"aperture/service-auth/rpc"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func create(c *gin.Context) {
	role := database.Role{}
	if err := c.BindJSON(&role); err != nil {
		//nolint:errcheck
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := role.Create(); err != nil {
		//nolint:errcheck
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	rpc.LogEvent(loggerlevel.Info, fmt.Sprintf("role created: %d", role.ID))
	c.JSON(http.StatusOK, role)
}

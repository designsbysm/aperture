package users

import (
	"fmt"
	"net/http"

	"aperture/go-types/loggerlevel"
	"aperture/service-auth/database"
	"aperture/service-auth/rpc"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		//nolint:errcheck
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user := database.User{
		ID: id,
	}
	if err = user.Delete(); err != nil {
		//nolint:errcheck
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	rpc.LogEvent(loggerlevel.Info, fmt.Sprintf("user deleted: %s", user.ID))
	c.Status(http.StatusOK)
}

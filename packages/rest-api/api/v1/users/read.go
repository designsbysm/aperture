package users

import (
	"fmt"
	"net/http"

	"aperture/go-types/loggerlevel"
	"aperture/service-authentication/database"
	"aperture/service-authentication/rpc"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func read(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		//nolint:errcheck
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user := database.User{
		ID: id,
	}
	if err = user.Read(); err != nil {
		status := http.StatusInternalServerError
		if err == gorm.ErrRecordNotFound {
			status = http.StatusBadRequest
		}

		//nolint:errcheck
		c.AbortWithError(status, err)
		return
	}

	rpc.LogEvent(loggerlevel.Info, fmt.Sprintf("user read: %s", user.ID))
	c.JSON(http.StatusOK, user)
}

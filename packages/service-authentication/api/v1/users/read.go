package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/smaperture/service-authentication/database"
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

	c.JSON(http.StatusOK, user)
}

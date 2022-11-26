package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/smaperture/service-authentication/database"
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

	c.Status(http.StatusOK)
}

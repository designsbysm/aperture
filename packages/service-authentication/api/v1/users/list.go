package users

import (
	"net/http"

	"github.com/smaperture/service-authentication/database"

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

	c.JSON(http.StatusOK, list)
}

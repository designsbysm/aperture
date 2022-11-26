package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smaperture/service-authentication/database"
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

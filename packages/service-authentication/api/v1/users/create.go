package users

import (
	"net/http"

	"github.com/smaperture/service-authentication/database"

	"github.com/gin-gonic/gin"
)

func create(c *gin.Context) {
	user := database.User{}
	if err := c.BindJSON(&user); err != nil {
		//nolint:errcheck
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := user.Create(); err != nil {
		//nolint:errcheck
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	role := database.Role{
		UserID: user.ID,
		Role:   database.RoleUser,
	}
	if err := user.Role.IsValid(); err == nil {
		role.Role = user.Role
	}

	if err := role.Create(); err != nil {
		//nolint:errcheck
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

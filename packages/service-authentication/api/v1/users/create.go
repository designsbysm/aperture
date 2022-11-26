package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smaperture/service-authentication/database"
)

type CreateRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func create(c *gin.Context) {
	req := CreateRequest{}
	if err := c.BindJSON(&req); err != nil {
		//nolint:errcheck
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user := database.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	}
	if err := user.PasswordEncrypt(req.Password); err != nil {
		//nolint:errcheck
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := user.Create(); err != nil {
		//nolint:errcheck
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

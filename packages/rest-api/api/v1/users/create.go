package users

import (
	"fmt"
	"net/http"

	"aperture/go-types/emailaddress"
	"aperture/go-types/loggerlevel"
	"aperture/go-types/phonenumber"
	"aperture/service-auth/database"
	"aperture/service-auth/rpc"

	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	FirstName string
	LastName  string
	Mobile    phonenumber.T
	Email     emailaddress.T
	Password  string
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
		Mobile:    req.Mobile,
		Email:     req.Email,
	}
	if err := user.PasswordEncrypt(req.Password); err != nil {
		//nolint:errcheck
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := user.Create(); err != nil {
		if err := user.IsDuplicateError(err); err != nil {
			//nolint:errcheck
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		//nolint:errcheck
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	rpc.LogEvent(loggerlevel.Info, fmt.Sprintf("user created: %s", user.ID))
	c.JSON(http.StatusOK, user)
}

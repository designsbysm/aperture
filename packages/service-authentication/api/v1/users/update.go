package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/smaperture/go-types/emailaddress"
	"github.com/smaperture/go-types/loggerlevel"
	"github.com/smaperture/service-authentication/database"
	"github.com/smaperture/service-authentication/rpc"
	"gorm.io/gorm"
)

type UpdateReqest struct {
	FirstName string         `json:"firstName"`
	LastName  string         `json:"lastName"`
	Email     emailaddress.T `json:"email"`
	Password  string         `json:"password"`
}

func update(c *gin.Context) {
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

	req := UpdateReqest{}
	if err = c.BindJSON(&req); err != nil {
		//nolint:errcheck
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}
	if req.LastName != "" {
		user.LastName = req.LastName
	}
	if req.Email != "" {
		user.Email = req.Email
	}

	//nolint:errcheck
	user.PasswordEncrypt(req.Password)

	if err = user.Update(); err != nil {
		if err := user.IsDuplicateError(err); err != nil {
			//nolint:errcheck
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		//nolint:errcheck
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	rpc.LogEvent(loggerlevel.Info, fmt.Sprintf("user updated: %s", user.ID))
	c.JSON(http.StatusOK, user)
}

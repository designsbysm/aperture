package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/smaperture/service-authentication/database"
	"gorm.io/gorm"
)

type UpdateReqest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
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
		//nolint:errcheck
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

package users

// import (
// 	"fmt"
// 	"net/http"

// 	"aperture/go-types/emailaddress"
// 	"aperture/go-types/loggerlevel"
// 	"aperture/go-types/phonenumber"
// 	"aperture/service-auth/database"
// 	"aperture/service-auth/rpc"

// 	"github.com/gin-gonic/gin"
// 	"github.com/gofrs/uuid"
// 	"gorm.io/gorm"
// )

// type UpdateRequest struct {
// 	FirstName string
// 	LastName  string
// 	Mobile    phonenumber.T
// 	Email     emailaddress.T
// 	Password  string
// }

// func update(c *gin.Context) {
// 	id, err := uuid.Parse(c.Param("id"))
// 	if err != nil {
// 		//nolint:errcheck
// 		c.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}

// 	user := database.User{
// 		ID: id,
// 	}
// 	if err = user.Read(); err != nil {
// 		status := http.StatusInternalServerError
// 		if err == gorm.ErrRecordNotFound {
// 			status = http.StatusBadRequest
// 		}

// 		//nolint:errcheck
// 		c.AbortWithError(status, err)
// 		return
// 	}

// 	req := UpdateRequest{}
// 	if err = c.BindJSON(&req); err != nil {
// 		//nolint:errcheck
// 		c.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}
// 	if req.FirstName != "" {
// 		user.FirstName = req.FirstName
// 	}
// 	if req.LastName != "" {
// 		user.LastName = req.LastName
// 	}
// 	if req.Mobile != "" {
// 		user.Mobile = req.Mobile
// 	}
// 	if req.Email != "" {
// 		user.Email = req.Email
// 	}

// 	if user.PasswordEncrypt(req.Password); err != nil {
// 		//nolint:errcheck
// 		c.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}

// 	if err = user.Update(); err != nil {
// 		if err := user.IsDuplicateError(err); err != nil {
// 			//nolint:errcheck
// 			c.AbortWithError(http.StatusBadRequest, err)
// 			return
// 		}

// 		//nolint:errcheck
// 		c.AbortWithError(http.StatusInternalServerError, err)
// 		return
// 	}

// 	rpc.LogEvent(loggerlevel.Info, fmt.Sprintf("user updated: %s", user.ID))
// 	c.JSON(http.StatusOK, user)
// }

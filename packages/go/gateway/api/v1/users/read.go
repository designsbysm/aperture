package users

// import (
// 	"fmt"
// 	"net/http"

// 	"aperture/types/loggerlevel"
// 	"aperture/service-auth/database"
// 	"aperture/service-auth/rpc"

// 	"github.com/gin-gonic/gin"
// 	"github.com/gofrs/uuid"
// 	"gorm.io/gorm"
// )

// func read(c *gin.Context) {
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

// 	rpc.LogEvent(loggerlevel.Info, fmt.Sprintf("user read: %s", user.ID))
// 	c.JSON(http.StatusOK, user)
// }

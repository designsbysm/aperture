package roles

// import (
// 	"fmt"
// 	"net/http"
// 	"strconv"

// 	"aperture/types/loggerlevel"
// 	"aperture/service-auth/database"
// 	"aperture/service-auth/rpc"

// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm"
// )

// func read(c *gin.Context) {
// 	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
// 	if err != nil {
// 		//nolint:errcheck
// 		c.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}

// 	role := database.Role{
// 		ID: id,
// 	}
// 	if err = role.Read(); err != nil {
// 		status := http.StatusInternalServerError
// 		if err == gorm.ErrRecordNotFound {
// 			status = http.StatusBadRequest
// 		}

// 		//nolint:errcheck
// 		c.AbortWithError(status, err)
// 		return
// 	}

// 	rpc.LogEvent(loggerlevel.Info, fmt.Sprintf("role read: %d", role.ID))
// 	c.JSON(http.StatusOK, role)
// }

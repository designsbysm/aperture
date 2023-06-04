package roles

// import (
// 	"fmt"
// 	"net/http"
// 	"strconv"

// 	"aperture/go-types/loggerlevel"
// 	"aperture/service-auth/database"
// 	"aperture/service-auth/rpc"

// 	"github.com/gin-gonic/gin"
// )

// func delete(c *gin.Context) {
// 	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
// 	if err != nil {
// 		//nolint:errcheck
// 		c.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}

// 	role := database.Role{
// 		ID: id,
// 	}
// 	if err = role.Delete(); err != nil {
// 		//nolint:errcheck
// 		c.AbortWithError(http.StatusInternalServerError, err)
// 		return
// 	}

// 	rpc.LogEvent(loggerlevel.Info, fmt.Sprintf("role deleted: %d", role.ID))
// 	c.Status(http.StatusOK)
// }

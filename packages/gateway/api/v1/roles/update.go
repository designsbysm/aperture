package roles

// import (
// 	"fmt"
// 	"net/http"
// 	"strconv"

// 	"aperture/go-types/loggerlevel"
// 	"aperture/go-types/userrole"
// 	"aperture/service-auth/database"
// 	"aperture/service-auth/rpc"

// 	"github.com/gin-gonic/gin"
// 	"github.com/google/uuid"
// 	"gorm.io/gorm"
// )

// type UpdateRequest struct {
// 	UserID string
// 	Role   userrole.T
// }

// func update(c *gin.Context) {
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

// 	req := UpdateRequest{}
// 	if err = c.BindJSON(&req); err != nil {
// 		//nolint:errcheck
// 		c.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}

// 	if req.UserID != "" {
// 		userID, err := uuid.Parse(req.UserID)
// 		if err != nil {
// 			//nolint:errcheck
// 			c.AbortWithError(http.StatusInternalServerError, err)
// 			return
// 		}

// 		role.UserID = userID
// 	}
// 	if req.Role.IsValid() {
// 		role.Role = req.Role
// 	}

// 	if err = role.Update(); err != nil {
// 		//nolint:errcheck
// 		c.AbortWithError(http.StatusInternalServerError, err)
// 		return
// 	}

// 	rpc.LogEvent(loggerlevel.Info, fmt.Sprintf("role updated: %d", role.ID))
// 	c.JSON(http.StatusOK, role)
// }

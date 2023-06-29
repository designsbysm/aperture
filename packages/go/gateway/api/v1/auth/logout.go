package auth

// import (
// 	"errors"
// 	"fmt"
// 	"net/http"

// 	"aperture/types/loggerlevel"
// 	"aperture/service-auth/database"
// 	"aperture/service-auth/rpc"

// 	"github.com/gin-gonic/gin"
// 	"github.com/gofrs/uuid"
// )

// func logout(c *gin.Context) {
// 	data, ok := c.Get("userID")
// 	if !ok {
// 		//nolint:errcheck
// 		c.AbortWithError(http.StatusInternalServerError, errors.New("missing user id"))
// 		return
// 	}
// 	userID := data.(uuid.UUID)

// 	refreshToken := database.RefreshToken{
// 		UserID: userID,
// 	}
// 	if err := refreshToken.Delete(); err != nil {
// 		//nolint:errcheck
// 		c.AbortWithError(http.StatusInternalServerError, err)
// 		return
// 	}

// 	rpc.LogEvent(loggerlevel.Info, fmt.Sprintf("user logged out: %s", userID))
// 	c.Status(http.StatusOK)
// }

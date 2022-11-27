package authentication

import (
	"errors"
	"fmt"
	"net/http"

	"aperture/go-types/loggerlevel"
	"aperture/service-authentication/database"
	"aperture/service-authentication/rpc"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func logout(c *gin.Context) {
	data, ok := c.Get("userID")
	if !ok {
		//nolint:errcheck
		c.AbortWithError(http.StatusInternalServerError, errors.New("missing user id"))
		return
	}
	userID := data.(uuid.UUID)

	refreshToken := database.RefreshToken{
		UserID: userID,
	}
	if err := refreshToken.Delete(); err != nil {
		//nolint:errcheck
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	rpc.LogEvent(loggerlevel.Info, fmt.Sprintf("user logged out: %s", userID))
	c.Status(http.StatusOK)
}

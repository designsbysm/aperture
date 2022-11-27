package authentication

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/smaperture/service-authentication/database"
	"github.com/smaperture/service-authentication/rpc"
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

	rpc.LogEvent("info", fmt.Sprintf("user logged out: %s", userID))
	c.Status(http.StatusOK)
}

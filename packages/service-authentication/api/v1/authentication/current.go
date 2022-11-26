package authentication

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/smaperture/service-authentication/database"
	"github.com/smaperture/service-authentication/types/role"
)

type CurrentResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Role      role.T    `json:"role"`
}

func current(c *gin.Context) {
	data, ok := c.Get("userID")
	if !ok {
		//nolint:errcheck
		c.AbortWithError(http.StatusInternalServerError, errors.New("missing user id"))
		return
	}
	userID := data.(uuid.UUID)

	user := database.User{
		ID: userID,
	}
	if err := user.Read(); err != nil {
		//nolint:errcheck
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res := CurrentResponse{}
	temp, err := json.Marshal(user)
	if err != nil {
		//nolint:errcheck
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if err = json.Unmarshal(temp, &res); err != nil {
		//nolint:errcheck
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	data, ok = c.Get("userRole")
	if !ok {
		//nolint:errcheck
		c.AbortWithError(http.StatusInternalServerError, errors.New("missing user role"))
		return
	}
	userRole := data.(string)

	res.Role = role.FromString(userRole)

	c.JSON(http.StatusOK, res)
}

package authentication

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smaperture/go-types/jwt"
	"github.com/smaperture/go-types/loggerlevel"
	"github.com/smaperture/go-types/userrole"
	"github.com/smaperture/service-authentication/database"
	"github.com/smaperture/service-authentication/rpc"
)

type SignupRequest struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func signup(c *gin.Context) {
	req := SignupRequest{}
	if err := c.BindJSON(&req); err != nil {
		//nolint:errcheck
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user := database.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	}
	if err := user.PasswordEncrypt(req.Password); err != nil {
		//nolint:errcheck
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if err := user.Create(); err != nil {
		//nolint:errcheck
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	role := database.Role{
		UserID: user.ID,
		Role:   userrole.User,
	}
	if err := role.Create(); err != nil {
		//nolint:errcheck
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	accessToken, err := jwt.Encode(user.ID, role.Role, false)
	if err != nil {
		//nolint:errcheck
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	refreshToken := database.RefreshToken{
		UserID: user.ID,
	}
	if err := refreshToken.Create(); err != nil {
		//nolint:errcheck
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	res := AuthenticationResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken.ID,
	}

	rpc.LogEvent(loggerlevel.Info, fmt.Sprintf("user signed up: %s", user.ID))
	c.JSON(http.StatusOK, res)
}

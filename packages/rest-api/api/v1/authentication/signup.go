package authentication

// import (
// 	"fmt"
// 	"net/http"

// 	"aperture/go-types/emailaddress"
// 	"aperture/go-types/jwt"
// 	"aperture/go-types/loggerlevel"
// 	"aperture/go-types/phonenumber"
// 	"aperture/go-types/userrole"
// 	"aperture/service-auth/database"
// 	"aperture/service-auth/rpc"

// 	"github.com/gin-gonic/gin"
// )

// type SignupRequest struct {
// 	FirstName string
// 	LastName  string
// 	Mobile    phonenumber.T
// 	Email     emailaddress.T
// 	Password  string
// }

// func signup(c *gin.Context) {
// 	req := SignupRequest{}
// 	if err := c.BindJSON(&req); err != nil {
// 		//nolint:errcheck
// 		c.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}

// 	user := database.User{
// 		FirstName: req.FirstName,
// 		LastName:  req.LastName,
// 		Mobile:    req.Mobile,
// 		Email:     req.Email,
// 	}
// 	if err := user.PasswordEncrypt(req.Password); err != nil {
// 		//nolint:errcheck
// 		c.AbortWithError(http.StatusInternalServerError, err)
// 		return
// 	}

// 	if err := user.Create(); err != nil {
// 		if err := user.IsDuplicateError(err); err != nil {
// 			//nolint:errcheck
// 			c.AbortWithError(http.StatusBadRequest, err)
// 			return
// 		}

// 		//nolint:errcheck
// 		c.AbortWithError(http.StatusInternalServerError, err)
// 		return
// 	}

// 	role := database.Role{
// 		UserID: user.ID,
// 		Role:   userrole.User,
// 	}
// 	if err := role.Create(); err != nil {
// 		//nolint:errcheck
// 		c.AbortWithError(http.StatusInternalServerError, err)
// 		return
// 	}

// 	accessToken, err := jwt.Encode(user.ID, role.Role, false)
// 	if err != nil {
// 		//nolint:errcheck
// 		c.AbortWithError(http.StatusInternalServerError, err)
// 		return
// 	}

// 	refreshToken := database.RefreshToken{
// 		UserID: user.ID,
// 	}
// 	if err := refreshToken.Create(); err != nil {
// 		//nolint:errcheck
// 		c.AbortWithError(http.StatusInternalServerError, err)
// 		return
// 	}

// 	res := AuthenticationResponse{
// 		AccessToken:  accessToken,
// 		RefreshToken: refreshToken.ID,
// 	}

// 	rpc.LogEvent(loggerlevel.Info, fmt.Sprintf("user signed up: %s", user.ID))
// 	c.JSON(http.StatusOK, res)
// }

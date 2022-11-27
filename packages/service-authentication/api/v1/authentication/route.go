package authentication

import (
	"aperture/service-authentication/middleware"

	"github.com/gin-gonic/gin"
)

func AddRoute(r *gin.RouterGroup) {
	group := r.Group("/authentication")
	{
		group.POST("/current", middleware.AuthorizeJWT(), current)
		group.POST("/login", login)
		group.POST("/logout", middleware.AuthorizeJWT(), logout)
		group.POST("/refreshAccess/:token", refreshAccess)
		group.POST("/signup", signup)
	}
}

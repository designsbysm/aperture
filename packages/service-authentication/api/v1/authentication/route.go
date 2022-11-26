package authentication

import (
	"github.com/gin-gonic/gin"
	"github.com/smaperture/service-authentication/middleware"
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

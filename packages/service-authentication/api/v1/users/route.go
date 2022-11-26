package users

import (
	"github.com/gin-gonic/gin"
	"github.com/smaperture/service-authentication/middleware"
)

func AddRoute(r *gin.RouterGroup) {
	group := r.Group("/users")
	{
		group.GET("/", middleware.AuthorizeJWT(), list)
		group.POST("/", middleware.AuthorizeJWT(), create)

		id := group.Group("/:id")
		{
			id.GET("/", middleware.AuthorizeJWT(), read)
			id.PUT("/", middleware.AuthorizeJWT(), update)
			id.DELETE("/", middleware.AuthorizeJWT(), delete)
		}
	}
}

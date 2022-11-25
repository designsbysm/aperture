package users

import (
	"github.com/gin-gonic/gin"
)

func AddRoute(r *gin.RouterGroup) {
	group := r.Group("/users")
	{
		group.GET("/", list)
		group.POST("/", create)

		id := group.Group("/:id")
		{
			id.GET("/", read)
			id.PUT("/", update)
			id.DELETE("/", delete)
		}
	}
}

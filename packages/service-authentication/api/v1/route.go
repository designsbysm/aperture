package v1

import (
	"github.com/smaperture/service-authentication/api/v1/users"

	"github.com/gin-gonic/gin"
)

func AddRoute(r *gin.RouterGroup) {
	group := r.Group("/v1")
	{
		users.AddRoute(group)
	}
}

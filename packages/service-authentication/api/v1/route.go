package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/smaperture/service-authentication/api/v1/authentication"
	"github.com/smaperture/service-authentication/api/v1/users"
)

func AddRoute(r *gin.RouterGroup) {
	group := r.Group("/v1")
	{
		authentication.AddRoute(group)
		users.AddRoute(group)
	}
}

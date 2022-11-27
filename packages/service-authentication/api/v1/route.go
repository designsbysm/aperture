package v1

import (
	"aperture/service-authentication/api/v1/authentication"
	"aperture/service-authentication/api/v1/users"

	"github.com/gin-gonic/gin"
)

func AddRoute(r *gin.RouterGroup) {
	group := r.Group("/v1")
	{
		authentication.AddRoute(group)
		users.AddRoute(group)
	}
}

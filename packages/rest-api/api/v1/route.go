package v1

import (
	"aperture/rest-api/api/v1/appointments"
	"aperture/rest-api/api/v1/auth"
	"aperture/rest-api/api/v1/healthcheck"

	"github.com/gin-gonic/gin"
)

func AddRoute(r *gin.RouterGroup) {
	group := r.Group("/v1")
	{
		appointments.AddRoute(group)
		auth.AddRoute(group)
		healthcheck.AddRoute(group)
	}
}

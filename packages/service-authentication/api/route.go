package api

import (
	v1 "github.com/smaperture/service-authentication/api/v1"

	"github.com/designsbysm/gibson"
	"github.com/gin-gonic/gin"
)

func AddRoute(r *gin.Engine) {
	r.Use(
		gin.Recovery(),
		gibson.Logger(),
		gibson.Error(),
	)

	group := r.Group("api")
	{
		v1.AddRoute(group)
	}
}

package api

import (
	v1 "aperture/rest-api/api/v1"

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

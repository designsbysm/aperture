package api

import (
	v1 "aperture/gateway/api/v1"
	"aperture/gateway/middleware"

	"github.com/designsbysm/gibson"
	"github.com/gin-gonic/gin"
)

func AddRoute(r *gin.Engine) {
	r.Use(
		gin.Recovery(),
		gibson.Logger(),
		middleware.RPCLogger(),
		gibson.Error(),
	)

	group := r.Group("api")
	{
		v1.AddRoute(group)
	}
}

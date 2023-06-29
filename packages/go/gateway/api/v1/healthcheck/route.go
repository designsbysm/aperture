package healthcheck

import (
	"github.com/gin-gonic/gin"
)

func AddRoute(r *gin.RouterGroup) {
	r.GET("/healthcheck", healthcheck)
}

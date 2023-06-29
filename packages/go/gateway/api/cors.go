package api

import (
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func addCORS(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return strings.HasPrefix(origin, "http://localhost:")
		},
		AllowHeaders:  []string{"Authorization", "Content-Length", "Content-Type", "Host", "Referrer", "Origin", "User-Agent"},
		AllowMethods:  []string{"DELETE", "GET", "POST", "OPTIONS", "PUT"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))
}

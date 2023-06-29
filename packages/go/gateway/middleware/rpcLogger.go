package middleware

import (
	"aperture/gateway/rpc"
	"aperture/types/loggerlevel"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func RPCLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		now := time.Now()

		c.Next()

		if gin.IsDebugging() {
			return
		}

		latency := time.Since(now)
		rpc.LogEvent(loggerlevel.Info, fmt.Sprintf("%d %s %s %v", c.Writer.Status(), c.Request.Method, c.Request.URL, latency.Round(time.Millisecond)))
	}
}

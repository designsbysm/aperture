package healthcheck

import (
	"aperture/gateway/rpc"
	"aperture/go-types/loggerlevel"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func healthcheck(c *gin.Context) {
	id := uuid.New().String()

	rpc.LogEvent(loggerlevel.Info, fmt.Sprintf("healthcheck: %s", id))
	c.String(http.StatusOK, id)
}

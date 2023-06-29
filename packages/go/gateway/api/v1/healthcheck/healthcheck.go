package healthcheck

import (
	"aperture/gateway/rpc"
	"aperture/types/loggerlevel"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func healthcheck(c *gin.Context) {
	id, _ := uuid.NewV7()

	rpc.LogEvent(loggerlevel.Info, fmt.Sprintf("healthcheck: %s", id))
	c.String(http.StatusOK, id.String())
}

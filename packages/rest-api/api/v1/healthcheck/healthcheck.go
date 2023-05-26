package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func healthcheck(c *gin.Context) {
	c.String(http.StatusOK, uuid.New().String())
}

package appointments

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func readOne(c *gin.Context) {
	// FIXME: remove dev code
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		//nolint:errcheck
		c.Error(err)
	}

	c.JSON(http.StatusOK, id)
}

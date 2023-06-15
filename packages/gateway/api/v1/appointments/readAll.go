package appointments

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func readAll(c *gin.Context) {
	time.Sleep(1 * time.Second)

	appointments := []Appointment{
		{
			ID:     uuid.New(),
			Date:   time.Now(),
			Status: "Pending",
		},
		{
			ID:     uuid.New(),
			Date:   time.Now(),
			Status: "Negative",
		},
		{
			ID:     uuid.New(),
			Date:   time.Now(),
			Status: "Positive",
		},
	}

	c.JSON(http.StatusOK, appointments)
}

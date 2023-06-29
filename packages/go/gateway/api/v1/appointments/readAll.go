package appointments

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func readAll(c *gin.Context) {
	time.Sleep(1 * time.Second)

	id1, _ := uuid.NewV7()
	id2, _ := uuid.NewV7()
	id3, _ := uuid.NewV7()

	appointments := []Appointment{
		{
			ID:     id1,
			Date:   time.Now(),
			Status: "Pending",
		},
		{
			ID:     id2,
			Date:   time.Now(),
			Status: "Negative",
		},
		{
			ID:     id3,
			Date:   time.Now(),
			Status: "Positive",
		},
	}

	c.JSON(http.StatusOK, appointments)
}

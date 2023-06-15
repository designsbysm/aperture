package appointments

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Appointment struct {
	ID     uuid.UUID `json:"id"`
	Date   time.Time `json:"date"`
	Status string    `json:"status"`
}

func AddRoute(r *gin.RouterGroup) {
	group := r.Group("/appointments")
	{
		group.GET("", readAll)

		idGroup := group.Group("/:id")
		{
			idGroup.GET("", readOne)
		}
	}
}

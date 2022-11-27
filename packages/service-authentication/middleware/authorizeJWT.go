package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/smaperture/go-types/jwt"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		token := strings.TrimPrefix(header, "Bearer ")

		claims, err := jwt.Decode(token)
		if err != nil {
			//nolint:errcheck
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		id, err := uuid.Parse(claims["userID"].(string))
		if err != nil {
			//nolint:errcheck
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Set("userID", id)
		c.Set("userRole", claims["userRole"].(string))
		c.Next()
	}
}

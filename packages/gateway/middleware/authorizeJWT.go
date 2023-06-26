package middleware

import (
	"errors"
	"net/http"
	"strings"

	"aperture/go-types/jwt"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		token := strings.TrimPrefix(header, "Bearer ")

		if token == "" {
			//nolint:errcheck
			c.AbortWithError(http.StatusUnauthorized, errors.New("authorization bearer token missing"))
			return
		}

		claims, err := jwt.Decode(token)
		if err != nil {
			//nolint:errcheck
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		id, err := uuid.FromString(claims["userID"].(string))
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

package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/smaperture/service-authentication/database"
	"github.com/spf13/viper"
)

func Encode(id uuid.UUID, role database.RoleEnum, longLived bool) (string, error) {
	defaulDuration := viper.GetInt("jwt.expiration")
	if longLived {
		defaulDuration = viper.GetInt("jwt.longLived.expiration")
	}
	expiration := time.Now().Add(time.Minute * time.Duration(defaulDuration)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":   id,
		"userRole": role,
		"exp":      expiration,
	})

	secretKey := []byte(viper.GetString("jwt.secret"))

	return token.SignedString(secretKey)
}

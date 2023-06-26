package jwt

import (
	"time"

	"aperture/go-types/userrole"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

func Encode(id uuid.UUID, role userrole.T, firstName string, lastName string, longLived bool) (string, error) {
	defaulDuration := viper.GetInt("jwt.expiration")
	if longLived {
		defaulDuration = viper.GetInt("jwt.longLived.expiration")
	}
	expiration := time.Now().Add(time.Minute * time.Duration(defaulDuration)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    id,
		"userRole":  role,
		"firstName": firstName,
		"lastName":  lastName,
		"exp":       expiration,
	})

	secretKey := []byte(viper.GetString("jwt.secret"))

	return token.SignedString(secretKey)
}

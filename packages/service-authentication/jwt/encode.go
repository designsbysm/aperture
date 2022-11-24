package jwt

// func Encode(id uuid.UUID, role database.Role) (string, error) {
// 	defaulDuration := 8
// 	if role.IsAdmin {
// 		defaulDuration = 24 * 365
// 	}
// 	expiration := time.Now().Add(time.Hour * time.Duration(defaulDuration)).Unix()

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"id":   id,
// 		"role": role.Name,
// 		"exp":  expiration,
// 	})

// 	secretKey := []byte(viper.GetString("jwt.secret"))

// 	return token.SignedString(secretKey)
// }

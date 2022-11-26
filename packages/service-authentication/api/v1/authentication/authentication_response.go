package authentication

import (
	"github.com/google/uuid"
)

type AuthenticationResponse struct {
	AccessToken  string    `json:"accessToken"`
	RefreshToken uuid.UUID `json:"refreshToken"`
}

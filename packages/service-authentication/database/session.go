package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Session struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *gorm.DeletedAt

	UserID uuid.UUID
	Token  string
}

func (s *Session) Read() error {
	return DB.First(&s, s).Error
}

func (s *Session) Upsert() error {
	tempToken := s.Token

	if err := DB.FirstOrCreate(&s, Session{
		UserID: s.UserID,
	}).Error; err != nil {
		return err
	}
	s.Token = tempToken

	return DB.Save(&s).Error
}

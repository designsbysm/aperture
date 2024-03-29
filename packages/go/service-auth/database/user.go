package database

import (
	"regexp"
	"time"

	"aperture/types/emailaddress"
	"aperture/types/phonenumber"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty"`
	FirstName string         `json:"firstName"`
	LastName  string         `json:"lastName"`
	Mobile    phonenumber.T  `json:"mobile"`
	Email     emailaddress.T `json:"email"`
	Password  string         `json:"-"`
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	if u.FirstName == "" {
		return ErrFistNameRequired
	}

	if u.LastName == "" {
		return ErrLastNameRequired
	}

	if !u.Mobile.IsValid() {
		return ErrInvalidMobile
	}

	if !u.Email.IsValid() {
		return ErrInvalidEmail
	}

	if u.Password == "" {
		return ErrPasswordRequired
	}

	return nil
}

func (u *User) Create() error {
	return DB.Create(&u).Error
}

func (u *User) Delete() error {
	return DB.Delete(&u, u).Error
}

func (u User) IsDuplicateError(err error) error {
	regex := regexp.MustCompile(`duplicate key.*users_email`)
	if regex.MatchString(err.Error()) {
		return ErrDuplicateEmail
	}

	regex = regexp.MustCompile(`duplicate key.*users_mobile`)
	if regex.MatchString(err.Error()) {
		return ErrDuplicateMobile
	}

	return nil
}

func (u *User) List() ([]User, error) {
	list := []User{}
	err := DB.Find(&list, u).Error

	return list, err
}

func (u *User) PasswordEncrypt(password string) error {
	if password == "" {
		return ErrPasswordRequired
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hash)

	return nil
}

func (u *User) PasswordValidate(password string) error {
	if password == "" {
		return ErrPasswordRequired
	}

	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func (u *User) Read() error {
	return DB.First(&u, u).Error
}

func (u *User) Update() error {
	return DB.Save(&u).Error
}

package database

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID      `json:"id"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt,omitempty"`
	FirstName   string         `json:"firstName"`
	LastName    string         `json:"lastName"`
	Email       string         `json:"email"`
	Password    string         `json:"-"`
	RawPassword string         `gorm:"-" json:"password,omitempty"`
	Role        RoleEnum       `gorm:"-" json:"role"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()

	return nil
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	if u.Password == "" && u.RawPassword == "" {
		return ErrPasswordRequired
	} else if u.RawPassword == "" {
		return nil
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(u.RawPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hash)
	u.RawPassword = ""

	return nil
}

// TODO: generic crud?

func (u *User) Create() error {
	return DB.Create(&u).Error
}

func (u *User) Delete() error {
	return DB.Delete(&u, u).Error
}

func (u *User) List() ([]User, error) {
	list := []User{}
	err := DB.Find(&list, u).Error

	return list, err
}

func (u *User) Read() error {
	return DB.First(&u, u).Error
}

func (u *User) Update() error {
	return DB.Save(&u).Error
}

func (u *User) ValidatePassword(password string) error {
	if password == "" {
		return ErrPasswordRequired
	}

	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

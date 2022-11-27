package database

import (
	"errors"
)

var ErrDuplicateEmail = errors.New("email address already in use")
var ErrDuplicateMobile = errors.New("mobile phone already in use")
var ErrFistNameRequired = errors.New("first name is required")
var ErrInvalidEmail = errors.New("email is invalid")
var ErrInvalidMobile = errors.New("mobile phone is invalid")
var ErrLastNameRequired = errors.New("last name is required")
var ErrPasswordRequired = errors.New("password is required")

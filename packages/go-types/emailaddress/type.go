package emailaddress

import (
	"errors"
	"regexp"
)

type T string

// TODO: add FromString & ToString functions with validity checks

func (t T) IsValid() bool {
	regex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return regex.MatchString(string(t))
}

func FromString(email string) (T, error) {
	temp := T(email)

	if !temp.IsValid() {
		return T(""), errors.New("email address is invalid")
	}

	return temp, nil
}

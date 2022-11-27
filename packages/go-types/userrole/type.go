package userrole

import "errors"

type T string

const (
	Admin      T = "admin"
	Diabled    T = "disabled"
	Operations T = "operations"
	Provider   T = "provider"
	User       T = "user"
)

func (t T) IsValid() error {
	switch t {
	case Admin, Diabled, Operations, Provider, User:
		return nil
	}

	return errors.New("invalid userrole")
}

func FromString(str string) T {
	roles := map[string]T{
		"admin":      Admin,
		"disabled":   Diabled,
		"operations": Operations,
		"provider":   Provider,
		"user":       User,
	}

	return roles[str]
}

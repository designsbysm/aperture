package database

import "errors"

type RoleEnum string

const (
	RoleAdmin      RoleEnum = "admin"
	RoleOperations RoleEnum = "operations"
	RoleProvider   RoleEnum = "provider"
	RoleUser       RoleEnum = "user"
)

func (e RoleEnum) Level() int {
	roles := map[RoleEnum]int{
		RoleAdmin:      0,
		RoleOperations: 1,
		RoleProvider:   2,
		RoleUser:       3,
	}

	return roles[e]
}

func (e RoleEnum) IsValid() error {
	switch e {
	case RoleAdmin, RoleOperations, RoleProvider, RoleUser:
		return nil
	}

	return errors.New("invalid role type")
}

func RoleFromString(str string) RoleEnum {
	roles := map[string]RoleEnum{
		"admin":      RoleAdmin,
		"operations": RoleOperations,
		"provider":   RoleProvider,
		"user":       RoleUser,
	}

	return roles[str]
}

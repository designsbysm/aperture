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
	levels := map[RoleEnum]int{
		RoleAdmin:      0,
		RoleOperations: 1,
		RoleProvider:   2,
		RoleUser:       3,
	}

	return levels[e]
}

func (r RoleEnum) IsValid() error {
	switch r {
	case RoleAdmin, RoleOperations, RoleProvider, RoleUser:
		return nil
	}

	return errors.New("invalid role type")
}

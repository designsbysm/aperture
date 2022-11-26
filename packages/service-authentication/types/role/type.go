package role

import "errors"

type T string

const (
	RoleAdmin      T = "admin"
	RoleDiabled    T = "disabled"
	RoleOperations T = "operations"
	RoleProvider   T = "provider"
	RoleUser       T = "user"
)

func (t T) Level() int {
	roles := map[T]int{
		RoleAdmin:      0,
		RoleOperations: 1,
		RoleProvider:   2,
		RoleUser:       3,
		RoleDiabled:    4,
	}

	return roles[t]
}

func (t T) IsValid() error {
	switch t {
	case RoleAdmin, RoleDiabled, RoleOperations, RoleProvider, RoleUser:
		return nil
	}

	return errors.New("invalid role type")
}

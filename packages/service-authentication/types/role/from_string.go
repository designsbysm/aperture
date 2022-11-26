package role

func FromString(str string) T {
	roles := map[string]T{
		"admin":      RoleAdmin,
		"disabled":   RoleDiabled,
		"operations": RoleOperations,
		"provider":   RoleProvider,
		"user":       RoleUser,
	}

	return roles[str]
}

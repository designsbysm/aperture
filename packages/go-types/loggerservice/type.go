package loggerservice

type T string

const (
	Authentication T = "authentication"
)

func (t T) IsValid() bool {
	switch t {
	case Authentication:
		return true
	}

	return false
}

func FromString(str string) T {
	services := map[string]T{
		"authentication": Authentication,
	}

	return services[str]
}

package loggerservice

type T string

const (
	Authentication T = "authentication"
	RestAPI        T = "rest-api"
)

func (t T) IsValid() bool {
	switch t {
	case Authentication:
	case RestAPI:
		return true
	}

	return false
}

// TODO: what does invalid string return
func FromString(str string) T {
	services := map[string]T{
		"authentication": Authentication,
		"rest-api":       RestAPI,
	}

	return services[str]
}

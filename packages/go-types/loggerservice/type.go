package loggerservice

import "errors"

type T string

const (
	Authentication T = "authentication"
)

func (t T) IsValid() error {
	switch t {
	case Authentication:
		return nil
	}

	return errors.New("invalid loggerservice")
}

func FromString(str string) T {
	services := map[string]T{
		"authentication": Authentication,
	}

	return services[str]
}

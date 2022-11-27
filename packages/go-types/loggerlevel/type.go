package loggerlevel

import "errors"

type T string

const (
	Error T = "error"
	Info  T = "info"
	Warn  T = "warn"
)

func (t T) IsValid() error {
	switch t {
	case Error, Info, Warn:
		return nil
	}

	return errors.New("invalid loggelevel")
}

func FromString(str string) T {
	levels := map[string]T{
		"error": Error,
		"info":  Info,
		"warn":  Warn,
	}

	return levels[str]
}

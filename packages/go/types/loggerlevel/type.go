package loggerlevel

type T string

const (
	Error T = "error"
	Info  T = "info"
	Warn  T = "warn"
)

func (t T) IsValid() bool {
	switch t {
	case Error, Info, Warn:
		return true
	}

	return false
}

// TODO: what does invalid string return?
func FromString(str string) T {
	levels := map[string]T{
		"error": Error,
		"info":  Info,
		"warn":  Warn,
	}

	return levels[str]
}

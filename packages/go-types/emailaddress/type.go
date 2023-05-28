package emailaddress

import "regexp"

type T string

// TODO: add FromString & ToString functions with validity checks

func (t T) IsValid() bool {
	regex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return regex.MatchString(string(t))
}

package phonenumber

import "regexp"

type T string

// TODO: add FromString & ToString functions with validity checks

func (t T) IsValid() bool {
	regex := regexp.MustCompile(`^\+1[2-9]\d{2}[2-9]\d{2}\d{4}$`)
	return regex.MatchString(string(t))
}

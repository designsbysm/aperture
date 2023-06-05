package emailaddress

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		id       int
		email    string
		expected bool
	}{
		{1, "", false},
		{2, "invalid", false},
		{3, "sm@sm.dev", true},
	}

	for _, test := range tests {
		email := T(test.email)
		valid := email.IsValid()

		if test.expected != valid {
			t.Errorf("%d: '%s' should %t", test.id, test.email, test.expected)
		}
	}
}

func TestFromString(t *testing.T) {
	tests := []struct {
		id       int
		email    string
		expected bool
	}{
		{1, "", false},
		{2, "invalid", false},
		{3, "sm@sm.dev", true},
	}

	for _, test := range tests {
		_, err := FromString(test.email)

		if !test.expected && err == nil {
			t.Errorf("%d: should have error, got nil", test.id)
		} else if test.expected && err != nil {
			t.Errorf("%d: should be nil, got %v", test.id, err)
		}
	}
}

package databaseconnect

import (
	"testing"
)

func TestNoConnectionString(t *testing.T) {
	_, err := Go("")

	if err == nil {
		t.Errorf("should have error, got nil")
	}
}

func TestInvalidConnectionString(t *testing.T) {
	_, err := Go("invalid")

	if err == nil {
		t.Errorf("should have error, got nil")
	}
}

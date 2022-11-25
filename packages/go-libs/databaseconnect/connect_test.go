package databaseconnect

import (
	"testing"
)

func TestNoConnectionString(t *testing.T) {
	_, err := Go("", 0)

	if err == nil {
		t.Errorf("should have error, got nil")
	}
}

func TestInvalidConnectionString(t *testing.T) {
	_, err := Go("invalid", 0)

	if err == nil {
		t.Errorf("should have error, got nil")
	}
}

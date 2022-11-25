package databaseconnect

import (
	"reflect"
	"testing"
)

func TestReturnLogger(t *testing.T) {
	logger := logger(1)

	if logger == nil {
		t.Fatalf("should have logger.Interface, got nil")
	}

	should := "*logger.logger"
	got := reflect.TypeOf(logger).String()

	if got != should {
		t.Errorf("should return %s, got %s", should, got)
	}
}

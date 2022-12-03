package testlib

import (
	"testing"
)

func AssertEqual[T int | int64 | string](t *testing.T, actual T, expected T) {
	if actual != expected {
		t.Errorf("Expected value %d, got %d", expected, actual)
	}
}

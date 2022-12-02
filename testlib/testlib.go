package testlib

import "testing"

func AssertEqualInt(t *testing.T, actual int, expected int) {
	if actual != expected {
		t.Errorf("Expected value %d, got %d", expected, actual)
	}
}

func AssertEqualInt64(t *testing.T, actual int64, expected int64) {
	if actual != expected {
		t.Errorf("Expected value %d, got %d", expected, actual)
	}
}

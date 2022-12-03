package testlib

import (
	"strings"
	"testing"
)

func AssertEqual[T comparable](t *testing.T, actual T, expected T) {
	if actual != expected {
		t.Errorf("Expected value %v, got %v", expected, actual)
	}
}

func AssertStringContains(t *testing.T, str string, substring string) {
	if !strings.Contains(str, substring) {
		t.Errorf("Expected to find substring \"%v\" in \"%v\"", str, substring)
	}
}

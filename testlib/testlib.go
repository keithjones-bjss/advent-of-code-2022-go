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
		t.Errorf("Expected to find substring \"%v\" in \"%v\"", substring, str)
	}
}

func AssertMapContainsKey[K comparable, V any](t *testing.T, m map[K]V, key K) {
	_, present := m[key]
	if !present {
		t.Errorf("Expected to find key \"%v\" in \"%v\"", key, m)
	}
}

func AssertMapDoesNotContainKey[K comparable, V any](t *testing.T, m map[K]V, key K) {
	_, present := m[key]
	if present {
		t.Errorf("Did not expect to find key \"%v\" in \"%v\"", key, m)
	}
}

package testlib

import (
	"testing"
)

func AssertEqual[T comparable](t *testing.T, actual T, expected T) {
	if actual != expected {
		t.Errorf("Expected value %v, got %v", expected, actual)
	}
}

func AssertArrayEqual[T comparable](t *testing.T, actual []T, expected []T) {
	if len(actual) != len(expected) {
		t.Errorf("Expected array %v(len=%v), got %v (len=%v)", expected, len(expected), actual, len(actual))
	}
	for index, _ := range expected {
		if len(actual) <= index {
			t.Errorf("Expected value %v at index %v, got nothing", expected[index], index)
		} else if actual[index] != expected[index] {
			t.Errorf("Expected value %v at index %v, got %v", expected[index], index, actual[index])
		}
	}
}

func AssertContains[T comparable](t *testing.T, array []T, element T) {
	for _, value := range array {
		if value == element {
			return
		}
	}
	t.Errorf("Expected to find item \"%v\" in array \"%v\"", element, array)
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

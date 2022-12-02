package day1

import "testing"

func AssertEqual(t *testing.T, actual int64, expected int64) {
	if actual != expected {
		t.Errorf("Expected value %d, got %d", expected, actual)
	}
}

func TestPart1(t *testing.T) {
	top := Top("../files/day1_test.txt")
	AssertEqual(t, top[0], 10000)
	AssertEqual(t, top[1], 11000)
	AssertEqual(t, top[2], 24000)
}

func TestPart2(t *testing.T) {
	top := Top("../files/day1_test.txt")
	sum := Sum(top)
	AssertEqual(t, sum, 45000)
}

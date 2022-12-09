package day9

import (
	"advent-of-code-2022/testlib"
	"testing"
)

func TestPart1(t *testing.T) {
	result, _ := Run("test.txt")
	testlib.AssertEqual(t, result, 13)
}

func TestPart2(t *testing.T) {
	_, result := Run("test.txt")
	testlib.AssertEqual(t, result, 1)
}

func TestPart2WithLargerGrid(t *testing.T) {
	_, result := Run("test2.txt")
	testlib.AssertEqual(t, result, 36)
}

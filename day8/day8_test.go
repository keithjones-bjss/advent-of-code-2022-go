package day8

import (
	"advent-of-code-2022/testlib"
	"testing"
)

func TestPart1(t *testing.T) {
	result, _ := Run("test.txt")
	testlib.AssertEqual(t, result, 21)
}

func TestPart2(t *testing.T) {
	_, result := Run("test.txt")
	testlib.AssertEqual(t, result, 8)
}

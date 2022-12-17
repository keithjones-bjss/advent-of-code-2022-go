package day16

import (
	"advent-of-code-2022/testlib"
	"testing"
)

func TestPart1(t *testing.T) {
	result, _ := Run("test.txt", false)
	testlib.AssertEqual(t, result, 1651)
}

func TestPart2(t *testing.T) {
	_, result := Run("test.txt", false)
	testlib.AssertEqual(t, result, 1707)
}

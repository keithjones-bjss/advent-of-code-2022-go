package day15

import (
	"advent-of-code-2022/testlib"
	"testing"
)

func TestPart1(t *testing.T) {
	result, _ := RunAtRow("test.txt", 10)
	testlib.AssertEqual(t, result, 26)
}

func TestPart2(t *testing.T) {
	_, result := RunAtRow("test.txt", 10)
	testlib.AssertEqual(t, result, 56000011)
}

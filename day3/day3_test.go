package day3

import (
	"advent-of-code-2022/testlib"
	"testing"
)

func TestPriorityLowerCaseA(t *testing.T) {
	result := Priority("a")
	testlib.AssertEqual(t, result, 1)
}

func TestPriorityLowerCaseZ(t *testing.T) {
	result := Priority("z")
	testlib.AssertEqual(t, result, 26)
}

func TestPriorityUpperCaseA(t *testing.T) {
	result := Priority("A")
	testlib.AssertEqual(t, result, 27)
}

func TestPriorityUpperCaseZ(t *testing.T) {
	result := Priority("Z")
	testlib.AssertEqual(t, result, 52)
}

func TestPart1Score(t *testing.T) {
	result := Part1Score("vJrwpWtwJgWrhcsFMMfFFhFp")
	testlib.AssertEqual(t, result, 16)
}

func TestPart2Score(t *testing.T) {
	result := Part2Score([]string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg"})
	testlib.AssertEqual(t, result, 18)
}

func TestPart1(t *testing.T) {
	result, _ := Run("test.txt")
	testlib.AssertEqual(t, result, 157)
}

func TestPart2(t *testing.T) {
	_, result := Run("test.txt")
	testlib.AssertEqual(t, result, 70)
}

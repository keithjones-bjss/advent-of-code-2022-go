package day4

import (
	"advent-of-code-2022/testlib"
	"testing"
)

func TestExpandIntRangeOneToFour(t *testing.T) {
	result := ExpandIntRange("1-4")
	testlib.AssertArrayEqual(t, result, []int{1, 2, 3, 4})
}

func TestExpandIntRangeTwoToSeven(t *testing.T) {
	result := ExpandIntRange("2-7")
	testlib.AssertArrayEqual(t, result, []int{2, 3, 4, 5, 6, 7})
}

func TestPart1(t *testing.T) {
	result, _ := Run("test.txt")
	testlib.AssertEqual(t, result, 2)
}

func TestPart2(t *testing.T) {
	_, result := Run("test.txt")
	testlib.AssertEqual(t, result, 4)
}

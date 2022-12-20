package day20

import (
	"advent-of-code-2022/testlib"
	"testing"
)

func TestMix(t *testing.T) {
	numbers := []int{1, 2, -3, 3, -2, 0, 4}
	result := Mix(numbers)
	testlib.AssertEqual(t, result[0], 1)
	testlib.AssertEqual(t, result[1], 2)
	testlib.AssertEqual(t, result[2], -3)
	testlib.AssertEqual(t, result[3], 4)
	testlib.AssertEqual(t, result[4], 0)
	testlib.AssertEqual(t, result[5], 3)
	testlib.AssertEqual(t, result[6], -2)
}

func TestWrapNumber(t *testing.T) {
	testlib.AssertEqual(t, WrapNumber(11, 7), 4)
	testlib.AssertEqual(t, WrapNumber(-11, 7), 3)
}

func TestIndexOf(t *testing.T) {
	numbers := []int{1, 2, -3, 4, 0, 3, -2}
	result := IndexOf(numbers, 0)
	testlib.AssertEqual(t, result, 4)
}

func TestValueAt(t *testing.T) {
	numbers := []int{1, 2, -3, 4, 0, 3, -2}
	testlib.AssertEqual(t, ValueAt(numbers, 1004), 4)
	testlib.AssertEqual(t, ValueAt(numbers, 2004), -3)
	testlib.AssertEqual(t, ValueAt(numbers, 3004), 2)
}

func TestPart1(t *testing.T) {
	result, _ := Run("test.txt")
	testlib.AssertEqual(t, result, 3)
}

func TestPart2(t *testing.T) {
	_, result := Run("test.txt")
	testlib.AssertEqual(t, result, -1)
}

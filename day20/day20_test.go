package day20

import (
	"advent-of-code-2022/aoc_library"
	"advent-of-code-2022/testlib"
	"testing"
)

func TestMix(t *testing.T) {
	numbers := []int{1, 2, -3, 3, -2, 0, 4}
	result := Mix(numbers, 1)
	testlib.AssertEqual(t, result[0], 1)
	testlib.AssertEqual(t, result[1], 2)
	testlib.AssertEqual(t, result[2], -3)
	testlib.AssertEqual(t, result[3], 4)
	testlib.AssertEqual(t, result[4], 0)
	testlib.AssertEqual(t, result[5], 3)
	testlib.AssertEqual(t, result[6], -2)
}

func TestMixWithKey(t *testing.T) {
	numbers := []int{1, 2, -3, 3, -2, 0, 4}
	result := Mix(aoc_library.ArrayTranslate(numbers, func(_ int, v int) int { return v * 811589153 }), 1)
	testlib.AssertEqual(t, result[0], 0)
	testlib.AssertEqual(t, result[1], -2434767459)
	testlib.AssertEqual(t, result[2], 3246356612)
	testlib.AssertEqual(t, result[3], -1623178306)
	testlib.AssertEqual(t, result[4], 2434767459)
	testlib.AssertEqual(t, result[5], 1623178306)
	testlib.AssertEqual(t, result[6], 811589153)
}

func TestMixTwiceWithKey(t *testing.T) {
	numbers := []int{1, 2, -3, 3, -2, 0, 4}
	result := Mix(aoc_library.ArrayTranslate(numbers, func(_ int, v int) int { return v * 811589153 }), 2)
	testlib.AssertEqual(t, result[0], 0)
	testlib.AssertEqual(t, result[1], 2434767459)
	testlib.AssertEqual(t, result[2], 1623178306)
	testlib.AssertEqual(t, result[3], 3246356612)
	testlib.AssertEqual(t, result[4], -2434767459)
	testlib.AssertEqual(t, result[5], -1623178306)
	testlib.AssertEqual(t, result[6], 811589153)
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
	testlib.AssertEqual(t, result, 1623178306)
}

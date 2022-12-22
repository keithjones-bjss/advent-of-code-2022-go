package aoc_library

import (
	"advent-of-code-2022/testlib"
	"testing"
)

func TestMapFilter(t *testing.T) {
	m := make(map[rune]int)
	m['a'] = 1
	m['b'] = 1
	m['c'] = 1
	result := MapFilter(m, func(key rune, _ int) bool { return key == 'a' })
	testlib.AssertMapContainsKey(t, result, 'a')
	testlib.AssertMapDoesNotContainKey(t, result, 'b')
	testlib.AssertMapDoesNotContainKey(t, result, 'c')
}

func TestMapKeysToArray(t *testing.T) {
	m := make(map[rune]bool)
	m['a'] = false
	m['b'] = false
	m['c'] = false
	result := MapKeysToArray(m)
	testlib.AssertContains(t, result, 'a')
	testlib.AssertContains(t, result, 'b')
	testlib.AssertContains(t, result, 'c')
}

func TestStringIntersection(t *testing.T) {
	result := AllStringIntersection([]string{"abc", "bdf"})
	testlib.AssertEqual(t, result, "b")
}

func TestIntersection(t *testing.T) {
	result := AllIntersection([][]int{{1, 2, 3}, {3, 4, 5}})
	testlib.AssertArrayEqual(t, result, []int{3})
}

func TestAnyMatch(t *testing.T) {
	result := AnyIntersection([][]int{{1, 2, 3}, {3, 4, 5}, {2, 6, 7, 8}})
	testlib.AssertContains(t, result, 2)
	testlib.AssertContains(t, result, 3)
}

func TestArrayEqualsWithNonMatchingArrays(t *testing.T) {
	result := ArrayEquals([]int{1, 2, 3}, []int{3, 4, 5})
	testlib.AssertEqual(t, result, false)
}

func TestArrayEqualsWithNonEmptyArray(t *testing.T) {
	result := ArrayEquals([]int{1, 2, 3}, []int{})
	testlib.AssertEqual(t, result, false)
}

func TestArrayEqualsWithElementsInDifferentOrder(t *testing.T) {
	result := ArrayEquals([]int{1, 2, 3}, []int{3, 2, 1})
	testlib.AssertEqual(t, result, false)
}

func TestArrayEqualsWithMatchingArrays(t *testing.T) {
	result := ArrayEquals([]int{1, 2, 3}, []int{1, 2, 3})
	testlib.AssertEqual(t, result, true)
}

func TestContainsWithMatchingFirstValue(t *testing.T) {
	result := Contains([]string{"abc", "bdf"}, "abc")
	testlib.AssertEqual(t, result, true)
}

func TestContainsWithMatchingSecondValue(t *testing.T) {
	result := Contains([]string{"abc", "bdf"}, "bdf")
	testlib.AssertEqual(t, result, true)
}

func TestContainsWithNonMatchingValue(t *testing.T) {
	result := Contains([]string{"abc", "bdf"}, "x")
	testlib.AssertEqual(t, result, false)
}

func TestContainsWithEmptyValue(t *testing.T) {
	result := Contains([]string{"abc", "bdf"}, "")
	testlib.AssertEqual(t, result, false)
}

func TestArrayTranslate(t *testing.T) {
	array := []int{1, 2, 3}
	result := ArrayTranslate(array, func(_, element int) []int { return []int{element} })
	testlib.AssertArrayEqual(t, result[0], []int{1})
	testlib.AssertArrayEqual(t, result[1], []int{2})
	testlib.AssertArrayEqual(t, result[2], []int{3})
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

func TestMaxWithRhsLarger(t *testing.T) {
	result := Max(2, 3)
	testlib.AssertEqual(t, result, 3)
}

func TestMaxWithLhsLarger(t *testing.T) {
	result := Max(-7, -11)
	testlib.AssertEqual(t, result, -7)
}

func TestMaxWithSameValues(t *testing.T) {
	result := Max(23, 23)
	testlib.AssertEqual(t, result, 23)
}

func TestMinWithRhsLarger(t *testing.T) {
	result := Min(2, 3)
	testlib.AssertEqual(t, result, 2)
}

func TestMinWithLhsLarger(t *testing.T) {
	result := Min(-7, -11)
	testlib.AssertEqual(t, result, -11)
}

func TestMinWithSameValues(t *testing.T) {
	result := Min(23, 23)
	testlib.AssertEqual(t, result, 23)
}

func TestAbsWithPositiveNumber(t *testing.T) {
	result := Abs(23)
	testlib.AssertEqual(t, result, 23)
}

func TestAbsWithNegativeNumber(t *testing.T) {
	result := Abs(-23)
	testlib.AssertEqual(t, result, 23)
}

func TestAbsWithZero(t *testing.T) {
	result := Abs(0)
	testlib.AssertEqual(t, result, 0)
}

func TestSignWithPositiveNumber(t *testing.T) {
	result := Sign(23)
	testlib.AssertEqual(t, result, 1)
}

func TestSignWithNegativeNumber(t *testing.T) {
	result := Sign(-23)
	testlib.AssertEqual(t, result, -1)
}

func TestSignWithZero(t *testing.T) {
	result := Sign(0)
	testlib.AssertEqual(t, result, 0)
}

func TestModWithPositiveNumber(t *testing.T) {
	result := Mod(11, 7)
	testlib.AssertEqual(t, result, 4)
}

func TestModWithNegativeNumber(t *testing.T) {
	result := Mod(-11, 7)
	testlib.AssertEqual(t, result, 3)
}

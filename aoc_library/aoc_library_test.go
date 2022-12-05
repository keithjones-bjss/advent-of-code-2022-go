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

func TestMapKeysToString(t *testing.T) {
	m := make(map[rune]bool)
	m['a'] = false
	m['b'] = false
	m['c'] = false
	result := MapKeysToString(m)
	testlib.AssertStringContains(t, result, "a")
	testlib.AssertStringContains(t, result, "b")
	testlib.AssertStringContains(t, result, "c")
}

func TestStringIntersection(t *testing.T) {
	result := StringIntersection([]string{"abc", "bdf"})
	testlib.AssertEqual(t, result, "b")
}

func TestIntersection(t *testing.T) {
	result := Intersection([][]int{{1, 2, 3}, {3, 4, 5}})
	testlib.AssertArrayEqual(t, result, []int{3})
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
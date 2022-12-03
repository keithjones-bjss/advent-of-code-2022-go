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

func TestDistinctChars(t *testing.T) {
	result := DistinctChars("aaabbc")
	testlib.AssertEqual(t, result, "abc")
}

func TestMapFilter(t *testing.T) {
	m := make(map[rune]int)
	m[97] = 1
	m[98] = 1
	m[99] = 1
	result := MapFilter(m, func(key rune, _ int) bool { return key == 97 })
	testlib.AssertMapContainsKey(t, result, 97)
	testlib.AssertMapDoesNotContainKey(t, result, 98)
	testlib.AssertMapDoesNotContainKey(t, result, 99)
}

func TestMapKeysToString(t *testing.T) {
	m := make(map[rune]int)
	m[97] = 1
	m[98] = 1
	m[99] = 1
	result := MapKeysToString(m)
	testlib.AssertStringContains(t, result, "a")
	testlib.AssertStringContains(t, result, "b")
	testlib.AssertStringContains(t, result, "c")
}

func TestIntersection(t *testing.T) {
	result := Intersection([]string{"abc", "bdf"})
	testlib.AssertEqual(t, result, "b")
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

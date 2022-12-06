package day6

import (
	"advent-of-code-2022/testlib"
	"testing"
)

func Test1Part1(t *testing.T) {
	result, _ := Run("test1.txt")
	testlib.AssertEqual(t, result, 7)
}

func Test2Part1(t *testing.T) {
	result, _ := Run("test2.txt")
	testlib.AssertEqual(t, result, 5)
}

func Test3Part1(t *testing.T) {
	result, _ := Run("test3.txt")
	testlib.AssertEqual(t, result, 6)
}

func Test4Part1(t *testing.T) {
	result, _ := Run("test4.txt")
	testlib.AssertEqual(t, result, 10)
}

func Test5Part1(t *testing.T) {
	result, _ := Run("test5.txt")
	testlib.AssertEqual(t, result, 11)
}

func Test1Part2(t *testing.T) {
	_, result := Run("test1.txt")
	testlib.AssertEqual(t, result, 19)
}

func Test2Part2(t *testing.T) {
	_, result := Run("test2.txt")
	testlib.AssertEqual(t, result, 23)
}

func Test3Part2(t *testing.T) {
	_, result := Run("test3.txt")
	testlib.AssertEqual(t, result, 23)
}

func Test4Part2(t *testing.T) {
	_, result := Run("test4.txt")
	testlib.AssertEqual(t, result, 29)
}

func Test5Part2(t *testing.T) {
	_, result := Run("test5.txt")
	testlib.AssertEqual(t, result, 26)
}

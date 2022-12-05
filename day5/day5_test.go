package day5

import (
	"advent-of-code-2022/testlib"
	"testing"
)

func TestParseCrates(t *testing.T) {
	result := ParseCrates([]string{"    [D]", "[N] [C]", "[Z] [M] [P]"})
	testlib.AssertArrayEqual(t, result[0], []string{"N", "Z"})
	testlib.AssertArrayEqual(t, result[1], []string{"D", "C", "M"})
	testlib.AssertArrayEqual(t, result[2], []string{"P"})
}

func TestMoveCratesPart1(t *testing.T) {
	stacks := [][]string{{"N", "Z"}, {"D", "C", "M"}, {"P"}}
	result := MoveCratesPart1(stacks, 2, 1, 0)
	testlib.AssertArrayEqual(t, result[0], []string{"C", "D", "N", "Z"})
	testlib.AssertArrayEqual(t, result[1], []string{"M"})
	testlib.AssertArrayEqual(t, result[2], []string{"P"})
}

func TestMoveCratesPart2(t *testing.T) {
	stacks := [][]string{{"N", "Z"}, {"D", "C", "M"}, {"P"}}
	result := MoveCratesPart2(stacks, 2, 1, 0)
	testlib.AssertArrayEqual(t, result[0], []string{"D", "C", "N", "Z"})
	testlib.AssertArrayEqual(t, result[1], []string{"M"})
	testlib.AssertArrayEqual(t, result[2], []string{"P"})
}

func TestPart1(t *testing.T) {
	result, _ := Run("test.txt")
	testlib.AssertEqual(t, result, "CMZ")
}

func TestPart2(t *testing.T) {
	_, result := Run("test.txt")
	testlib.AssertEqual(t, result, "MCD")
}

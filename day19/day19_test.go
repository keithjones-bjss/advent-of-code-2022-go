package day19

import (
	"advent-of-code-2022/testlib"
	"testing"
)

func TestParse(t *testing.T) {
	result := Parse("Blueprint 1: " +
		"Each ore robot costs 4 ore. " +
		"Each clay robot costs 2 ore. " +
		"Each obsidian robot costs 3 ore and 14 clay. " +
		"Each geode robot costs 2 ore and 7 obsidian.")
	testlib.AssertEqual(t, result.id, 1)
	testlib.AssertEqual(t, result.ore, 4)
	testlib.AssertEqual(t, result.clay, 2)
	testlib.AssertEqual(t, result.obsidianOre, 3)
	testlib.AssertEqual(t, result.obsidianClay, 14)
	testlib.AssertEqual(t, result.geodeOre, 2)
	testlib.AssertEqual(t, result.geodeObsidian, 7)
}

func TestWalk24(t *testing.T) {
	blueprint := Blueprint{
		id:            1,
		ore:           4,
		clay:          2,
		obsidianOre:   3,
		obsidianClay:  14,
		geodeOre:      2,
		geodeObsidian: 7,
	}
	result := Walk(24, blueprint, [][][]int{{{1, 0}, {0, 0}, {0, 0}, {0, 0}}})
	testlib.AssertEqual(t, result, 9)
}

func TestWalk32(t *testing.T) {
	blueprint := Blueprint{
		id:            1,
		ore:           4,
		clay:          2,
		obsidianOre:   3,
		obsidianClay:  14,
		geodeOre:      2,
		geodeObsidian: 7,
	}
	result := Walk(32, blueprint, [][][]int{{{1, 0}, {0, 0}, {0, 0}, {0, 0}}})
	testlib.AssertEqual(t, result, 56)
}

func TestPart1(t *testing.T) {
	result, _ := Run("test.txt")
	testlib.AssertEqual(t, result, 33)
}

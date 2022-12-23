package day23

import (
	"advent-of-code-2022/testlib"
	"testing"
)

func TestMove(t *testing.T) {
	grid := [][]rune{
		[]rune(".##"),
		[]rune(".#."),
		[]rune("..."),
		[]rune(".##"),
	}
	MakeDirectionMap()
	result, emptySpaces, _ := Move(grid, 0)
	testlib.AssertEqual(t, emptySpaces, 5)
	testlib.AssertEqual(t, string(result[0]), "##")
	testlib.AssertEqual(t, string(result[1]), "..")
	testlib.AssertEqual(t, string(result[2]), "#.")
	testlib.AssertEqual(t, string(result[3]), ".#")
	testlib.AssertEqual(t, string(result[4]), "#.")
}

func TestMoveFour(t *testing.T) {
	grid := [][]rune{
		[]rune(".##"),
		[]rune(".#."),
		[]rune("..."),
		[]rune(".##"),
	}
	emptySpaces := 0
	MakeDirectionMap()
	result, _, _ := Move(grid, 0)
	result, emptySpaces, _ = Move(result, 1)
	result, emptySpaces, _ = Move(result, 2)
	result, emptySpaces, _ = Move(result, 3)
	testlib.AssertEqual(t, emptySpaces, 25)
	testlib.AssertEqual(t, string(result[0]), "..#..")
	testlib.AssertEqual(t, string(result[1]), "....#")
	testlib.AssertEqual(t, string(result[2]), "#....")
	testlib.AssertEqual(t, string(result[3]), "....#")
	testlib.AssertEqual(t, string(result[4]), ".....")
	testlib.AssertEqual(t, string(result[5]), "..#..")
}

func TestPart1(t *testing.T) {
	result, _ := Run("test.txt")
	testlib.AssertEqual(t, result, 110)
}

func TestPart2(t *testing.T) {
	_, result := Run("test.txt")
	testlib.AssertEqual(t, result, 20)
}

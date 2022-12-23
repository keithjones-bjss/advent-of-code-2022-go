package day22

import (
	"advent-of-code-2022/aoc_library"
	"advent-of-code-2022/testlib"
	"testing"
)

func TestFindCube(t *testing.T) {
	lines := []string{
		"        ...#",
		"        .#..",
		"        #...",
		"        ....",
		"...#.......#",
		"........#...",
		"..#....#....",
		"..........#.",
		"        ...#....",
		"        .....#..",
		"        .#......",
		"        ......#.",
	}
	grid := aoc_library.ArrayTranslate(lines, func(_ int, line string) []rune { return []rune(line) })
	cube := FindCube(grid)

	testlib.AssertEqual(t, cube.width, 4)
	testlib.AssertEqual(t, cube.height, 4)

	testlib.AssertEqual(t, cube.faceGrid[0][0], -1)
	testlib.AssertEqual(t, cube.faceGrid[0][1], -1)
	testlib.AssertEqual(t, cube.faceGrid[0][2], 0)
	testlib.AssertEqual(t, cube.faceGrid[0][3], -1)

	testlib.AssertEqual(t, cube.faceGrid[1][0], 1)
	testlib.AssertEqual(t, cube.faceGrid[1][1], 2)
	testlib.AssertEqual(t, cube.faceGrid[1][2], 3)
	testlib.AssertEqual(t, cube.faceGrid[1][3], -1)

	testlib.AssertEqual(t, cube.faceGrid[2][0], -1)
	testlib.AssertEqual(t, cube.faceGrid[2][1], -1)
	testlib.AssertEqual(t, cube.faceGrid[2][2], 4)
	testlib.AssertEqual(t, cube.faceGrid[2][3], 5)

	testlib.AssertEqual(t, cube.faceGrid[3][0], -1)
	testlib.AssertEqual(t, cube.faceGrid[3][1], -1)
	testlib.AssertEqual(t, cube.faceGrid[3][2], -1)
	testlib.AssertEqual(t, cube.faceGrid[3][3], -1)
}

func TestPart1(t *testing.T) {
	result, _ := Run("test.txt")
	testlib.AssertEqual(t, result, 6032)
}

func TestPart2(t *testing.T) {
	_, result := Run("test.txt")
	testlib.AssertEqual(t, result, 5031)
}

package day9

import (
	"advent-of-code-2022/aoc_library"
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type area struct {
	knots        [10]point
	visitedPart1 map[point]bool
	visitedPart2 map[point]bool
}

var directions = map[string]point{
	"U": {0, 1},
	"D": {0, -1},
	"L": {-1, 0},
	"R": {1, 0},
}

func (p point) Move(dir string) point {
	return point{p.x + directions[dir].x, p.y + directions[dir].y}
}

func (p point) IsTouching(other point) bool {
	return aoc_library.Abs(other.x-p.x) <= 1 && aoc_library.Abs(other.y-p.y) <= 1
}

func (p point) MoveTowards(other point) point {
	if p.IsTouching(other) {
		return p
	}
	p.x += aoc_library.Sign(other.x - p.x)
	p.y += aoc_library.Sign(other.y - p.y)
	return p
}

func Run(filename string) (int, int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Can't open %v: %v", filename, err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	grid := area{}
	grid.knots = [10]point{}
	for n := 0; n <= 9; n++ {
		grid.knots[n] = point{0, 0}
	}
	grid.visitedPart1 = make(map[point]bool)
	grid.visitedPart2 = make(map[point]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			UpdateGrid(&grid, line)
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1 := Part1(grid)
	part2 := Part2(grid)

	return part1, part2
}

func UpdateGrid(grid *area, line string) {
	args := strings.Split(line, " ")
	dir := args[0]
	steps, _ := strconv.Atoi(args[1])
	for count := 1; count <= steps; count++ {
		grid.knots[0] = grid.knots[0].Move(dir)
		for knot := 1; knot < len(grid.knots); knot++ {
			grid.knots[knot] = grid.knots[knot].MoveTowards(grid.knots[knot-1])
		}
		grid.visitedPart1[grid.knots[1]] = true
		grid.visitedPart2[grid.knots[9]] = true
	}
}

func Part1(grid area) int {
	return len(grid.visitedPart1)
}

func Part2(grid area) int {
	return len(grid.visitedPart2)
}

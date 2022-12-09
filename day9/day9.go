package day9

import (
	"bufio"
	"log"
	"math"
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
	"U": point{0, 1},
	"D": point{0, -1},
	"L": point{-1, 0},
	"R": point{1, 0},
}

func (p point) Move(dir string) point {
	return point{p.x + directions[dir].x, p.y + directions[dir].y}
}

func (p point) IsTouching(other point) bool {
	return math.Abs(float64(other.x-p.x)) <= 1 && math.Abs(float64(other.y-p.y)) <= 1
}

func Sign(value int) int {
	if value < 0 {
		return -1
	}
	if value > 0 {
		return 1
	}
	return 0
}

func (p point) MoveTowards(other point) point {
	if p.IsTouching(other) {
		return p
	}
	p.x += Sign(other.x - p.x)
	p.y += Sign(other.y - p.y)
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

package day22

import (
	"advent-of-code-2022/aoc_library"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	Right int = 0
	Down  int = 1
	Left  int = 2
	Up    int = 3
)

type Position struct {
	row    int
	column int
	facing int
}

type Direction struct {
	steps     int
	clockwise bool
}

func Run(filename string) (int, int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Can't open %v: %v", filename, err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var grid [][]rune
	var directions []Direction
	var doneGrid bool

	position := Position{
		row:    0,
		column: 0,
		facing: Right,
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			doneGrid = true
		} else if doneGrid {
			directions = Parse(line)
		} else {
			if len(grid) == 0 {
				position.column = strings.Index(line, ".")
			}
			grid = append(grid, []rune(line))
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1 := Walk(grid, position, directions)
	part2 := 0

	return part1, part2
}

func Parse(line string) []Direction {
	var directions []Direction
	next := line
	for next != "" {
		var steps int
		res, _ := fmt.Sscanf(next, "%d", &steps)
		if res == 1 {
			directions = append(directions, Direction{steps: steps})
			next = next[len(fmt.Sprint(steps)):]
		} else {
			directions = append(directions, Direction{clockwise: next[0] == 'R'})
			next = next[1:]
		}
	}
	return directions
}

func Walk(grid [][]rune, originalPosition Position, directions []Direction) int {
	position := originalPosition
	for _, dir := range directions {
		if dir.steps == 0 {
			if dir.clockwise {
				position.facing++
			} else {
				position.facing--
			}
			position.facing = aoc_library.Mod(position.facing, 4)
		} else {
			for count := 1; count <= dir.steps; count++ {
				position = CalculateNextPosition(grid, position)
			}
		}
	}
	return 1000*(position.row+1) + 4*(position.column+1) + position.facing
}

var offset = []int{0, 1, 0, -1, 0}

func CalculateNextPosition(grid [][]rune, position Position) Position {
	dx := offset[position.facing+1]
	dy := offset[position.facing]
	var next = MoveWithWrap(grid, position, dx, dy)
	if grid[next.row][next.column] != '.' {
		return position
	}
	return next
}

func MoveWithWrap(grid [][]rune, originalPosition Position, dx int, dy int) Position {
	position := originalPosition
	nextRow := aoc_library.Mod(position.row+dy, len(grid))
	nextColumn := aoc_library.Mod(position.column+dx, len(grid[0]))
	for len(grid) <= nextRow || len(grid[nextRow]) <= nextColumn || grid[nextRow][nextColumn] == ' ' {
		nextRow = aoc_library.Mod(nextRow+dy, len(grid))
		nextColumn = aoc_library.Mod(nextColumn+dx, len(grid[0]))
	}
	position.row = nextRow
	position.column = nextColumn
	return position
}

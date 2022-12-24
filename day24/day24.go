package day24

import (
	"advent-of-code-2022/aoc_library"
	"bufio"
	"log"
	"os"
	"strings"
)

func Run(filename string) (int, int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Can't open %v: %v", filename, err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var grid [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			grid = append(grid, []rune(line))
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	area, blizzards := ParseGrid(grid)
	positions := InitialPosition(area.start)
	moves := 0

	blizzards, moves = Move(area, blizzards, positions, moves)
	part1 := moves

	part2start := area.goal
	area.goal = area.start
	area.start = part2start
	positions = InitialPosition(part2start)
	blizzards, moves = Move(area, blizzards, positions, moves)

	part2return := area.goal
	area.goal = area.start
	area.start = part2return
	positions = InitialPosition(part2return)
	_, moves = Move(area, blizzards, positions, moves)

	part2 := moves

	return part1, part2
}

type Blizzard struct {
	position  aoc_library.Point
	direction aoc_library.Point
}

type Area struct {
	height int
	width  int
	start  aoc_library.Point
	goal   aoc_library.Point
}

var directions = []aoc_library.Point{
	{X: 0, Y: -1},
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: -1, Y: 0},
	{X: 0, Y: 0},
}

func ParseGrid(grid [][]rune) (Area, []Blizzard) {
	area := Area{
		height: len(grid),
		width:  len(grid[0]),
	}
	var blizzards []Blizzard
	for y, row := range grid {
		for x, terrain := range row {
			if terrain != '#' {
				if terrain == '.' {
					if y == 0 {
						area.start = aoc_library.Point{X: x, Y: y}
					} else if y == len(grid)-1 {
						area.goal = aoc_library.Point{X: x, Y: y}
					}
				} else {
					blizzard := Blizzard{
						position:  aoc_library.Point{X: x, Y: y},
						direction: directions[strings.Index("^>v<", string(terrain))],
					}
					blizzards = append(blizzards, blizzard)
				}
			}
		}
	}
	return area, blizzards
}

func InitialPosition(point aoc_library.Point) map[aoc_library.Point]bool {
	singleMapEntry := make(map[aoc_library.Point]bool)
	singleMapEntry[point] = true
	return singleMapEntry
}

func Move(area Area, blizzards []Blizzard, positions map[aoc_library.Point]bool, moves int) ([]Blizzard, int) {
	futureBlizzards := UpdateBlizzards(area, blizzards)
	futurePositions := make(map[aoc_library.Point]bool)
	for position := range positions {
		for _, direction := range directions {
			if TryMove(area, futureBlizzards, position, direction) {
				futurePosition := DoMove(position, direction)
				if futurePosition == area.goal {
					return futureBlizzards, moves + 1
				}
				futurePositions[futurePosition] = true
			}
		}
	}
	if len(futurePositions) == 0 {
		return []Blizzard{}, 0
	}
	return Move(area, futureBlizzards, futurePositions, moves+1)
}

func UpdateBlizzards(area Area, blizzards []Blizzard) []Blizzard {
	var futureBlizzards []Blizzard
	for _, blizzard := range blizzards {
		x := blizzard.position.X + blizzard.direction.X
		if x >= area.width-1 {
			x = 1
		}
		if x <= 0 {
			x = area.width - 2
		}
		y := blizzard.position.Y + blizzard.direction.Y
		if y >= area.height-1 {
			y = 1
		}
		if y <= 0 {
			y = area.height - 2
		}
		nextBlizzard := Blizzard{
			position:  aoc_library.Point{X: x, Y: y},
			direction: blizzard.direction,
		}
		futureBlizzards = append(futureBlizzards, nextBlizzard)
	}
	return futureBlizzards
}

func TryMove(area Area, blizzards []Blizzard, position aoc_library.Point, direction aoc_library.Point) bool {
	x := position.X + direction.X
	if x <= 0 || x >= area.width-1 {
		return false
	}
	y := position.Y + direction.Y
	if y < 0 || y >= area.height {
		return false
	}
	if (y == 0 || y == area.height-1) && x != area.start.X && x != area.goal.X {
		return false
	}
	for _, blizzard := range blizzards {
		if x == blizzard.position.X && y == blizzard.position.Y {
			return false
		}
	}
	return true
}

func DoMove(position aoc_library.Point, direction aoc_library.Point) aoc_library.Point {
	return aoc_library.Point{X: position.X + direction.X, Y: position.Y + direction.Y}
}

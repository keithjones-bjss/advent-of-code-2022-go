package day14

import (
	"advent-of-code-2022/aoc_library"
	"bufio"
	"log"
	"os"
	"strconv"
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

	part1 := 0
	part2 := 0

	var walls []aoc_library.Point

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			points := Parse(line)
			walls = append(walls, Explode(points)...)
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	maxY := GetMaxY(walls)

	start := aoc_library.Point{X: 500, Y: 0}

	// Part 1
	filled := MakeMap(walls)
	for part1 = 0; ; part1++ {
		grain := Flow(start, filled, maxY)
		if grain.Y > maxY {
			break
		}
		filled[grain] = true
	}

	// Part 2
	filled = MakeMap(walls)
	for part2 = 0; ; part2++ {
		grain := Flow(start, filled, maxY)
		if grain.Y == 0 {
			break
		}
		filled[grain] = true
	}

	return part1, part2 + 1
}

func MakeMap(walls []aoc_library.Point) map[aoc_library.Point]bool {
	filled := make(map[aoc_library.Point]bool)
	for _, v := range walls {
		filled[v] = true
	}
	return filled
}

func Parse(line string) []aoc_library.Point {
	var points []aoc_library.Point
	parts := strings.Split(line, " -> ")
	for _, v := range parts {
		coords := aoc_library.ArrayTranslate(strings.Split(v, ","), func(_ int, s string) int {
			i, _ := strconv.Atoi(s)
			return i
		})
		points = append(points, aoc_library.Point{X: coords[0], Y: coords[1]})
	}
	return points
}

func Explode(points []aoc_library.Point) []aoc_library.Point {
	var exploded []aoc_library.Point
	var last aoc_library.Point
	for i, p := range points {
		if i != 0 {
			if last.X < p.X {
				for x := last.X; x < p.X; x++ {
					exploded = append(exploded, aoc_library.Point{X: x, Y: p.Y})
				}
			} else if last.X > p.X {
				for x := last.X; x > p.X; x-- {
					exploded = append(exploded, aoc_library.Point{X: x, Y: p.Y})
				}
			} else if last.Y < p.Y {
				for y := last.Y; y < p.Y; y++ {
					exploded = append(exploded, aoc_library.Point{X: p.X, Y: y})
				}
			} else if last.Y > p.Y {
				for y := last.Y; y > p.Y; y-- {
					exploded = append(exploded, aoc_library.Point{X: p.X, Y: y})
				}
			}
		}
		last = p
	}
	exploded = append(exploded, last)
	return exploded
}

func GetMaxY(walls []aoc_library.Point) int {
	maxY := 0
	for _, p := range walls {
		if maxY < p.Y {
			maxY = p.Y
		}
	}
	return maxY
}

func Flow(grain aoc_library.Point, filled map[aoc_library.Point]bool, maxY int) aoc_library.Point {
	next := grain.Move("D")
	if IsBlocked(next, filled) {
		next = grain.Move("D+L")
		if IsBlocked(next, filled) {
			next = grain.Move("D+R")
			if IsBlocked(next, filled) {
				return grain
			}
		}
	}
	if next.Y > maxY {
		return next
	}
	return Flow(next, filled, maxY)
}

func IsBlocked(point aoc_library.Point, filled map[aoc_library.Point]bool) bool {
	_, ok := filled[point]
	return ok
}

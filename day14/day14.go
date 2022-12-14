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

	var walls [][]aoc_library.Point

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			walls = append(walls, Parse(line))
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	maxY := GetMaxY(walls)

	start := aoc_library.Point{X: 500, Y: 0}

	// Part 1
	var sand []aoc_library.Point
	for part1 = 0; ; part1++ {
		grain := Flow(start, walls, sand, maxY)
		if grain.Y > maxY {
			break
		}
		sand = append(sand, grain)
	}

	// Part 2
	var sand2 []aoc_library.Point
	for part2 = 0; ; part2++ {
		grain := Flow(start, walls, sand2, maxY)
		if grain.Y == 0 {
			break
		}
		sand2 = append(sand2, grain)
	}

	return part1, part2 + 1
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

func GetMaxY(walls [][]aoc_library.Point) int {
	maxY := 0
	for _, v := range walls {
		for _, p := range v {
			if maxY < p.Y {
				maxY = p.Y
			}
		}
	}
	return maxY
}

func Flow(
	grain aoc_library.Point,
	walls [][]aoc_library.Point,
	sand []aoc_library.Point,
	maxY int,
) aoc_library.Point {
	next := grain.Move("D")
	if IsBlocked(next, walls, sand) {
		next = grain.Move("D+L")
		if IsBlocked(next, walls, sand) {
			next = grain.Move("D+R")
			if IsBlocked(next, walls, sand) {
				return grain
			}
		}
	}
	if next.Y > maxY {
		return next
	}
	return Flow(next, walls, sand, maxY)
}

func IsBlocked(point aoc_library.Point, walls [][]aoc_library.Point, sand []aoc_library.Point) bool {
	for _, v := range walls {
		var last aoc_library.Point
		for i, next := range v {
			if i != 0 {
				if point.IsBetween(last, next) {
					return true
				}
			}
			last = next
		}
	}
	for _, grain := range sand {
		if point.X == grain.X && point.Y == grain.Y {
			return true
		}
	}
	return false
}

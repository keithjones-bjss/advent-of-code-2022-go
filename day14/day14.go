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

	// Part 1
	filled := MakeMap(walls, maxY)
	for part1 = 0; ; part1++ {
		x, y := Flow(maxY+1, 0, filled, maxY)
		if y > maxY {
			break
		}
		filled[y][x] = 'o'
	}

	// Part 2
	filled = MakeMap(walls, maxY)
	for part2 = 0; ; part2++ {
		x, y := Flow(maxY+1, 0, filled, maxY)
		if y == 0 {
			break
		}
		filled[y][x] = 'o'
	}

	return part1, part2 + 1
}

func MakeMap(walls []aoc_library.Point, maxY int) [][]rune {
	var filled [][]rune
	row := strings.Repeat(".", maxY*2+3)
	for count := 0; count <= maxY+1; count++ {
		filled = append(filled, []rune(row))
	}
	for _, v := range walls {
		filled[v.Y][v.X+maxY-499] = '#'
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

func Flow(lastX int, lastY int, filled [][]rune, maxY int) (int, int) {
	x := lastX
	y := lastY + 1
	if filled[y][x] != '.' {
		x--
		if filled[y][x] != '.' {
			x += 2
			if filled[y][x] != '.' {
				return lastX, lastY
			}
		}
	}
	if y > maxY {
		return x, y
	}
	return Flow(x, y, filled, maxY)
}

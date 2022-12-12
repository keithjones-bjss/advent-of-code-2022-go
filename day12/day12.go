package day12

import (
	"advent-of-code-2022/aoc_library"
	"bufio"
	"log"
	"os"
	"strings"
)

type point struct {
	x int
	y int
}

var directions = map[string]point{
	"U": {0, -1},
	"D": {0, 1},
	"L": {-1, 0},
	"R": {1, 0},
}

func (p point) Move(dir string) point {
	return point{p.x + directions[dir].x, p.y + directions[dir].y}
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

	start := point{0, 0}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			x := strings.Index(line, "S")
			if x >= 0 {
				start.y = len(grid)
				start.x = x
			}
			grid = append(grid, []rune(line))
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1 := Hike(grid, [][]point{{start}})

	var paths [][]point
	for y, v := range grid {
		for x, r := range v {
			if r == 'a' || r == 'S' {
				var start = point{x, y}
				paths = append(paths, []point{start})
			}
		}
	}
	part2 := Hike(grid, paths)

	return part1, part2
}

func Hike(grid [][]rune, paths [][]point) int {
	var visited [][]bool
	for _, v := range grid {
		falseValues := aoc_library.ArrayTranslate(v, func(_ int, _ rune) bool { return false })
		row := append([]bool{}, falseValues...)
		visited = append(visited, row)
	}
	for _, v := range paths {
		visited[v[0].y][v[0].x] = true
	}
	found := Walk(grid, visited, paths)
	result := len(found)
	if result > 0 {
		result = len(found[0]) - 1
	}
	return result
}

func Walk(grid [][]rune, visited [][]bool, paths [][]point) [][]point {
	var nextPaths [][]point
	for _, path := range paths {
		last := path[len(path)-1]
		for dir := range directions {
			next := last.Move(dir)
			if next.x < 0 || next.y < 0 || next.y >= len(grid) || next.x >= len(grid[0]) || visited[next.y][next.x] {
				continue
			}
			lastArea := grid[last.y][last.x]
			nextArea := grid[next.y][next.x]
			if Height(nextArea)-Height(lastArea) <= 1 {
				nextPath := append([]point{}, path...)
				nextPath = append(nextPath, next)
				if grid[next.y][next.x] == 'E' {
					return [][]point{nextPath}
				}
				nextPaths = append(nextPaths, nextPath)
				visited[next.y][next.x] = true
			}
		}
	}
	if len(nextPaths) > 0 {
		return Walk(grid, visited, nextPaths)
	}
	return nextPaths
}

func Height(value rune) rune {
	if value == 'E' {
		return 'z'
	}
	if value == 'S' {
		return 'a'
	}
	return value
}

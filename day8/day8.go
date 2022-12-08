package day8

import (
	"advent-of-code-2022/aoc_library"
	"bufio"
	"log"
	"os"
)

func Run(filename string) (int, int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Can't open %v: %v", filename, err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var grid [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			grid = append(grid,
				aoc_library.ArrayTranslate([]rune(line), func(_ int, v rune) int { return int(v - '0') }))
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1 := Part1(grid)
	part2 := Part2(grid)

	return part1, part2
}

func Part1(grid [][]int) int {
	count := 2*(len(grid)+len(grid[0])) - 4 // all edge trees are visible - don't count corners twice
	for y := 1; y < len(grid)-1; y++ {
		for x := 1; x < len(grid[0])-1; x++ {
			if IsTreeVisible(grid, y, x) {
				count++
			}
		}
	}
	return count
}

func Part2(grid [][]int) int {
	maxScore := 0
	for y := 1; y < len(grid)-1; y++ {
		for x := 1; x < len(grid[0])-1; x++ {
			score := ScenicScore(grid, y, x)
			if maxScore < score {
				maxScore = score
			}
		}
	}
	return maxScore
}

func IsTreeVisible(grid [][]int, y int, x int) bool {
	visible := true
	for n := 0; n < y; n++ {
		if grid[n][x] >= grid[y][x] {
			visible = false
		}
	}
	if visible {
		return visible
	}
	visible = true
	for n := y + 1; n < len(grid); n++ {
		if grid[n][x] >= grid[y][x] {
			visible = false
		}
	}
	if visible {
		return visible
	}
	visible = true
	for n := 0; n < x; n++ {
		if grid[y][n] >= grid[y][x] {
			visible = false
		}
	}
	if visible {
		return visible
	}
	visible = true
	for n := x + 1; n < len(grid[y]); n++ {
		if grid[y][n] >= grid[y][x] {
			visible = false
		}
	}
	return visible
}

func ScenicScore(grid [][]int, y int, x int) int {
	top := 0
	bottom := 0
	left := 0
	right := 0
	for n := y - 1; n >= 0; n-- {
		top++
		if grid[n][x] >= grid[y][x] {
			break
		}
	}
	for n := y + 1; n < len(grid); n++ {
		bottom++
		if grid[n][x] >= grid[y][x] {
			break
		}
	}
	for n := x - 1; n >= 0; n-- {
		left++
		if grid[y][n] >= grid[y][x] {
			break
		}
	}
	for n := x + 1; n < len(grid[y]); n++ {
		right++
		if grid[y][n] >= grid[y][x] {
			break
		}
	}
	return top * bottom * left * right
}

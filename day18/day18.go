package day18

import (
	"advent-of-code-2022/aoc_library"
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point3D struct {
	X int
	Y int
	Z int
}

func Run(filename string) (int, int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Can't open %v: %v", filename, err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var points []Point3D
	var min *Point3D
	var max *Point3D

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			var p Point3D
			_, _ = fmt.Sscanf(line, "%d,%d,%d", &p.X, &p.Y, &p.Z)
			min, max = UpdateGridSize(p, min, max)
			points = append(points, p)
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	size := Point3D{max.X - min.X + 1, max.Y - min.Y + 1, max.Z - min.Z + 1}

	grid := InitializeGrid(size)
	PopulateGrid(points, grid, min)

	part1 := CalculateSurfaceArea(size, grid)

	FillOutsideArea(size, grid)
	outerArea := 2 * (size.X*size.Y + size.X*size.Z + size.Y*size.Z)
	innerArea := CalculateSurfaceArea(size, grid) - outerArea
	part2 := part1 - innerArea

	return part1, part2
}

func UpdateGridSize(p Point3D, min *Point3D, max *Point3D) (*Point3D, *Point3D) {
	if min == nil {
		min = &Point3D{p.X, p.Y, p.Z}
		max = &Point3D{p.X, p.Y, p.Z}
	} else {
		min.X = aoc_library.Min(min.X, p.X)
		min.Y = aoc_library.Min(min.Y, p.Y)
		min.Z = aoc_library.Min(min.Z, p.Z)
		max.X = aoc_library.Max(max.X, p.X)
		max.Y = aoc_library.Max(max.Y, p.Y)
		max.Z = aoc_library.Max(max.Z, p.Z)
	}
	return min, max
}

func InitializeGrid(size Point3D) [][][]bool {
	var grid [][][]bool
	for x := 0; x < size.X; x++ {
		grid = append(grid, [][]bool{})
		for y := 0; y < size.Y; y++ {
			grid[x] = append(grid[x], []bool{})
			for z := 0; z < size.Z; z++ {
				grid[x][y] = append(grid[x][y], false)
			}
		}
	}
	return grid
}

func PopulateGrid(points []Point3D, grid [][][]bool, min *Point3D) {
	for _, p := range points {
		grid[p.X-min.X][p.Y-min.Y][p.Z-min.Z] = true
	}
}

func CalculateSurfaceArea(size Point3D, grid [][][]bool) int {
	area := 0
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			for z := 0; z < size.Z; z++ {
				if grid[x][y][z] {
					area += ExposedFaces(size, grid, x, y, z)
				}
			}
		}
	}
	return area
}

func FillOutsideArea(size Point3D, grid [][][]bool) {
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			FillPoint(size, grid, 0, y, x)
			FillPoint(size, grid, size.X-1, y, x)
			if x != 0 && x != size.X-1 {
				FillPoint(size, grid, x, 0, y)
				FillPoint(size, grid, x, size.Y-1, y)
				if y != 0 && y != size.Y-1 {
					FillPoint(size, grid, x, y, 0)
					FillPoint(size, grid, x, y, size.Z-1)
				}
			}
		}
	}
}

func FillPoint(size Point3D, grid [][][]bool, x int, y int, z int) {
	if grid[x][y][z] {
		return
	}
	grid[x][y][z] = true
	if x > 0 && !grid[x-1][y][z] {
		FillPoint(size, grid, x-1, y, z)
	}
	if x < size.X-1 && !grid[x+1][y][z] {
		FillPoint(size, grid, x+1, y, z)
	}
	if y > 0 && !grid[x][y-1][z] {
		FillPoint(size, grid, x, y-1, z)
	}
	if y < size.Y-1 && !grid[x][y+1][z] {
		FillPoint(size, grid, x, y+1, z)
	}
	if z > 0 && !grid[x][y][z-1] {
		FillPoint(size, grid, x, y, z-1)
	}
	if z < size.Z-1 && !grid[x][y][z+1] {
		FillPoint(size, grid, x, y, z+1)
	}
}

func ExposedFaces(size Point3D, grid [][][]bool, x int, y int, z int) int {
	count := 0
	if x == 0 || !grid[x-1][y][z] {
		count++
	}
	if x == size.X-1 || !grid[x+1][y][z] {
		count++
	}
	if y == 0 || !grid[x][y-1][z] {
		count++
	}
	if y == size.Y-1 || !grid[x][y+1][z] {
		count++
	}
	if z == 0 || !grid[x][y][z-1] {
		count++
	}
	if z == size.Z-1 || !grid[x][y][z+1] {
		count++
	}
	return count
}

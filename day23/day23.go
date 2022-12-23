package day23

import (
	"advent-of-code-2022/aoc_library"
	"bufio"
	"log"
	"os"
	"strings"
)

const (
	North     = '↑'
	NorthEast = '↗'
	South     = '↓'
	SouthEast = '↘'
	East      = '→'
	SouthWest = '↙'
	West      = '←'
	NorthWest = '↖'
)

type Elf struct {
	pos      aoc_library.Point
	next     aoc_library.Point
	dirIndex int
}

type ProposedMoves struct {
	pos   aoc_library.Point
	dir   rune
	count int
}

type Direction struct {
	dir  rune
	move aoc_library.Point
}

var directions = []Direction{
	{dir: North, move: aoc_library.Point{X: 0, Y: -1}},
	{dir: NorthEast, move: aoc_library.Point{X: 1, Y: -1}},
	{dir: East, move: aoc_library.Point{X: 1, Y: 0}},
	{dir: SouthEast, move: aoc_library.Point{X: 1, Y: 1}},
	{dir: South, move: aoc_library.Point{X: 0, Y: 1}},
	{dir: SouthWest, move: aoc_library.Point{X: -1, Y: 1}},
	{dir: West, move: aoc_library.Point{X: -1, Y: 0}},
	{dir: NorthWest, move: aoc_library.Point{X: -1, Y: -1}},
}
var directionOrder = []rune{North, South, West, East}
var directionMap map[rune]int

func Run(filename string) (int, int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Can't open %v: %v", filename, err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	MakeDirectionMap()
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

	// Ten moves
	part1 := 0
	part2 := 0
	for round := 0; part2 == 0; round++ {
		nextGrid, emptySpaces, elvesMoved := Move(grid, round)
		grid = nextGrid
		if round == 9 {
			part1 = emptySpaces
		}
		if elvesMoved == 0 {
			part2 = round + 1
			break
		}
	}

	return part1, part2
}

func MakeDirectionMap() {
	directionMap = make(map[rune]int)
	directionMap[North] = 0
	directionMap[NorthEast] = 1
	directionMap[East] = 2
	directionMap[SouthEast] = 3
	directionMap[South] = 4
	directionMap[SouthWest] = 5
	directionMap[West] = 6
	directionMap[NorthWest] = 7
}

func Move(grid [][]rune, round int) ([][]rune, int, int) {
	var nextGrid [][]rune
	// Generate empty grid and locate elves
	height := len(grid)
	width := len(grid[0])
	elves, nextGrid := PrepareMoves(grid, width, height, nextGrid)
	// Propose moves
	proposedMoves := make(map[aoc_library.Point]int)
	for i, elf := range elves {
		next := ProposeMove(grid, round, width, height, elf)
		elves[i].next = next
		proposedMoves[next]++
		//log.Printf("Proposed move %v -> %v count %v", elf.pos, next, proposedMoves[next])
	}
	// Perform all moves
	result, elvesMoved := PerformMoves(nextGrid, elves, proposedMoves)
	return result, (len(result) * len(result[0])) - len(elves), elvesMoved
}

func PrepareMoves(grid [][]rune, width int, height int, nextGrid [][]rune) ([]Elf, [][]rune) {
	emptyRow := strings.Repeat(".", width+2)
	var elves []Elf
	for row := 0; row < height+2; row++ {
		nextGrid = append(nextGrid, []rune(emptyRow))
		if row < len(grid) {
			for column, v := range grid[row] {
				if v != '.' {
					if v == '#' {
						v = '^'
					}
					//dirIndex := directionMap[v]
					//elves = append(elves, Elf{pos: aoc_library.Point{X: column, Y: row}, dirIndex: dirIndex})
					elves = append(elves, Elf{pos: aoc_library.Point{X: column, Y: row}})
				}
			}
		}
	}
	return elves, nextGrid
}

func ProposeMove(grid [][]rune, round int, width int, height int, elf Elf) aoc_library.Point {
	// Check elf's surroundings
	var emptySpaces []bool
	allEmpty := true
	for _, dir := range directions {
		isEmpty := IsEmptySpace(grid, width, height, elf.pos, dir.move)
		allEmpty = allEmpty && isEmpty
		emptySpaces = append(emptySpaces, isEmpty)
	}
	next := aoc_library.Point{X: elf.pos.X + 1, Y: elf.pos.Y + 1}
	if allEmpty {
		return next
	}
	// Propose elf's move
	//found := false
	for count := 0; count <= 3; count++ {
		dirIndex := directionMap[directionOrder[(count+round)%4]]
		dir := directions[dirIndex]
		left := aoc_library.Mod(dirIndex-1, 8)
		right := aoc_library.Mod(dirIndex+1, 8)
		if emptySpaces[dirIndex] && emptySpaces[left] && emptySpaces[right] {
			next.Y += dir.move.Y
			next.X += dir.move.X
			//elves[i].dirIndex = (count + elf.dirIndex) % 4
			//log.Printf("Elf at row %v col %v proposes to move %v.", elf.pos.Y, elf.pos.X, string(dir.dir))
			//found = true
			break
			//} else {
			//	log.Printf("Elf at row %v col %v can't move %v.", elf.pos.Y, elf.pos.X, string(dir.dir))
		}
	}
	//if !found {
	//	log.Printf("Elf at row %v col %v cannot move in any cardinal direction.", elf.pos.Y, elf.pos.X)
	//}
	//} else {
	//	log.Printf("Elf at row %v col %v is alone, so chooses not to move.", elf.pos.Y, elf.pos.X)
	return next
}

func PerformMoves(nextGrid [][]rune, elves []Elf, proposedMoves map[aoc_library.Point]int) ([][]rune, int) {
	// Attempt moves
	elvesMoved := 0
	for _, elf := range elves {
		if proposedMoves[elf.next] <= 1 {
			nextGrid[elf.next.Y][elf.next.X] = '#'
			if elf.next.X != elf.pos.X+1 || elf.next.Y != elf.pos.Y+1 {
				elvesMoved++
			}
			//log.Printf("Actual move %v -> %v", elf.pos, elf.next)
		} else {
			nextGrid[elf.pos.Y+1][elf.pos.X+1] = '#'
			//log.Printf("Actual move %v -> {%v %v}", elf.pos, elf.pos.X+1, elf.pos.Y+1)
		}
	}
	// Trim grid if possible
	result := Trim(nextGrid)
	return result, elvesMoved
}

func IsEmptySpace(grid [][]rune, width int, height int, point aoc_library.Point, dir aoc_library.Point) bool {
	return point.Y+dir.Y < 0 ||
		point.Y+dir.Y >= height ||
		point.X+dir.X < 0 ||
		point.X+dir.X >= width ||
		grid[point.Y+dir.Y][point.X+dir.X] == '.'
}

func Trim(grid [][]rune) [][]rune {
	skipTop := 0
	for !aoc_library.Contains(grid[skipTop], '#') {
		skipTop++
	}
	skipBottom := len(grid) - 1
	for !aoc_library.Contains(grid[skipBottom], '#') {
		skipBottom--
	}
	minIndex := len(grid[0]) - 1
	maxIndex := 0
	for index := skipTop; index < skipBottom; index++ {
		firstIndex := aoc_library.IndexOf(grid[index], '#')
		if firstIndex >= 0 {
			if minIndex > firstIndex {
				minIndex = firstIndex
			}
			lastIndex := aoc_library.LastIndexOf(grid[index], '#')
			if maxIndex < lastIndex {
				maxIndex = lastIndex
			}
		}
	}
	return aoc_library.ArrayTranslate(grid[skipTop:skipBottom+1], func(index int, value []rune) []rune {
		return value[minIndex : maxIndex+1]
	})
}

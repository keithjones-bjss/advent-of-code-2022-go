package day17

import (
	"bufio"
	"log"
	"os"
)

var blockStrings = [][]string{
	{"@@@@"},
	{".@.", "@@@", ".@."},
	{"@@@", "..@", "..@"},
	{"@", "@", "@", "@"},
	{"@@", "@@"},
}

func Run(filename string) (int, int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Can't open %v: %v", filename, err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var jets string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			jets = line
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1 := SolveTower(jets, 2022)
	part2 := SolveTower(jets, 1000000000000)

	return part1, part2
}

func GenerateBlocks() [][][]rune {
	var blocks [][][]rune
	for i := range blockStrings {
		var block [][]rune
		for j := range blockStrings[i] {
			block = append(block, []rune(blockStrings[i][j]))
		}
		blocks = append(blocks, block)
	}
	return blocks
}

type TowerHistory struct {
	height  int
	count   int
	jet     int
	looping bool
}

func SolveTower(jets string, rocks int) int {
	var towerHistory []TowerHistory
	var tower [][]rune
	height := 0
	addedHeight := 0
	jetCount := 0
	blocks := GenerateBlocks()
	for count := 0; count < rocks; count++ {
		x := 2
		y := height + 3
		block := blocks[count%len(blocks)]
		w := len(block[0])
		rest := false
		for !rest {
			ox := x
			oy := y
			// Move block left or right
			jet := jets[jetCount%len(jets)]
			jetCount++
			x = MoveBlockSideways(jet, x, w)
			if CheckCollision(x, y, block, tower, height) {
				x = ox
			}
			// Drop block
			y--
			rest = CheckCollision(x, y, block, tower, height)
			if rest {
				// Update tower
				y = oy
				for len(tower) < y+len(block) {
					tower = append(tower, []rune("......."))
				}
				height = len(tower)
				for dy := 0; dy < len(block); dy++ {
					if y+dy < height {
						for dx := 0; dx < len(block[dy]); dx++ {
							if block[dy][dx] != '.' {
								tower[y+dy][x+dx] = '#'
							}
						}
					}
				}
				// Iteration detection
				currentState := TowerHistory{
					len(tower),
					count % len(blocks),
					jetCount % len(jets),
					false,
				}
				lastLoop := -1
				for i, v := range towerHistory {
					if v.count == currentState.count && v.jet == currentState.jet {
						currentState.looping = true
						lastLoop = i
					}
				}
				currentPosition := len(towerHistory)
				towerHistory = append(towerHistory, currentState)
				iterationSize := currentPosition - lastLoop
				previousIteration := lastLoop - iterationSize
				if previousIteration >= 0 {
					heightDifference := currentState.height - towerHistory[lastLoop].height
					previousIterationMatches := towerHistory[previousIteration].count == currentState.count && towerHistory[previousIteration].jet == currentState.jet
					heightsMatch := heightDifference == towerHistory[lastLoop].height-towerHistory[previousIteration].height
					if previousIterationMatches && heightsMatch {
						if rocks-count >= iterationSize {
							// Simulate further complete iterations
							remainingIterations := (rocks - count) / iterationSize
							addedRocks := iterationSize * remainingIterations
							addedHeight = heightDifference * remainingIterations
							count += addedRocks
						}
					}
				}
			}
		}
	}

	return len(tower) + addedHeight
}

func MoveBlockSideways(jet uint8, x int, w int) int {
	if jet == '>' && x+w < 7 {
		x++
	}
	if jet == '<' && x > 0 {
		x--
	}
	return x
}

func CheckCollision(x int, y int, block [][]rune, tower [][]rune, height int) bool {
	if y < 0 {
		return true
	}
	if y < height {
		for dy := 0; dy < len(block); dy++ {
			if y+dy < height {
				for dx := 0; dx < len(block[dy]); dx++ {
					if block[dy][dx] != '.' && tower[y+dy][x+dx] != '.' {
						return true
					}
				}
			}
		}
	}
	return false
}

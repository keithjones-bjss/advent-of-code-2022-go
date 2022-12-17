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

	var blocks [][][]rune
	for i := range blockStrings {
		var block [][]rune
		for j := range blockStrings[i] {
			block = append(block, []rune(blockStrings[i][j]))
		}
		blocks = append(blocks, block)
	}

	part1 := SolveTower(blocks, jets, 2022)
	//part2 := SolveTower(blocks, jets, 1000000000000)
	part2 := 0

	return part1, part2
}

func SolveTower(blocks [][][]rune, jets string, rocks int) int {
	var tower [][]rune
	height := 0
	jetCount := 0
	debug := false
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
								if debug {
									log.Printf("%v,%v -> %v,%v", dy, dx, y+dy, x+dx)
								}
								tower[y+dy][x+dx] = '#'
							}
						}
					}
				}
			}
			if debug {
				log.Printf("%v-%v,%v-%v %v %v-%v,%v-%v [%v]",
					ox, ox+w-1, oy, oy+len(block)-1, string(jet), x, x+w-1, y, y+len(block)-1, rest)
			}
		}
		if debug {
			for idx := len(tower) - 1; idx >= 0; idx-- {
				log.Printf("|%v|", string(tower[idx]))
			}
			log.Printf("+-------+")
			log.Printf("")
		}
	}

	return len(tower)
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

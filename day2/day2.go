package day2

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Run(filename string) (int, int) {
	part1 := 0
	part2 := 0

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Can't open %s: %s", filename, err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			part1 += Part1Score(line)
			part2 += Part2Score(line)
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return part1, part2
}

func Part1Score(round string) int {
	moves := strings.Split(round, " ")
	theirs := int(moves[0][0]) - int('A')
	ours := int(moves[1][0]) - int('X')
	outcome := (4 + ours - theirs) % 3 // 0 = loss, 1 = draw, 2 = win
	score := ours + outcome*3 + 1
	return score
}

func Part2Score(round string) int {
	moves := strings.Split(round, " ")
	theirs := int(moves[0][0]) - int('A')
	outcome := int(moves[1][0]) - int('X') // 0 = loss, 1 = draw, 2 = win
	ours := (outcome + theirs + 2) % 3     // 0 = rock, 1 = paper, 2 = scissors
	score := ours + outcome*3 + 1
	return score
}

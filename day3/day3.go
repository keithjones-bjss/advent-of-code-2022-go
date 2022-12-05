package day3

import (
	"advent-of-code-2022/aoc_library"
	"bufio"
	"log"
	"os"
)

func Run(filename string) (int, int) {
	part1 := 0
	part2 := 0

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Can't open %v: %v", filename, err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)
	var batch []string
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			part1 += Part1Score(line)
			batch = append(batch, line)
			if len(batch) == 3 {
				part2 += Part2Score(batch)
				batch = nil
			}
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return part1, part2
}

func Part1Score(rucksack string) int {
	size := len(rucksack)
	half := size / 2
	compartments := []string{rucksack[0:half], rucksack[half:size]}
	result := aoc_library.StringIntersection(compartments)
	return Priority(result)
}

func Part2Score(rucksacks []string) int {
	result := aoc_library.StringIntersection(rucksacks)
	return Priority(result)
}

func Priority(str string) int {
	char := int(str[0])
	if char >= 97 && char <= 122 {
		return char - 96 // a-z => 1-26
	}
	if char >= 65 && char <= 90 {
		return char - 38 // A-Z => 27-52
	}
	return 0
}

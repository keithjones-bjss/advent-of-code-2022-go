package day6

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

	part1 := 0
	part2 := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			part1 = Part1(line)
			part2 = Part2(line)
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return part1, part2
}

func Part1(line string) int {
	return FindStartOfMessageMarker(line, 4)
}

func Part2(line string) int {
	return FindStartOfMessageMarker(line, 14)
}

func FindStartOfMessageMarker(line string, length int) int {
	for index := length; index <= len(line); index++ {
		substr := []rune(line[index-length : index])
		split := aoc_library.ArrayTranslate(substr, func(_ int, element rune) []rune { return []rune{element} })
		result := aoc_library.AnyIntersection(split)
		if len(result) == 0 {
			return index
		}
	}
	return 0
}

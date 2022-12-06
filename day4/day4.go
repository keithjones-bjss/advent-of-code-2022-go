package day4

import (
	"advent-of-code-2022/aoc_library"
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
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
	lineCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineCount++
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

func Part1Score(line string) int {
	intRangePair := strings.Split(line, ",")
	var expandedRanges [][]int
	for _, intRange := range intRangePair {
		expandedRanges = append(expandedRanges, ExpandIntRange(intRange))
	}
	intersection := aoc_library.AllIntersection(expandedRanges)
	sort.Ints(intersection)
	if aoc_library.ArrayContains(expandedRanges, intersection) {
		return 1
	}
	return 0
}

func Part2Score(line string) int {
	intRangePair := strings.Split(line, ",")
	var expandedRanges [][]int
	for _, intRange := range intRangePair {
		expandedRanges = append(expandedRanges, ExpandIntRange(intRange))
	}
	intersection := aoc_library.AllIntersection(expandedRanges)
	if len(intersection) > 0 {
		return 1
	}
	return 0
}

func ExpandIntRange(intRange string) []int {
	boundary := strings.Split(intRange, "-")
	first, _ := strconv.Atoi(boundary[0])
	last, _ := strconv.Atoi(boundary[1])
	var result []int
	for value := first; value <= last; value++ {
		result = append(result, value)
	}
	return result
}

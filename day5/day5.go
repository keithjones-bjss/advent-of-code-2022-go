package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Run(filename string) (string, string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Can't open %v: %v", filename, err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)
	var stackLines []string
	var part1stacks [][]string
	var part2stacks [][]string
	for scanner.Scan() {
		line := scanner.Text()
		if len(part1stacks) > 0 {
			numberOfCrates := 0
			source := 0
			target := 0
			_, err = fmt.Sscanf(line, "move %d from %d to %d", &numberOfCrates, &source, &target)
			if err != nil {
				log.Fatalf("Can't parse line %v: %v", line, err)
			}
			part1stacks = MoveCratesPart1(part1stacks, numberOfCrates, source-1, target-1)
			part2stacks = MoveCratesPart2(part2stacks, numberOfCrates, source-1, target-1)
		} else {
			if line == "" {
				part1stacks = ParseCrates(stackLines[:len(stackLines)-1])
				part2stacks = ParseCrates(stackLines[:len(stackLines)-1])
			} else {
				stackLines = append(stackLines, line)
			}
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1 := ""
	for _, stack := range part1stacks {
		part1 += stack[0]
	}

	part2 := ""
	for _, stack := range part2stacks {
		part2 += stack[0]
	}

	return part1, part2
}

func ParseCrates(stackLines []string) [][]string {
	var stacks [][]string
	for _, line := range stackLines {
		for index := 0; index < len(line); index += 4 {
			row := index / 4
			if len(stacks) <= row {
				stacks = append(stacks, []string{})
			}
			if index+1 < len(line) {
				value := line[index+1]
				if value >= 'A' && value <= 'Z' {
					stacks[row] = append(stacks[row], string(value))
				}
			}
		}
	}
	return stacks
}

func MoveCratesPart1(stacks [][]string, numberOfCrates int, source int, target int) [][]string {
	for n := 1; n <= numberOfCrates; n++ {
		if len(stacks[source]) > 0 {
			item := stacks[source][0]
			stacks[source] = stacks[source][1:]
			stacks[target] = append([]string{item}, stacks[target]...)
		}
	}
	return stacks
}

func MoveCratesPart2(stacks [][]string, numberOfCrates int, source int, target int) [][]string {
	if len(stacks[source]) >= numberOfCrates {
		items := append([]string{}, stacks[source][:numberOfCrates]...)
		stacks[source] = stacks[source][numberOfCrates:]
		stacks[target] = append(items, stacks[target]...)
	}
	return stacks
}

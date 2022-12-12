package day11

import (
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
			// do something
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return part1, part2
}

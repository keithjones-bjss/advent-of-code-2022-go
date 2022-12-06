package day1

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func Top(filename string) []int64 {
	current := int64(0)
	top := []int64{0, 0, 0}

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
		if line == "" {
			if top[0] < current {
				top[0] = current
				sort.Slice(top, func(a, b int) bool {
					return top[a] < top[b]
				})
			}
			current = int64(0)
		} else {
			value, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				log.Fatalf("cannot parse %s: %s", line, err)
			}
			current += value
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if top[0] < current {
		top[0] = current
		sort.Slice(top, func(a, b int) bool {
			return top[a] < top[b]
		})
	}

	return top
}

func Sum(array []int64) int64 {
	sum := int64(0)
	for _, value := range array {
		sum += value
	}
	return sum
}

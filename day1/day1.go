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

	file, error := os.Open(filename)
	if error != nil {
		log.Fatalf("Can't open %s: %s", filename, error)
	}
	defer file.Close()

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
			value, error := strconv.ParseInt(line, 10, 64)
			if error != nil {
				log.Fatalf("cannot parse %s: %s", line, error)
			}
			current += value
		}
	}
	if error = scanner.Err(); error != nil {
		log.Fatal(error)
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

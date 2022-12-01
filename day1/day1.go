package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func Run() {
	current := int64(0)
	top := []int64{0, 0, 0}

	filename := "files/day1_input.txt"
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

	if top[0] < current {
		top[0] = current
		sort.Slice(top, func(a, b int) bool {
			return top[a] < top[b]
		})
	}

	sum := int64(0)
	for _, value := range top {
		sum += value
	}

	fmt.Printf("Part 1: %d\n", top[2]) // Part 1
	fmt.Printf("Part 2: %d\n", sum)    // Part 2

	if error = scanner.Err(); error != nil {
		log.Fatal(error)
	}
}

package main

import (
	"advent-of-code-2022/day1"
	"fmt"
)

func main() {
	top := day1.Top("files/day1_input.txt")
	sum := day1.Sum(top)

	fmt.Printf("Part 1: %d\n", top[2])
	fmt.Printf("Part 2: %d\n", sum)
}

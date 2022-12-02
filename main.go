package main

import (
	"advent-of-code-2022/day1"
	"advent-of-code-2022/day2"
	"fmt"
)

func RunDay1() {
	top := day1.Top("files/day1_input.txt")
	sum := day1.Sum(top)

	fmt.Printf("Day 1  Part 1: %d\n", top[2])
	fmt.Printf("Day 1  Part 2: %d\n", sum)
}

func RunDay2() {
	part1, part2 := day2.Run("files/day2_input.txt")
	fmt.Printf("Day 2  Part 1: %d\n", part1)
	fmt.Printf("Day 2  Part 2: %d\n", part2)
	//fmt.Printf("Part 2: %d\n", sum)
}

func main() {
	RunDay1()
	RunDay2()
}

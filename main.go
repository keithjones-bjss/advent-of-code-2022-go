package main

import (
	"advent-of-code-2022/day1"
	"advent-of-code-2022/day2"
	"advent-of-code-2022/day3"
	"fmt"
)

func RunDay1() {
	top := day1.Top("day1/input.txt")
	sum := day1.Sum(top)

	fmt.Printf("Day 1  Part 1: %v\n", top[2])
	fmt.Printf("Day 1  Part 2: %v\n", sum)
}

func RunDay2() {
	part1, part2 := day2.Run("day2/input.txt")
	fmt.Printf("Day 2  Part 1: %v\n", part1)
	fmt.Printf("Day 2  Part 2: %v\n", part2)
}

func RunDay3() {
	part1, part2 := day3.Run("day3/input.txt")
	fmt.Printf("Day 3  Part 1: %v\n", part1)
	fmt.Printf("Day 3  Part 2: %v\n", part2)
}

func main() {
	RunDay1()
	RunDay2()
	RunDay3()
}

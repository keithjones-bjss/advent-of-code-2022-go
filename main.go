package main

import (
	"advent-of-code-2022/day1"
	"advent-of-code-2022/day10"
	"advent-of-code-2022/day11"
	"advent-of-code-2022/day2"
	"advent-of-code-2022/day3"
	"advent-of-code-2022/day4"
	"advent-of-code-2022/day5"
	"advent-of-code-2022/day6"
	"advent-of-code-2022/day7"
	"advent-of-code-2022/day8"
	"advent-of-code-2022/day9"
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

func RunDay4() {
	part1, part2 := day4.Run("day4/input.txt")
	fmt.Printf("Day 4  Part 1: %v\n", part1)
	fmt.Printf("Day 4  Part 2: %v\n", part2)
}

func RunDay5() {
	part1, part2 := day5.Run("day5/input.txt")
	fmt.Printf("Day 5  Part 1: %v\n", part1)
	fmt.Printf("Day 5  Part 2: %v\n", part2)
}

func RunDay6() {
	part1, part2 := day6.Run("day6/input.txt")
	fmt.Printf("Day 6  Part 1: %v\n", part1)
	fmt.Printf("Day 6  Part 2: %v\n", part2)
}

func RunDay7() {
	part1, part2 := day7.Run("day7/input.txt")
	fmt.Printf("Day 7  Part 1: %v\n", part1)
	fmt.Printf("Day 7  Part 2: %v\n", part2)
}

func RunDay8() {
	part1, part2 := day8.Run("day8/input.txt")
	fmt.Printf("Day 8  Part 1: %v\n", part1)
	fmt.Printf("Day 8  Part 2: %v\n", part2)
}

func RunDay9() {
	part1, part2 := day9.Run("day9/input.txt")
	fmt.Printf("Day 9  Part 1: %v\n", part1)
	fmt.Printf("Day 9  Part 2: %v\n", part2)
}

func RunDay10() {
	part1, part2 := day10.Run("day10/input.txt")
	fmt.Printf("Day 10 Part 1: %v\n", part1)
	fmt.Printf("Day 10 Part 2\n%v\n", day10.Stringify(part2))
}

func RunDay11() {
	part1, part2 := day11.Run("day11/input.txt")
	fmt.Printf("Day 11 Part 1: %v\n", part1)
	fmt.Printf("Day 11 Part 2: %v\n", part2)
}

func main() {
	RunDay1()
	RunDay2()
	RunDay3()
	RunDay4()
	RunDay5()
	RunDay6()
	RunDay7()
	RunDay8()
	RunDay9()
	RunDay10()
	RunDay11()
}

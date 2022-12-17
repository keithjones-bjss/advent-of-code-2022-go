package main

import (
	"advent-of-code-2022/day1"
	"advent-of-code-2022/day10"
	"advent-of-code-2022/day11"
	"advent-of-code-2022/day12"
	"advent-of-code-2022/day13"
	"advent-of-code-2022/day14"
	"advent-of-code-2022/day15"
	"advent-of-code-2022/day16"
	"advent-of-code-2022/day17"
	"advent-of-code-2022/day2"
	"advent-of-code-2022/day3"
	"advent-of-code-2022/day4"
	"advent-of-code-2022/day5"
	"advent-of-code-2022/day6"
	"advent-of-code-2022/day7"
	"advent-of-code-2022/day8"
	"advent-of-code-2022/day9"
	"fmt"
	"time"
)

func Day1() {
	top := day1.Top("day1/input.txt")
	sum := day1.Sum(top)

	fmt.Printf("Day 1  Part 1: %v\n", top[2])
	fmt.Printf("Day 1  Part 2: %v\n", sum)
}

func Day2() {
	part1, part2 := day2.Run("day2/input.txt")
	fmt.Printf("Day 2  Part 1: %v\n", part1)
	fmt.Printf("Day 2  Part 2: %v\n", part2)
}

func Day3() {
	part1, part2 := day3.Run("day3/input.txt")
	fmt.Printf("Day 3  Part 1: %v\n", part1)
	fmt.Printf("Day 3  Part 2: %v\n", part2)
}

func Day4() {
	part1, part2 := day4.Run("day4/input.txt")
	fmt.Printf("Day 4  Part 1: %v\n", part1)
	fmt.Printf("Day 4  Part 2: %v\n", part2)
}

func Day5() {
	part1, part2 := day5.Run("day5/input.txt")
	fmt.Printf("Day 5  Part 1: %v\n", part1)
	fmt.Printf("Day 5  Part 2: %v\n", part2)
}

func Day6() {
	part1, part2 := day6.Run("day6/input.txt")
	fmt.Printf("Day 6  Part 1: %v\n", part1)
	fmt.Printf("Day 6  Part 2: %v\n", part2)
}

func Day7() {
	part1, part2 := day7.Run("day7/input.txt")
	fmt.Printf("Day 7  Part 1: %v\n", part1)
	fmt.Printf("Day 7  Part 2: %v\n", part2)
}

func Day8() {
	part1, part2 := day8.Run("day8/input.txt")
	fmt.Printf("Day 8  Part 1: %v\n", part1)
	fmt.Printf("Day 8  Part 2: %v\n", part2)
}

func Day9() {
	part1, part2 := day9.Run("day9/input.txt")
	fmt.Printf("Day 9  Part 1: %v\n", part1)
	fmt.Printf("Day 9  Part 2: %v\n", part2)
}

func Day10() {
	part1, part2 := day10.Run("day10/input.txt")
	fmt.Printf("Day 10 Part 1: %v\n", part1)
	fmt.Printf("Day 10 Part 2\n%v\n", day10.Stringify(part2))
}

func Day11() {
	part1, part2 := day11.Run("day11/input.txt")
	fmt.Printf("Day 11 Part 1: %v\n", part1)
	fmt.Printf("Day 11 Part 2: %v\n", part2)
}

func Day12() {
	part1, part2 := day12.Run("day12/input.txt")
	fmt.Printf("Day 12 Part 1: %v\n", part1)
	fmt.Printf("Day 12 Part 2: %v\n", part2)
}

func Day13() {
	part1, part2 := day13.Run("day13/input.txt")
	fmt.Printf("Day 13 Part 1: %v\n", part1)
	fmt.Printf("Day 13 Part 2: %v\n", part2)
}

func Day14() {
	part1, part2 := day14.Run("day14/input.txt")
	fmt.Printf("Day 14 Part 1: %v\n", part1)
	fmt.Printf("Day 14 Part 2: %v\n", part2)
}

func Day15() {
	part1, part2 := day15.Run("day15/input.txt")
	fmt.Printf("Day 15 Part 1: %v\n", part1)
	fmt.Printf("Day 15 Part 2: %v\n", part2)
}

func Day16() {
	part1, part2 := day16.Run("day16/input.txt")
	fmt.Printf("Day 16 Part 1: %v\n", part1)
	fmt.Printf("Day 16 Part 2: %v\n", part2)
}

func Day17() {
	part1, part2 := day17.Run("day17/input.txt")
	fmt.Printf("Day 17 Part 1: %v\n", part1)
	fmt.Printf("Day 17 Part 2: %v\n", part2)
}

func RunTimed(function func()) {
	started := time.Now()
	function()
	elapsed := time.Since(started)
	fmt.Printf("Completed in %v\n", elapsed)
}

var skipDay16 = true

func main() {
	started := time.Now()
	RunTimed(Day1)
	RunTimed(Day2)
	RunTimed(Day3)
	RunTimed(Day4)
	RunTimed(Day5)
	RunTimed(Day6)
	RunTimed(Day7)
	RunTimed(Day8)
	RunTimed(Day9)
	RunTimed(Day10)
	RunTimed(Day11)
	RunTimed(Day12)
	RunTimed(Day13)
	RunTimed(Day14)
	RunTimed(Day15)
	if skipDay16 {
		fmt.Print("Day 16: ** SKIPPED **\n")
	} else {
		RunTimed(Day16)
	}
	RunTimed(Day17)
	elapsed := time.Since(started)
	fmt.Printf("Completed all days in %v\n", elapsed)
}

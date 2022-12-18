package main

import (
	"advent-of-code-2022/aoc_library"
	"advent-of-code-2022/day1"
	"advent-of-code-2022/day10"
	"advent-of-code-2022/day11"
	"advent-of-code-2022/day12"
	"advent-of-code-2022/day13"
	"advent-of-code-2022/day14"
	"advent-of-code-2022/day15"
	"advent-of-code-2022/day16"
	"advent-of-code-2022/day17"
	"advent-of-code-2022/day18"
	"advent-of-code-2022/day2"
	"advent-of-code-2022/day3"
	"advent-of-code-2022/day4"
	"advent-of-code-2022/day5"
	"advent-of-code-2022/day6"
	"advent-of-code-2022/day7"
	"advent-of-code-2022/day8"
	"advent-of-code-2022/day9"
	"flag"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"
)

func Day1(_ bool) {
	top := day1.Top("day1/input.txt")
	sum := day1.Sum(top)

	fmt.Printf("Day 1  Part 1: %v\n", top[2])
	fmt.Printf("Day 1  Part 2: %v\n", sum)
}

func Day2(_ bool) {
	part1, part2 := day2.Run("day2/input.txt")
	fmt.Printf("Day 2  Part 1: %v\n", part1)
	fmt.Printf("Day 2  Part 2: %v\n", part2)
}

func Day3(_ bool) {
	part1, part2 := day3.Run("day3/input.txt")
	fmt.Printf("Day 3  Part 1: %v\n", part1)
	fmt.Printf("Day 3  Part 2: %v\n", part2)
}

func Day4(_ bool) {
	part1, part2 := day4.Run("day4/input.txt")
	fmt.Printf("Day 4  Part 1: %v\n", part1)
	fmt.Printf("Day 4  Part 2: %v\n", part2)
}

func Day5(_ bool) {
	part1, part2 := day5.Run("day5/input.txt")
	fmt.Printf("Day 5  Part 1: %v\n", part1)
	fmt.Printf("Day 5  Part 2: %v\n", part2)
}

func Day6(_ bool) {
	part1, part2 := day6.Run("day6/input.txt")
	fmt.Printf("Day 6  Part 1: %v\n", part1)
	fmt.Printf("Day 6  Part 2: %v\n", part2)
}

func Day7(_ bool) {
	part1, part2 := day7.Run("day7/input.txt")
	fmt.Printf("Day 7  Part 1: %v\n", part1)
	fmt.Printf("Day 7  Part 2: %v\n", part2)
}

func Day8(_ bool) {
	part1, part2 := day8.Run("day8/input.txt")
	fmt.Printf("Day 8  Part 1: %v\n", part1)
	fmt.Printf("Day 8  Part 2: %v\n", part2)
}

func Day9(_ bool) {
	part1, part2 := day9.Run("day9/input.txt")
	fmt.Printf("Day 9  Part 1: %v\n", part1)
	fmt.Printf("Day 9  Part 2: %v\n", part2)
}

func Day10(_ bool) {
	part1, part2 := day10.Run("day10/input.txt")
	fmt.Printf("Day 10 Part 1: %v\n", part1)
	fmt.Printf("Day 10 Part 2\n%v\n", day10.Stringify(part2))
}

func Day11(_ bool) {
	part1, part2 := day11.Run("day11/input.txt")
	fmt.Printf("Day 11 Part 1: %v\n", part1)
	fmt.Printf("Day 11 Part 2: %v\n", part2)
}

func Day12(_ bool) {
	part1, part2 := day12.Run("day12/input.txt")
	fmt.Printf("Day 12 Part 1: %v\n", part1)
	fmt.Printf("Day 12 Part 2: %v\n", part2)
}

func Day13(_ bool) {
	part1, part2 := day13.Run("day13/input.txt")
	fmt.Printf("Day 13 Part 1: %v\n", part1)
	fmt.Printf("Day 13 Part 2: %v\n", part2)
}

func Day14(_ bool) {
	part1, part2 := day14.Run("day14/input.txt")
	fmt.Printf("Day 14 Part 1: %v\n", part1)
	fmt.Printf("Day 14 Part 2: %v\n", part2)
}

func Day15(_ bool) {
	part1, part2 := day15.Run("day15/input.txt")
	fmt.Printf("Day 15 Part 1: %v\n", part1)
	fmt.Printf("Day 15 Part 2: %v\n", part2)
}

func Day16(all bool) {
	part1, part2 := day16.Run("day16/input.txt", !all)
	fmt.Printf("Day 16 Part 1: %v\n", part1)
	if all {
		fmt.Printf("Day 16 Part 2: %v\n", part2)
	} else {
		fmt.Print("Day 16 Part 2: (skipped due to excessive processing time)\n")
	}
}

func Day17(_ bool) {
	part1, part2 := day17.Run("day17/input.txt")
	fmt.Printf("Day 17 Part 1: %v\n", part1)
	fmt.Printf("Day 17 Part 2: %v\n", part2)
}

func Day18(_ bool) {
	part1, part2 := day18.Run("day18/input.txt")
	fmt.Printf("Day 18 Part 1: %v\n", part1)
	fmt.Printf("Day 18 Part 2: %v\n", part2)
}

func RunTimed(all bool, days []func(bool)) {
	for _, function := range days {
		started := time.Now()
		function(all)
		elapsed := time.Since(started)
		fmt.Printf("Completed in %v\n", elapsed)
	}
}

func GetDays(s string) []int {
	var days []int
	if s == "" {
		return days
	}
	for _, v := range strings.Split(s, ",") {
		var a int
		var b int
		count, _ := fmt.Sscanf(v, "%d-%d", &a, &b)
		if count == 1 {
			days = append(days, a)
		} else if count == 2 {
			for i := a; i <= b; i++ {
				days = append(days, i)
			}
		}
	}
	return days
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func DaysToRun(days []int) []func(bool) {
	allDays := []func(bool){
		Day1, Day2, Day3, Day4, Day5, Day6, Day7, Day8, Day9, Day10, Day11,
		Day12, Day13, Day14, Day15, Day16, Day17, Day18,
	}
	if len(days) == 0 {
		return allDays
	}
	var selectedDays []func(bool)
	var number int
	for _, day := range allDays {
		_, _ = fmt.Sscanf(GetFunctionName(day), "main.Day%d", &number)
		if aoc_library.Contains(days, number) {
			selectedDays = append(selectedDays, day)
		}
	}
	return selectedDays
}

func main() {
	all := flag.Bool("include-all", false, "If set, does not skip CPU-intensive parts of solutions.")
	days := flag.String("days", "", "Comma-separated list of specific days to run.")
	flag.Parse()

	started := time.Now()
	daysToRun := DaysToRun(GetDays(*days))
	RunTimed(*all, daysToRun)
	elapsed := time.Since(started)
	fmt.Printf("Completed %v days in %v\n", len(daysToRun), elapsed)
}

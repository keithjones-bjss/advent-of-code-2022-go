package day11

import (
	"advent-of-code-2022/aoc_library"
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	items         []int
	items2        []int
	operation     string
	divisibleTest int
	trueTarget    int
	falseTarget   int
	inspections   int
	inspections2  int
}

func Run(filename string) (int, int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Can't open %v: %v", filename, err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var monkeys []monkey
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			monkeys = UpdateMonkeys(monkeys, line)
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Get lowest common multiple for all monkeys
	divisors := aoc_library.ArrayTranslate(monkeys, func(_ int, this monkey) int {
		return this.divisibleTest
	})
	lcm := 1
	for _, v := range divisors {
		lcm *= v
	}

	// Perform first 20 rounds for part 1
	for round := 1; round <= 20; round++ {
		monkeys = PerformPart1(monkeys, lcm)
		monkeys = PerformPart2(monkeys, lcm)
	}

	// Part 1
	inspections := aoc_library.ArrayTranslate(monkeys, func(_ int, this monkey) int {
		return this.inspections
	})
	sort.Ints(inspections)
	part1 := inspections[len(inspections)-2] * inspections[len(inspections)-1]

	// Perform remaining rounds for part 2
	for round := 21; round <= 10000; round++ {
		monkeys = PerformPart2(monkeys, lcm)
	}

	// Part 2
	inspections2 := aoc_library.ArrayTranslate(monkeys, func(_ int, this monkey) int {
		return this.inspections2
	})
	sort.Ints(inspections2)
	part2 := inspections2[len(inspections2)-2] * inspections2[len(inspections2)-1]

	return part1, part2
}

func PerformPart1(monkeys []monkey, lcm int) []monkey {
	for index, thisMonkey := range monkeys {
		for _, item := range thisMonkey.items {
			worryLevel := Evaluate(thisMonkey.operation, item) / 3
			var target int
			if worryLevel%thisMonkey.divisibleTest == 0 {
				target = thisMonkey.trueTarget
			} else {
				target = thisMonkey.falseTarget
			}
			monkeys[target].items = append(monkeys[target].items, worryLevel)
		}
		monkeys[index].inspections += len(monkeys[index].items)
		monkeys[index].items = []int{}
	}
	return monkeys
}

func PerformPart2(monkeys []monkey, lcm int) []monkey {
	for index, thisMonkey := range monkeys {
		for _, item := range thisMonkey.items2 {
			worryLevel := Evaluate(thisMonkey.operation, item) % lcm
			var target int
			if worryLevel%thisMonkey.divisibleTest == 0 {
				target = thisMonkey.trueTarget
			} else {
				target = thisMonkey.falseTarget
			}
			monkeys[target].items2 = append(monkeys[target].items2, worryLevel)
		}
		monkeys[index].inspections2 += len(monkeys[index].items2)
		monkeys[index].items2 = []int{}
	}
	return monkeys
}

func UpdateMonkeys(monkeys []monkey, line string) []monkey {
	last := len(monkeys) - 1
	// New monkey
	if line[:6] == "Monkey" {
		return append(monkeys, monkey{})
	}
	// Starting items
	if line[2:16] == "Starting items" {
		items := aoc_library.ArrayTranslate(strings.Split(line[18:], ", "),
			func(_ int, s string) int {
				v, _ := strconv.Atoi(s)
				return v
			})
		return append(monkeys[:last], monkey{items: items, items2: items})
	}
	// Operation
	if line[2:11] == "Operation" {
		return append(monkeys[:last], monkey{
			items:     monkeys[last].items,
			items2:    monkeys[last].items2,
			operation: line[19:]})
	}
	// Divisible test
	if line[2:20] == "Test: divisible by" {
		divisor, _ := strconv.Atoi(line[21:])
		return append(monkeys[:last], monkey{
			items:         monkeys[last].items,
			items2:        monkeys[last].items2,
			operation:     monkeys[last].operation,
			divisibleTest: divisor,
		})
	}
	// True target
	if line[4:11] == "If true" {
		target, _ := strconv.Atoi(line[29:])
		return append(monkeys[:last], monkey{
			items:         monkeys[last].items,
			items2:        monkeys[last].items2,
			operation:     monkeys[last].operation,
			divisibleTest: monkeys[last].divisibleTest,
			trueTarget:    target,
		})
	}
	// False target
	if line[4:12] == "If false" {
		target, _ := strconv.Atoi(line[30:])
		return append(monkeys[:last], monkey{
			items:         monkeys[last].items,
			items2:        monkeys[last].items2,
			operation:     monkeys[last].operation,
			divisibleTest: monkeys[last].divisibleTest,
			trueTarget:    monkeys[last].trueTarget,
			falseTarget:   target,
		})
	}
	// Unrecognised command
	return monkeys
}

func Evaluate(expression string, old int) int {
	parts := strings.Split(expression, " ")
	lhs := EvaluateOperand(parts[0], old)
	rhs := EvaluateOperand(parts[2], old)
	if parts[1] == "+" {
		return lhs + rhs
	}
	if parts[1] == "-" {
		return lhs - rhs
	}
	if parts[1] == "*" {
		return lhs * rhs
	}
	if parts[1] == "/" {
		return lhs / rhs
	}
	return 0 // Unrecognised expression
}

func EvaluateOperand(operand string, old int) int {
	if operand == "old" {
		return old
	}
	v, _ := strconv.Atoi(operand)
	return v
}

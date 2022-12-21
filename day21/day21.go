package day21

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Expression struct {
	lhs     string
	rhs     string
	operand rune
	value   int
}

const (
	Plus   rune = '+'
	Minus  rune = '-'
	Times  rune = '*'
	Divide rune = '/'
	Number rune = '#'
)

func Run(filename string) (int, int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Can't open %v: %v", filename, err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	expressions := make(map[string]Expression)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			k, v := Parse(line)
			expressions[k] = v
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1, ok := Evaluate(expressions, "root")
	if !ok {
		log.Fatalf("Unable to evaluate expression")
	}

	// TODO
	part2 := 0

	return part1, part2
}

func Parse(line string) (string, Expression) {
	var expression Expression
	keyValuePair := strings.Split(line, ": ")
	if strings.Contains(keyValuePair[1], " ") {
		operation := strings.Split(keyValuePair[1], " ")
		expression.lhs = operation[0]
		expression.operand = []rune(operation[1])[0]
		expression.rhs = operation[2]
	} else {
		expression.operand = Number
		expression.value, _ = strconv.Atoi(keyValuePair[1])
	}
	return keyValuePair[0], expression
}

func Evaluate(expressions map[string]Expression, name string) (int, bool) {
	expression, ok := expressions[name]
	if ok {
		if expression.operand == Number {
			//log.Printf("%v = %v", name, expression.value)
			return expression.value, true
		}
		lhs, ok := Evaluate(expressions, expression.lhs)
		if !ok {
			return 0, ok
		}
		rhs, ok := Evaluate(expressions, expression.rhs)
		if !ok {
			return 0, ok
		}
		//log.Printf("%v = %v(%v) %v %v(%v)", name, expression.lhs, lhs, string(expression.operand), expression.rhs, rhs)
		result := 0
		switch expression.operand {
		case Plus:
			result = lhs + rhs
		case Minus:
			result = lhs - rhs
		case Times:
			result = lhs * rhs
		case Divide:
			result = lhs / rhs
		default:
			log.Fatalf("Invalid operand '%v' in expression %v: %v", string(expression.operand), name, expression)
			return 0, false
		}
		return result, true
	}
	log.Fatalf("Invalid expression '%v'", name)
	return 0, false
}

package day21

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Expression struct {
	lhs      string
	rhs      string
	operand  rune
	value    int
	human    bool
	lhsHuman bool
	rhsHuman bool
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

	part1, _ := Evaluate(expressions, "root")

	part2 := Part2(expressions)

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
			if name == "humn" {
				expression.human = true
			}
			//log.Printf("%v = %v", name, expression.value)
			return expression.value, expression.human
		}
		lhs, human := Evaluate(expressions, expression.lhs)
		expression.lhsHuman = human
		rhs, human := Evaluate(expressions, expression.rhs)
		expression.rhsHuman = human
		expressions[name] = expression // Do we need this?
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
		return result, expression.lhsHuman || expression.rhsHuman
	}
	log.Fatalf("Invalid expression '%v'", name)
	return 0, false
}

func Part2(expressions map[string]Expression) int {
	target := 0
	expressionName := ""
	root := expressions["root"]
	if expressions["root"].lhsHuman {
		target, _ = Evaluate(expressions, expressions["root"].rhs)
		expressionName = root.lhs
	} else {
		target, _ = Evaluate(expressions, expressions["root"].lhs)
		expressionName = root.rhs
	}
	log.Printf("Target value is %v", target)
	result := target
	for expressionName != "humn" {
		expression, _ := expressions[expressionName]
		//log.Printf("[%v] %v: %v[%v] %v %v[%v]",
		//	result, expressionName, expression.lhs, expression.lhsHuman,
		//	string(expression.operand), expression.rhs, expression.rhsHuman)
		if expression.lhsHuman {
			rhsValue, _ := Evaluate(expressions, expression.rhs)
			//log.Printf("[%v] %v: %v %v %v",
			//	result, expressionName, expression.lhs, string(expression.operand), rhsValue)
			switch expression.operand {
			case Plus:
				result = result - rhsValue
			case Minus:
				result = result + rhsValue
			case Times:
				result = result / rhsValue
			case Divide:
				result = result * rhsValue
			}
			expressionName = expression.lhs
		} else {
			lhsValue, _ := Evaluate(expressions, expression.lhs)
			//log.Printf("[%v] %v: %v %v %v",
			//	result, expressionName, lhsValue, string(expression.operand), expression.rhs)
			switch expression.operand {
			case Plus:
				result = result - lhsValue
			case Minus:
				result = lhsValue - result
			case Times:
				result = result / lhsValue
			case Divide:
				result = lhsValue / result
			}
			expressionName = expression.rhs
		}
	}
	part2 := result
	return part2
}

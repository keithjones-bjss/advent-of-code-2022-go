package day11

import (
	"advent-of-code-2022/testlib"
	"testing"
)

func TestEvaluateOperandIfOld(t *testing.T) {
	result := EvaluateOperand("old", 4)
	testlib.AssertEqual(t, result, 4)
}

func TestEvaluateOperandIfInt(t *testing.T) {
	result := EvaluateOperand("7", 4)
	testlib.AssertEqual(t, result, 7)
}

func TestEvaluateAddition(t *testing.T) {
	result := Evaluate("old + 7", 4)
	testlib.AssertEqual(t, result, 11)
}

func TestEvaluateSubtraction(t *testing.T) {
	result := Evaluate("old - 7", 4)
	testlib.AssertEqual(t, result, -3)
}

func TestEvaluateMultiplication(t *testing.T) {
	result := Evaluate("old * 7", 4)
	testlib.AssertEqual(t, result, 28)
}

func TestEvaluateDivision(t *testing.T) {
	result := Evaluate("old / 7", 28)
	testlib.AssertEqual(t, result, 4)
}

func TestEvaluateSquare(t *testing.T) {
	result := Evaluate("old * old", 4)
	testlib.AssertEqual(t, result, 16)
}

func TestPart1(t *testing.T) {
	result, _ := Run("test.txt")
	testlib.AssertEqual(t, result, 10605)
}

func TestPart2(t *testing.T) {
	_, result := Run("test.txt")
	testlib.AssertEqual(t, result, 2713310158)
}

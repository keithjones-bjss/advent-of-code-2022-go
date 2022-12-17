package day17

import (
	"advent-of-code-2022/testlib"
	"testing"
)

//func TestTenSteps(t *testing.T) {
//	result := SolveTower(">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>", 100, true)
//	testlib.AssertEqual(t, result, 3068)
//}

func TestPart1(t *testing.T) {
	result, _ := Run("test.txt")
	testlib.AssertEqual(t, result, 3068)
}

func TestPart2(t *testing.T) {
	_, result := Run("test.txt")
	testlib.AssertEqual(t, result, 1514285714288)
}

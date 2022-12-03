package day1

import (
	"advent-of-code-2022/testlib"
	"testing"
)

func TestPart1(t *testing.T) {
	top := Top("test.txt")
	testlib.AssertEqual(t, top[0], 10000)
	testlib.AssertEqual(t, top[1], 11000)
	testlib.AssertEqual(t, top[2], 24000)
}

func TestPart2(t *testing.T) {
	top := Top("test.txt")
	sum := Sum(top)
	testlib.AssertEqual(t, sum, 45000)
}

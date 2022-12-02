package day1

import (
	"advent-of-code-2022/testlib"
	"testing"
)

func TestPart1(t *testing.T) {
	top := Top("../files/day1_test.txt")
	testlib.AssertEqualInt64(t, top[0], 10000)
	testlib.AssertEqualInt64(t, top[1], 11000)
	testlib.AssertEqualInt64(t, top[2], 24000)
}

func TestPart2(t *testing.T) {
	top := Top("../files/day1_test.txt")
	sum := Sum(top)
	testlib.AssertEqualInt64(t, sum, 45000)
}

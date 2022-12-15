package aoc_library

import (
	"advent-of-code-2022/testlib"
	"testing"
)

func TestMoveUp(t *testing.T) {
	original := Point{X: 4, Y: 10}
	result := original.Move("U")
	testlib.AssertEqual(t, result.X, original.X)
	testlib.AssertEqual(t, result.Y, original.Y-1)
}

func TestMoveDown(t *testing.T) {
	original := Point{X: 4, Y: 10}
	result := original.Move("D")
	testlib.AssertEqual(t, result.X, original.X)
	testlib.AssertEqual(t, result.Y, original.Y+1)
}

func TestMoveLeft(t *testing.T) {
	original := Point{X: 4, Y: 10}
	result := original.Move("L")
	testlib.AssertEqual(t, result.X, original.X-1)
	testlib.AssertEqual(t, result.Y, original.Y)
}

func TestMoveRight(t *testing.T) {
	original := Point{X: 4, Y: 10}
	result := original.Move("R")
	testlib.AssertEqual(t, result.X, original.X+1)
	testlib.AssertEqual(t, result.Y, original.Y)
}

func TestDistanceFromGivenSamePosition(t *testing.T) {
	a := Point{X: 4, Y: 10}
	b := Point{X: 4, Y: 10}
	result := a.DistanceFrom(b)
	testlib.AssertEqual(t, result, 0)
}

func TestDistanceFromGivenDifferentPositions(t *testing.T) {
	a := Point{X: 4, Y: 23}
	b := Point{X: 17, Y: 10}
	result := a.DistanceFrom(b)
	testlib.AssertEqual(t, result, 26)
}

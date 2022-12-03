package day2

import (
	"advent-of-code-2022/testlib"
	"testing"
)

func TestPart1Score_Rock_Rock(t *testing.T) {
	score := Part1Score("A X")
	testlib.AssertEqual(t, score, 4)
}

func TestPart1Score_Rock_Paper(t *testing.T) {
	score := Part1Score("B X")
	testlib.AssertEqual(t, score, 1)
}

func TestPart1Score_Rock_Scissors(t *testing.T) {
	score := Part1Score("C X")
	testlib.AssertEqual(t, score, 7)
}

func TestPart1Score_Paper_Rock(t *testing.T) {
	score := Part1Score("A Y")
	testlib.AssertEqual(t, score, 8)
}

func TestPart1Score_Paper_Paper(t *testing.T) {
	score := Part1Score("B Y")
	testlib.AssertEqual(t, score, 5)
}

func TestPart1Score_Paper_Scissors(t *testing.T) {
	score := Part1Score("C Y")
	testlib.AssertEqual(t, score, 2)
}

func TestPart1Score_Scissors_Rock(t *testing.T) {
	score := Part1Score("A Z")
	testlib.AssertEqual(t, score, 3)
}

func TestPart1Score_Scissors_Paper(t *testing.T) {
	score := Part1Score("B Z")
	testlib.AssertEqual(t, score, 9)
}

func TestPart1Score_Scissors_Scissors(t *testing.T) {
	score := Part1Score("C Z")
	testlib.AssertEqual(t, score, 6)
}

func TestPart2Score_Rock_Rock(t *testing.T) {
	score := Part2Score("A Y") // Rock / Draw
	testlib.AssertEqual(t, score, 4)
}

func TestPart2Score_Rock_Paper(t *testing.T) {
	score := Part2Score("B X") // Paper / Loss
	testlib.AssertEqual(t, score, 1)
}

func TestPart2Score_Rock_Scissors(t *testing.T) {
	score := Part2Score("C Z") // Scissors / Win
	testlib.AssertEqual(t, score, 7)
}

func TestPart2Score_Paper_Rock(t *testing.T) {
	score := Part2Score("A Z") // Paper / Win
	testlib.AssertEqual(t, score, 8)
}

func TestPart2Score_Paper_Paper(t *testing.T) {
	score := Part2Score("B Y") // Paper / Draw
	testlib.AssertEqual(t, score, 5)
}

func TestPart2Score_Paper_Scissors(t *testing.T) {
	score := Part2Score("C X") // Scissors / Loss
	testlib.AssertEqual(t, score, 2)
}

func TestPart2Score_Scissors_Rock(t *testing.T) {
	score := Part2Score("A X") // Rock / Loss
	testlib.AssertEqual(t, score, 3)
}

func TestPart2Score_Scissors_Paper(t *testing.T) {
	score := Part2Score("B Z") // Paper / Win
	testlib.AssertEqual(t, score, 9)
}

func TestPart2Score_Scissors_Scissors(t *testing.T) {
	score := Part2Score("C Y") // Scissors / Draw
	testlib.AssertEqual(t, score, 6)
}

func TestPart1(t *testing.T) {
	score, _ := Run("test.txt")
	testlib.AssertEqual(t, score, 15)
}

func TestPart2(t *testing.T) {
	_, score := Run("test.txt")
	testlib.AssertEqual(t, score, 12)
}

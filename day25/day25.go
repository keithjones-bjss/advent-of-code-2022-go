package day25

import (
	"bufio"
	"log"
	"os"
)

func Run(filename string) (string, string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Can't open %v: %v", filename, err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lines = append(lines, line)
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum := 0
	for _, line := range lines {
		sum += ConvertSnafuToInt(line)
	}

	part1 := ConvertIntToSnafu(sum)
	part2 := ""

	return part1, part2
}

const (
	MinusTwo = '='
	MinusOne = '-'
	Zero     = '0'
	One      = '1'
	Two      = '2'
)

var digitMap = map[rune]int{
	MinusTwo: -2,
	MinusOne: -1,
	Zero:     0,
	One:      1,
	Two:      2,
}

var snafuDigits = []rune("012=-")

func ConvertSnafuToInt(snafu string) int {
	result := 0
	runes := []rune(snafu)
	for _, digit := range runes {
		result = result*5 + digitMap[digit]
	}
	return result
}

func ConvertIntToSnafu(number int) string {
	if number == 0 {
		return "0"
	}
	result := ""
	for number != 0 {
		digit := number % 5
		number = number / 5
		if digit > 2 {
			number++
		}
		result = string(snafuDigits[digit]) + result
	}
	return result
}

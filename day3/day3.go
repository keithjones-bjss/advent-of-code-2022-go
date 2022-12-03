package day3

import (
	"bufio"
	"log"
	"os"
	"unicode/utf8"
)

func Run(filename string) (int, int) {
	part1 := 0
	part2 := 0

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Can't open %v: %v", filename, err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)
	var batch []string
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			part1 += Part1Score(line)
			batch = append(batch, line)
			if len(batch) == 3 {
				part2 += Part2Score(batch)
				batch = nil
			}
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return part1, part2
}

func Part1Score(rucksack string) int {
	size := len(rucksack)
	half := size / 2
	compartments := []string{rucksack[0:half], rucksack[half:size]}
	result := Intersection(compartments)
	return Priority(result)
}

func Part2Score(rucksacks []string) int {
	result := Intersection(rucksacks)
	return Priority(result)
}

type MapFilterFunction[K comparable, V any] func(key K, value V) bool

func MapFilter[K comparable, V any](m map[K]V, function MapFilterFunction[K, V]) map[K]V {
	n := make(map[K]V)
	for key, value := range m {
		if function(key, value) {
			n[key] = value
		}
	}
	return n
}

func Intersection(strings []string) string {
	commonChars := make(map[rune]int)
	for _, value := range strings {
		distinctChars := DistinctChars(value)
		for _, char := range distinctChars {
			commonChars[char]++
		}
	}
	filteredChars := MapFilter(commonChars, func(_ rune, value int) bool { return value == len(strings) })
	return MapKeysToString(filteredChars)
}

func MapKeysToString(m map[rune]int) string {
	var result []byte
	for char := range m {
		result = utf8.AppendRune(result, char)
	}
	return string(result)
}

func DistinctChars(str string) string {
	var result []byte
	distinctChars := make(map[rune]bool)
	for _, char := range str {
		_, present := distinctChars[char]
		if !present {
			distinctChars[char] = true
			result = utf8.AppendRune(result, char)
		}
	}
	return string(result)
}

func Priority(str string) int {
	char := int(str[0])
	if char >= 97 && char <= 122 {
		return char - 96 // a-z => 1-26
	}
	if char >= 65 && char <= 90 {
		return char - 38 // A-Z => 27-52
	}
	return 0
}

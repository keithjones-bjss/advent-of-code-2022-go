package day20

import (
	"advent-of-code-2022/aoc_library"
	"bufio"
	"log"
	"os"
	"strconv"
)

func Run(filename string) (int, int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Can't open %v: %v", filename, err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var numbers []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			value, _ := strconv.Atoi(line)
			if aoc_library.Contains(numbers, value) {
				log.Fatalf("Duplicate value %v!", value)
			}
			numbers = append(numbers, value)
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	mixed := Mix(numbers)
	zero := IndexOf(mixed, 0)
	part1 := ValueAt(mixed, zero+1000) + ValueAt(mixed, zero+2000) + ValueAt(mixed, zero+3000)

	part2 := 0

	return part1, part2
}

func Mix(numbers []int) []int {
	size := len(numbers)
	nextNumbers := append([]int{}, numbers...)
	done := make(map[int]bool)
	done[0] = true
	index := 0
	for index < size {
		value := nextNumbers[index]
		if done[value] {
			index++
			continue
		}
		done[value] = true
		newPosition := index + nextNumbers[index]
		if newPosition <= 0 && value < 0 {
			newPosition--
		}
		if newPosition >= size && value > 0 {
			newPosition++
		}
		newPosition = WrapNumber(newPosition, size)
		//log.Printf("Moving %d from position %d to position %d.", value, index, newPosition)
		var slice []int
		if newPosition < index {
			//log.Printf("Moving backwards %d places", index-newPosition)
			slice = append(slice, nextNumbers[:newPosition]...)
			//log.Printf("Slice left of new position: %v", slice)
			slice = append(slice, nextNumbers[index])
			slice = append(slice, nextNumbers[newPosition:index]...)
			//log.Printf("Slice up to index: %v", slice)
			if index < size {
				slice = append(slice, nextNumbers[index+1:]...)
			}
		} else {
			//log.Printf("Moving forwards %d places", newPosition-index)
			if index > 0 {
				slice = append(slice, nextNumbers[:index]...)
				//log.Printf("Slice left of value: %v", slice)
			}
			slice = append(slice, nextNumbers[index+1:newPosition+1]...)
			//log.Printf("Slice up to new position: %v", slice)
			slice = append(slice, nextNumbers[index])
			if newPosition+1 < size {
				slice = append(slice, nextNumbers[newPosition+1:]...)
			}
		}
		//log.Printf("%v -> %v", nextNumbers, slice)
		nextNumbers = slice
	}
	return nextNumbers
}

func WrapNumber(i int, size int) int {
	if i >= 0 {
		//log.Printf("WrapNumber(%d, %d) -> %d", i, size, i%size)
		return i % size
	}
	//log.Printf("WrapNumber(%d, %d) -> %d", i, size, size + i%size)
	return size + i%size
}

func IndexOf[T comparable](array []T, value T) int {
	for i, v := range array {
		if v == value {
			return i
		}
	}
	return -1
}

func ValueAt[T comparable](array []T, index int) T {
	return array[index%len(array)]
}

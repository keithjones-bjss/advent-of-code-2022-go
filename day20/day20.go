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

	decryptionKey := 811589153
	var numbers1 []int
	var numbers2 []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			value, _ := strconv.Atoi(line)
			numbers1 = append(numbers1, value)
			numbers2 = append(numbers2, value*decryptionKey)
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	mixed1 := Mix(numbers1, 1)
	zero1 := IndexOf(mixed1, 0)
	part1 := ValueAt(mixed1, zero1+1000) + ValueAt(mixed1, zero1+2000) + ValueAt(mixed1, zero1+3000)

	mixed2 := Mix(numbers2, 10)
	zero2 := IndexOf(mixed2, 0)
	part2 := ValueAt(mixed2, zero2+1000) + ValueAt(mixed2, zero2+2000) + ValueAt(mixed2, zero2+3000)

	return part1, part2
}

type Mixable struct {
	value           int
	initialPosition int
	done            bool
}

func Mix(numbers []int, times int) []int {
	size := len(numbers)
	nextNumbers := append([]Mixable{}, aoc_library.ArrayTranslate(numbers, func(i int, v int) Mixable {
		return Mixable{v, i, v == 0}
	})...)
	for count := 0; count < size; count++ {
		index := Find(nextNumbers, func(_ int, value Mixable) bool { return value.initialPosition == count })
		value := nextNumbers[index].value
		newPosition := index + value
		if newPosition <= 0 && value < 0 {
			//log.Printf("%d %d >>>", value, newPosition)
			newPosition += (size - 1) * (-newPosition / size)
			for newPosition <= 0 && value < 0 {
				newPosition += size - 1
			}
		}
		if newPosition >= size && value > 0 {
			//log.Printf("%d %d <<<", value, newPosition)
			newPosition += (1 - size) * (newPosition / size)
			for newPosition >= size && value > 0 {
				newPosition += 1 - size
			}
		}
		//log.Printf("Moving %d from position %d to position %d.", value, index, newPosition)
		var slice []Mixable
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
	return aoc_library.ArrayTranslate(nextNumbers, func(_ int, v Mixable) int {
		return v.value
	})
}

type FilterFunction[T comparable, U any] func(index T, value U) bool

func IndexOf[T comparable](array []T, value T) int {
	return Find(array, func(_ int, v T) bool { return v == value })
}

func Find[T comparable](array []T, function FilterFunction[int, T]) int {
	for index, value := range array {
		if function(index, value) {
			return index
		}
	}
	return -1
}

func ValueAt[T comparable](array []T, index int) T {
	return array[index%len(array)]
}

package aoc_library

import "unicode/utf8"

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

func MapKeysToString[T any](m map[rune]T) string {
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

func StringIntersection(strings []string) string {
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

func MapKeysToArray[T comparable, U any](m map[T]U) []T {
	var result []T
	for value := range m {
		result = append(result, value)
	}
	return result
}

func DistinctValues[T comparable](array []T) []T {
	var result []T
	distinctValues := make(map[T]bool)
	for _, value := range array {
		_, present := distinctValues[value]
		if !present {
			distinctValues[value] = true
			result = append(result, value)
		}
	}
	return result
}

func Intersection[T comparable](arrays [][]T) []T {
	commonValues := make(map[T]int)
	for _, value := range arrays {
		distinctValues := DistinctValues(value)
		for _, value := range distinctValues {
			commonValues[value]++
		}
	}
	filteredValues := MapFilter(commonValues, func(_ T, value int) bool { return value == len(arrays) })
	return MapKeysToArray(filteredValues)
}

func Contains[T comparable](array []T, element T) bool {
	for _, value := range array {
		if value == element {
			return true
		}
	}
	return false
}

func ArrayEquals[T comparable](array1 []T, array2 []T) bool {
	if len(array1) != len(array2) {
		return false
	}
	for index, _ := range array1 {
		if array1[index] != array2[index] {
			return false
		}
	}
	return true
}

func ArrayContains[T comparable](arrays [][]T, array []T) bool {
	for _, value := range arrays {
		if ArrayEquals(value, array) {
			return true
		}
	}
	return false
}

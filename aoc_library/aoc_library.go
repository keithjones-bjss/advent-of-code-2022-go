package aoc_library

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

func CommonValues[T comparable](arrays [][]T) map[T]int {
	commonValues := make(map[T]int)
	for _, value := range arrays {
		distinctValues := DistinctValues(value)
		for _, value := range distinctValues {
			commonValues[value]++
		}
	}
	return commonValues
}

func AllIntersection[T comparable](arrays [][]T) []T {
	commonValues := CommonValues(arrays)
	filteredValues := MapFilter(commonValues, func(_ T, value int) bool { return value == len(arrays) })
	return MapKeysToArray(filteredValues)
}

func AnyIntersection[T comparable](arrays [][]T) []T {
	commonValues := CommonValues(arrays)
	filteredValues := MapFilter(commonValues, func(_ T, value int) bool { return value > 1 })
	return MapKeysToArray(filteredValues)
}

func AllStringIntersection(strings []string) string {
	runeArrays := ArrayTranslate(strings, func(_ int, value string) []rune { return []rune(value) })
	result := AllIntersection(runeArrays)
	return string(result)
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
	for index := range array1 {
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

type TranslateFunction[T any, U any] func(index int, initialValue T) U

func ArrayTranslate[T any, U any](array []T, function TranslateFunction[T, U]) []U {
	var newArray []U
	for index, value := range array {
		newArray = append(newArray, function(index, value))
	}
	return newArray
}

type FilterFunction[T comparable, U any] func(index T, value U) bool

func IndexOf[T comparable](array []T, value T) int {
	return Find(array, func(_ int, v T) bool { return v == value })
}

func LastIndexOf[T comparable](array []T, value T) int {
	return FindReverse(array, func(_ int, v T) bool { return v == value })
}

func Find[T comparable](array []T, function FilterFunction[int, T]) int {
	for index, value := range array {
		if function(index, value) {
			return index
		}
	}
	return -1
}

func FindReverse[T comparable](array []T, function FilterFunction[int, T]) int {
	for index := len(array) - 1; index >= 0; index-- {
		value := array[index]
		if function(index, value) {
			return index
		}
	}
	return -1
}

func ValueAt[T comparable](array []T, index int) T {
	return array[index%len(array)]
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func Sign(value int) int {
	if value < 0 {
		return -1
	}
	if value > 0 {
		return 1
	}
	return 0
}

func Mod(dividend int, divisor int) int {
	if dividend < 0 {
		return (divisor + (dividend % divisor)) % divisor
	}
	return dividend % divisor
}

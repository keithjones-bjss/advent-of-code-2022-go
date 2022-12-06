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

func AnyMatch[T comparable](arrays [][]T) []T {
	commonValues := make(map[T]int)
	for _, value := range arrays {
		distinctValues := DistinctValues(value)
		for _, value := range distinctValues {
			commonValues[value]++
		}
	}
	filteredValues := MapFilter(commonValues, func(_ T, value int) bool { return value > 1 })
	return MapKeysToArray(filteredValues)
}

func StringIntersection(strings []string) string {
	var runeArrays [][]rune
	for _, value := range strings {
		runeArrays = append(runeArrays, []rune(value))
	}
	result := Intersection(runeArrays)
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

type TranslateFunction[T any, U any] func(index int, initialValue T) U

func ArrayTranslate[T any, U any](array []T, function TranslateFunction[T, U]) []U {
	var newArray []U
	for index, value := range array {
		newArray = append(newArray, function(index, value))
	}
	return newArray
}

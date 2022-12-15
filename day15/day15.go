package day15

import (
	"advent-of-code-2022/aoc_library"
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Sensor struct {
	position    aoc_library.Point
	beacon      aoc_library.Point
	signalRange int
}

type Rectangle struct {
	minX int
	minY int
	maxX int
	maxY int
}

func Run(filename string) (int, int) {
	return RunAtRow(filename, 2000000)
}

func RunAtRow(filename string, row int) (int, int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Can't open %v: %v", filename, err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var sensors []Sensor
	var thisSensor Sensor
	var bounds *Rectangle

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			thisSensor, bounds = Parse(line, bounds)
			sensors = append(sensors, thisSensor)
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1 := CountBeaconlessPositions(sensors, *bounds, row)
	part2 := FindMissingBeacon(sensors, row)

	return part1, part2
}

func Parse(line string, oldBounds *Rectangle) (Sensor, *Rectangle) {
	var result Sensor
	var bounds Rectangle
	_, _ = fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
		&result.position.X, &result.position.Y, &result.beacon.X, &result.beacon.Y)
	result.signalRange = result.position.DistanceFrom(result.beacon)
	bounds.maxX = result.position.X + result.signalRange
	bounds.maxY = result.position.Y + result.signalRange
	bounds.minX = result.position.X - result.signalRange
	bounds.minY = result.position.Y - result.signalRange
	if oldBounds != nil {
		bounds.maxX = aoc_library.Max(bounds.maxX, oldBounds.maxX)
		bounds.maxY = aoc_library.Max(bounds.maxY, oldBounds.maxY)
		bounds.minX = aoc_library.Min(bounds.minX, oldBounds.minX)
		bounds.minY = aoc_library.Min(bounds.minY, oldBounds.minY)
	}
	return result, &bounds
}

func CountBeaconlessPositions(sensors []Sensor, bounds Rectangle, y int) int {
	count := 0
	row := []rune(strings.Repeat(".", bounds.maxX-bounds.minX+1))
	for _, v := range sensors {
		signalRangeAtRow := v.signalRange - aoc_library.Abs(y-v.position.Y)
		if signalRangeAtRow >= 0 {
			rowX := v.position.X - bounds.minX
			if v.position.Y == y {
				row[v.position.X-bounds.minX] = 'S'
				count++
			}
			if v.beacon.Y == y {
				row[v.beacon.X-bounds.minX] = 'B'
			}
			for x := rowX - signalRangeAtRow; x <= rowX+signalRangeAtRow; x++ {
				if row[x] == '.' {
					row[x] = '#'
					count++
				}
			}
		}
	}
	return count
}

type minmax struct {
	min int
	max int
}

func FindMissingBeacon(sensors []Sensor, rows int) int {
	lastY := rows*2 + 1
	for y := 0; y < lastY; y++ {
		var ranges []minmax
		for _, v := range sensors {
			signalRangeAtRow := v.signalRange - aoc_library.Abs(y-v.position.Y)
			if signalRangeAtRow >= 0 {
				minX := aoc_library.Max(0, v.position.X-signalRangeAtRow)
				maxX := aoc_library.Min(lastY-1, v.position.X+signalRangeAtRow)
				ranges = append(ranges, minmax{minX, maxX})
			}
		}
		sort.Slice(ranges, func(i int, j int) bool { return ranges[i].min < ranges[j].min })
		last := 0
		for _, v := range ranges {
			if v.min > last {
				return 4000000*last + y
			}
			if v.max >= last {
				last = v.max + 1
			}
		}
	}
	return -1
}

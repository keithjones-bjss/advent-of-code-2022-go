package day19

import (
	"advent-of-code-2022/aoc_library"
	"bufio"
	"fmt"
	"log"
	"os"
)

func Run(filename string, skipPart2 bool) (int, int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Can't open %v: %v", filename, err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var blueprints []Blueprint

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			blueprints = append(blueprints, Parse(line))
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1 := 0
	for _, v := range blueprints {
		geodesCracked := Walk(24, v, [][][]int{{{1, 0}, {0, 0}, {0, 0}, {0, 0}}})
		qualityNumber := geodesCracked * v.id
		part1 += qualityNumber
	}

	part2 := 1
	if !skipPart2 {
		for count := 0; count < 3; count++ {
			part2 = part2 * Walk(32, blueprints[count], [][][]int{{{1, 0}, {0, 0}, {0, 0}, {0, 0}}})
		}
	}

	return part1, part2
}

type Blueprint struct {
	id            int
	ore           int
	clay          int
	obsidianOre   int
	obsidianClay  int
	geodeOre      int
	geodeObsidian int
	geodesCracked int
	qualityNumber int
}

const (
	Ore      int = 0
	Clay     int = 1
	Obsidian int = 2
	Geode    int = 3
)

const (
	Collectors int = 0
	Available  int = 1
)

func Parse(line string) Blueprint {
	var blueprint Blueprint
	_, _ = fmt.Sscanf(line, "Blueprint %d: "+
		"Each ore robot costs %d ore. "+
		"Each clay robot costs %d ore. "+
		"Each obsidian robot costs %d ore and %d clay. "+
		"Each geode robot costs %d ore and %d obsidian.",
		&blueprint.id,
		&blueprint.ore, &blueprint.clay,
		&blueprint.obsidianOre, &blueprint.obsidianClay,
		&blueprint.geodeOre, &blueprint.geodeObsidian)
	return blueprint
}

func Walk(time int, blueprint Blueprint, paths [][][]int) int {
	var possiblePaths [][][]int
	maxOreCollectors := aoc_library.Max(aoc_library.Max(blueprint.ore, blueprint.clay), aoc_library.Max(blueprint.geodeOre, blueprint.obsidianOre))
	for _, initialResources := range paths {
		//log.Printf("time %d blueprint %v resources %v", 24-time, blueprint, initialResources)
		// Check if we can produce more robots; don't produce more than we need
		canMakeGeodeRobot := initialResources[Ore][Available] >= blueprint.geodeOre &&
			initialResources[Obsidian][Available] >= blueprint.geodeObsidian
		canMakeObsidianRobot := initialResources[Ore][Available] >= blueprint.obsidianOre &&
			initialResources[Clay][Available] >= blueprint.obsidianClay &&
			initialResources[Obsidian][Collectors] < blueprint.geodeObsidian
		canMakeClayRobot := initialResources[Ore][Available] >= blueprint.clay &&
			initialResources[Clay][Collectors] < blueprint.obsidianClay
		canMakeOreRobot := initialResources[Ore][Available] >= blueprint.ore &&
			initialResources[Ore][Collectors] < maxOreCollectors
		resources := CopyAndUpdateResources(initialResources)
		// Should we just collect resources?
		canCollectObsidian := !canMakeGeodeRobot && initialResources[Obsidian][Collectors] > 0
		canCollectClay := !canMakeObsidianRobot && initialResources[Clay][Collectors] > 0
		canCollectOre := !canMakeOreRobot || !canMakeClayRobot
		if canCollectOre || canCollectClay || canCollectObsidian {
			possiblePaths = append(possiblePaths, resources)
		}
		// Produce geode robot if we can
		if canMakeGeodeRobot {
			copyResources := CopyResources(resources)
			copyResources[Ore][Available] -= blueprint.geodeOre
			copyResources[Obsidian][Available] -= blueprint.geodeObsidian
			copyResources[Geode][Collectors]++
			possiblePaths = append(possiblePaths, copyResources)
		}
		// Produce obsidian robot if we can
		if canMakeObsidianRobot {
			copyResources := CopyResources(resources)
			copyResources[Ore][Available] -= blueprint.obsidianOre
			copyResources[Clay][Available] -= blueprint.obsidianClay
			copyResources[Obsidian][Collectors]++
			possiblePaths = append(possiblePaths, copyResources)
		}
		// Produce a clay robot if we can
		if canMakeClayRobot {
			copyResources := CopyResources(resources)
			copyResources[Ore][Available] -= blueprint.clay
			copyResources[Clay][Collectors]++
			possiblePaths = append(possiblePaths, copyResources)
		}
		// Produce an ore robot if we can
		if canMakeOreRobot {
			copyResources := CopyResources(resources)
			copyResources[Ore][Available] -= blueprint.ore
			copyResources[Ore][Collectors]++
			possiblePaths = append(possiblePaths, copyResources)
		}
	}
	geodesCracked := 0
	if time > 1 {
		var goodPaths [][][]int
		var scores []int
		threshold := 0 // filter out paths that don't meet a minimum geode cracking threshold
		for _, v := range possiblePaths {
			calc := v[Geode][Available]
			threshold = aoc_library.Max(threshold, calc + v[Geode][Collectors] * (time - 1))
			for count := 1; time - count > 0; count++ {
				calc += v[Geode][Collectors] + count
			}
			scores = append(scores, calc)
		}
		for i, v := range possiblePaths {
			if scores[i] >= threshold && IsGoodPath(possiblePaths, i) {
				goodPaths = append(goodPaths, v)
			}
		}
		if len(goodPaths) == 0 {
			goodPaths = possiblePaths
		}
		geodesCracked = Walk(time - 1, blueprint, goodPaths)
	} else {
		for _, v := range possiblePaths {
			geodesCracked = aoc_library.Max(geodesCracked, v[Geode][Available])
		}
	}
	return geodesCracked
}

func CopyAndUpdateResources(resources [][]int) [][]int {
	var copyResources [][]int
	for i := range resources {
		collectors := resources[i][Collectors]
		available := resources[i][Available] + collectors
		copyResources = append(copyResources, []int{collectors, available})
	}
	return copyResources
}

func CopyResources(resources [][]int) [][]int {
	var copyResources [][]int
	for i := range resources {
		collectors := resources[i][Collectors]
		available := resources[i][Available]
		copyResources = append(copyResources, []int{collectors, available})
	}
	return copyResources
}

func IsGoodPath(possiblePaths [][][]int, index int) bool {
	for i, v := range possiblePaths {
		if i != index {
			betterPath := false // to filter out paths that are worse in all values than some other path
			samePath := true // to filter out duplicate paths
			for j, w := range v {
				for k, x := range w {
					if possiblePaths[index][j][k] < x {
						samePath = false
					}
					if possiblePaths[index][j][k] > x {
						betterPath = true
					}
				}
			}
			if !betterPath && (!samePath || index > i) {
				return false
			}
		}
	}
	return true
}

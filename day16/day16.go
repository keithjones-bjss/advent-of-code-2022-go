package day16

import (
	"advent-of-code-2022/aoc_library"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Node struct {
	flowRate int
	valves   []string
	links    map[string]int
}

func Run(filename string) (int, int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Can't open %v: %v", filename, err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	nodes := make(map[string]Node)
	start := "AA"

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			id, node := Parse(line)
			nodes[id] = node
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	nodes = CollapseNodes(start, nodes)
	CalculatePaths(nodes)

	visited := make(map[string]bool)
	visited[start] = true
	part1 := Walk1(&nodes, visited, start, 30, 0)
	part2 := Walk2(&nodes, visited, start, 26, start, 26, 0, []string{})
	//part2 := 0
	return part1, part2
}

func Parse(line string) (string, Node) {
	var nodeId string
	var result Node
	_, _ = fmt.Sscanf(line, "Valve %s has flow rate=%d", &nodeId, &result.flowRate)
	index := strings.Index(line, "valves") + 7
	if index < 7 {
		index = strings.Index(line, "valve") + 6
	}
	result.valves = strings.Split(line[index:], ", ")
	result.links = make(map[string]int)
	return nodeId, result
}

func CollapseNodes(start string, nodes map[string]Node) map[string]Node {
	collapsedNodes := make(map[string]Node)
	for id, node := range nodes {
		if id == start || node.flowRate != 0 {
			for _, v := range node.valves {
				distance := 1
				var prev []string
				next := v
				for next != start && nodes[next].flowRate == 0 && !aoc_library.Contains(prev, next) {
					prev = append(prev, next)
					for _, x := range nodes[next].valves {
						if !aoc_library.Contains(prev, x) && x != id {
							next = x
							distance++
						}
					}
				}
				if next != id && (next == start || nodes[next].flowRate != 0) {
					node.links[next] = distance
				}
			}
			collapsedNodes[id] = node
		}
	}
	return collapsedNodes
}

func CalculatePaths(nodes map[string]Node) {
	for nodeId, node := range nodes {
		for len(node.links) < len(nodes)-1 {
			for linkNodeId, linkDistance := range node.links {
				linkNode := nodes[linkNodeId]
				for onwardNodeId, onwardDistance := range linkNode.links {
					if onwardNodeId != nodeId {
						currentDistance, ok := node.links[onwardNodeId]
						if !ok || currentDistance > linkDistance+onwardDistance {
							node.links[onwardNodeId] = linkDistance + onwardDistance
						}
					}
				}
			}
		}
	}
}

func Walk1(nodes *map[string]Node, visited map[string]bool, here string, moves int, total int) int {
	bestTotal := total
	for next, distance := range (*nodes)[here].links {
		if visited[next] {
			continue
		}
		movesLeft := moves - distance - 1
		if movesLeft > 0 {
			newTotal := total + movesLeft*(*nodes)[next].flowRate
			copyVisited := make(map[string]bool)
			for k := range visited {
				copyVisited[k] = true
			}
			copyVisited[next] = true
			newTotal = Walk1(nodes, copyVisited, next, movesLeft, newTotal)
			if bestTotal < newTotal {
				bestTotal = newTotal
			}
		}
	}
	return bestTotal
}

func Walk2(
	nodes *map[string]Node,
	visited map[string]bool,
	here1 string,
	moves1 int,
	here2 string,
	moves2 int,
	total int,
	history []string,
) int {
	bestTotal := total
	for next, distance := range (*nodes)[here1].links {
		if visited[next] {
			continue
		}
		movesLeft := moves1 - distance - 1
		if movesLeft > 0 && (*nodes)[next].flowRate > 0 {
			if len(history) == 1 {
				log.Printf("%v I:%v\n", history[0], next)
			}
			newTotal := total + movesLeft*(*nodes)[next].flowRate
			copyVisited := make(map[string]bool)
			for k := range visited {
				copyVisited[k] = true
			}
			copyVisited[next] = true
			newTotal = Walk2(nodes, copyVisited, next, movesLeft, here2, moves2, newTotal, append(history, "I:"+next))
			if bestTotal < newTotal {
				bestTotal = newTotal
			}
		}
	}
	for next, distance := range (*nodes)[here2].links {
		if visited[next] {
			continue
		}
		movesLeft := moves2 - distance - 1
		if movesLeft > 0 && (*nodes)[next].flowRate > 0 {
			if len(history) == 1 {
				log.Printf("%v E:%v\n", history[0], next)
			}
			newTotal := total + movesLeft*(*nodes)[next].flowRate
			copyVisited := make(map[string]bool)
			for k := range visited {
				copyVisited[k] = true
			}
			copyVisited[next] = true
			newTotal = Walk2(nodes, copyVisited, here1, moves1, next, movesLeft, newTotal, append(history, "E:"+next))
			if bestTotal < newTotal {
				bestTotal = newTotal
			}
		}
	}
	return bestTotal
}

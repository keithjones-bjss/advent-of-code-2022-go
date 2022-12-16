package day16

import (
	"advent-of-code-2022/aoc_library"
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Node struct {
	flowRate int
	valves   []string
	links    map[string]int
}

type IndexedNode struct {
	flowRate int
	links    []int
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

	indexedNodes := IndexNodes(nodes)
	var visitedIndex []bool
	for i := 0; i < len(indexedNodes); i++ {
		visitedIndex = append(visitedIndex, false)
	}
	part1 := Walk1(&indexedNodes, &visitedIndex, 0, 30, 0)
	part2 := Walk2(&indexedNodes, &visitedIndex, 0, 26, 0, 26, 0)

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
				prev := []string{id}
				next := v
				for next != start && nodes[next].flowRate == 0 && !aoc_library.Contains(prev, next) {
					prev = append(prev, next)
					for _, x := range nodes[next].valves {
						if !aoc_library.Contains(prev, x) {
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

func IndexNodes(nodes map[string]Node) []IndexedNode {
	var nodeIds []string
	for nodeId := range nodes {
		nodeIds = append(nodeIds, nodeId)
	}
	sort.Strings(nodeIds)
	var nodeIndexes = make(map[string]int)
	for index, nodeId := range nodeIds {
		nodeIndexes[nodeId] = index
	}
	var indexedNodes []IndexedNode
	for _, nodeId := range nodeIds {
		var links []int
		for _, linkNodeId := range nodeIds {
			distance := nodes[nodeId].links[linkNodeId]
			links = append(links, distance)
		}
		indexedNodes = append(indexedNodes, IndexedNode{
			flowRate: nodes[nodeId].flowRate,
			links:    links,
		})
	}
	return indexedNodes
}

func Walk1(nodes *[]IndexedNode, visited *[]bool, here int, moves int, total int) int {
	bestTotal := total
	for next, distance := range (*nodes)[here].links {
		if (*visited)[next] {
			continue
		}
		movesLeft := moves - distance - 1
		if movesLeft > 0 {
			newTotal := total + movesLeft*(*nodes)[next].flowRate
			(*visited)[next] = true
			newTotal = Walk1(nodes, visited, next, movesLeft, newTotal)
			(*visited)[next] = false
			if bestTotal < newTotal {
				bestTotal = newTotal
			}
		}
	}
	return bestTotal
}

func Walk2(
	nodes *[]IndexedNode,
	visited *[]bool,
	here1 int,
	moves1 int,
	here2 int,
	moves2 int,
	total int,
) int {
	bestTotal := total
	if here1 == here2 && moves1 == moves2 {
		bestTotal = IndexedMove(nodes, visited, here1, moves1, here2, moves2, total, bestTotal)
	} else if moves1 > moves2 {
		bestTotal = IndexedMove(nodes, visited, here1, moves1, here2, moves2, total, bestTotal)
		bestTotal = IndexedMove(nodes, visited, here2, moves2, here1, moves1, total, bestTotal)
	} else {
		bestTotal = IndexedMove(nodes, visited, here2, moves2, here1, moves1, total, bestTotal)
		bestTotal = IndexedMove(nodes, visited, here1, moves1, here2, moves2, total, bestTotal)
	}
	return bestTotal
}

func IndexedMove(
	nodes *[]IndexedNode,
	visited *[]bool,
	myNode int,
	myMovesLeft int,
	otherNode int,
	otherMovesLeft int,
	total int,
	bestTotal int,
) int {
	for next, distance := range (*nodes)[myNode].links {
		if (*visited)[next] || distance == 0 {
			continue
		}
		flowRate := (*nodes)[next].flowRate
		movesLeft := myMovesLeft - distance - 1
		if movesLeft > 0 && flowRate > 0 {
			newTotal := total + flowRate*movesLeft
			(*visited)[next] = true
			newTotal = Walk2(
				nodes,
				visited,
				next,
				movesLeft,
				otherNode,
				otherMovesLeft,
				newTotal,
			)
			(*visited)[next] = false
			if bestTotal < newTotal {
				bestTotal = newTotal
			}
		}
	}
	return bestTotal
}

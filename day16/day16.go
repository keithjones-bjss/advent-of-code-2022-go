package day16

import (
	"advent-of-code-2022/aoc_library"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Link struct {
	node     string
	distance int
}

type Node struct {
	flowRate int
	valves   []string
	links    []Link
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

	var path1 = Walk(nodes, State{
		currentNode: []string{start, start},
		trail:       [][]string{{start}, {start}},
		activeNodes: []string{},
		moves:       []int{30, 0},
	})
	log.Printf("%v", path1)
	part1 := path1.total

	//var path2 = Walk(nodes, State{
	//	currentNode: []string{start, start},
	//	trail:       [][]string{{start}, {start}},
	//	activeNodes: []string{},
	//	moves:       []int{26, 26},
	//})
	//log.Printf("%v", path2)
	//part2 := path2.total
	part2 := 0

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
					node.links = append(node.links, Link{next, distance})
				}
			}
			collapsedNodes[id] = node
		}
	}
	return collapsedNodes
}

func Walk(nodes map[string]Node, state State) State {
	// If all valves are open, stop...
	if len(state.activeNodes) == len(nodes)-1 {
		return state
	}
	// Play out all possible move combinations
	nextState := state
	for index, moves := range state.moves {
		if moves > 0 {
			currentNode := state.currentNode[index]
			// Options: (1) activate current node
			if nodes[currentNode].flowRate != 0 && !aoc_library.Contains(state.activeNodes, currentNode) {
				possibleState := state.OpenValve(nodes[currentNode].flowRate, index)
				possibleState = Walk(nodes, possibleState)
				if nextState.total < possibleState.total {
					nextState = possibleState
				}
			}
			// Options: (2) Move elsewhere
			for _, link := range nodes[currentNode].links {
				if link.distance <= state.moves[index] && !aoc_library.Contains(state.trail[index], link.node) {
					possibleState := state.MoveTo(link, index)
					possibleState = Walk(nodes, possibleState)
					if nextState.total < possibleState.total {
						nextState = possibleState
					}
				}
			}
		}
	}
	return nextState
}

type State struct {
	currentNode []string
	trail       [][]string
	moves       []int
	activeNodes []string
	total       int
}

func (state State) MoveTo(link Link, index int) State {
	nodes := append([]string{}, state.currentNode...)
	nodes[1-index] = state.currentNode[1-index]
	nodes[index] = link.node
	trail := append([][]string{}, append([]string{}, state.trail[0]...), append([]string{}, state.trail[1]...))
	trail[index] = append(trail[index], link.node)
	moves := append([]int{}, state.moves...)
	moves[index] = state.moves[index] - link.distance
	return State{
		currentNode: nodes,
		trail:       trail,
		activeNodes: state.activeNodes,
		total:       state.total,
		moves:       moves,
	}
}

func (state State) OpenValve(flowRate int, index int) State {
	moves := append([]int{}, state.moves...)
	moves[index] = state.moves[index] - 1
	trail := append([][]string{}, append([]string{}, state.trail[0]...), append([]string{}, state.trail[1]...))
	trail[index] = append([]string{}, state.currentNode[index])
	return State{
		currentNode: state.currentNode,
		trail:       trail,
		activeNodes: append(state.activeNodes, state.currentNode[index]),
		total:       state.total + flowRate*(state.moves[index]-1),
		moves:       moves,
	}
}

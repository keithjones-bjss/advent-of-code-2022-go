package day13

import (
	"advent-of-code-2022/aoc_library"
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	value    int
	children []*Node
}

func Run(filename string) (int, int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Can't open %v: %v", filename, err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var packets []*Node
	var allPackets []*Node
	part1 := 0
	index := 1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			inOrder := ComparePackets(packets[0], packets[1]) >= 0
			if inOrder {
				part1 += index
			}
			index++
			packets = []*Node{}
		} else {
			root := Node{children: []*Node{}}
			Parse(&root, line)
			packets = append(packets, &root)
			allPackets = append(allPackets, &root)
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if len(packets) == 2 {
		inOrder := ComparePackets(packets[0], packets[1]) >= 0
		if inOrder {
			part1 += index
		}
	}

	part2 := GetDecoderKey(allPackets)

	return part1, part2
}

func Parse(parent *Node, line string) *Node {
	// If line is empty, just return the parent
	if line == "" {
		return parent
	}
	// Create a leaf node
	leaf := Node{children: []*Node{}}
	next := 0
	if line[0] == '[' {
		// If line starts with a bracket, next leaf node is a list
		depth := 0
		closeBracket := 0
		for index, v := range line {
			if v == '[' {
				depth++
			}
			if v == ']' {
				depth--
				if depth == 0 {
					closeBracket = index
					break
				}
			}
		}
		Parse(&leaf, line[1:closeBracket])
		next = closeBracket + 1
	} else {
		// Next item is a plain value
		next = strings.Index(line, ",")
		if next < 0 {
			next = len(line)
		}
		leaf.value, _ = strconv.Atoi(line[:next])
	}
	parent.children = append(parent.children, &leaf)
	if next < len(line) {
		if line[next] == ',' {
			Parse(parent, line[next+1:])
		} else {
			log.Fatalf("Unexpected character %v at %v index %v. Expected comma.", line[next], line, next)
		}
	}
	return parent
}

func ComparePackets(packetA *Node, packetB *Node) int {
	// If we don't have two packets, do nothing
	if packetA == nil || packetB == nil {
		return 0
	}
	// If both values are integers, lower integer should come first
	if packetA.value != 0 && packetB.value != 0 {
		return packetB.value - packetA.value
	}
	// If both values are lists, compare lists
	if packetA.value == 0 && packetB.value == 0 {
		for index, childNode := range packetA.children {
			// If right list runs out of items first, packets are in wrong order
			if len(packetB.children) <= index {
				return len(packetB.children) - len(packetA.children)
			}
			result := ComparePackets(childNode, packetB.children[index])
			if result != 0 {
				return result
			}
		}
		// If left list runs out of items first, packets are in right order
		return len(packetB.children) - len(packetA.children)
	}
	// If exactly one value is an integer, convert it to a list and compare
	if packetA.value != 0 {
		list := Node{children: []*Node{packetA}}
		return ComparePackets(&list, packetB)
	}
	if packetB.value != 0 {
		list := Node{children: []*Node{packetB}}
		return ComparePackets(packetA, &list)
	}
	// No comparison available
	return -1
}

func GetDecoderKey(allPackets []*Node) int {
	root2 := Node{children: []*Node{}}
	allPackets = append(allPackets, Parse(&root2, "[[2]]"))
	root6 := Node{children: []*Node{}}
	allPackets = append(allPackets, Parse(&root6, "[[6]]"))

	// Sort packets
	newOrder := append([]*Node{}, allPackets...)
	for count := 0; count < len(newOrder); count++ {
		for index := 0; index < len(newOrder)-count-1; index++ {
			this := newOrder[index]
			next := newOrder[index+1]
			result := ComparePackets(this, next)
			if result < 0 {
				newOrder[index] = next
				newOrder[index+1] = this
			}
		}
	}

	// Calculate decoder key
	result := 1
	for count := 0; count < len(newOrder); count++ {
		line := newOrder[count].children[0].ToString()
		if line == "[[2]]" || line == "[[6]]" {
			result *= count + 1
		}
	}

	return result
}

func (node *Node) ToString() string {
	if node.value == 0 {
		values := aoc_library.ArrayTranslate(node.children, func(_ int, value *Node) string { return value.ToString() })
		return "[" + strings.Join(values, ",") + "]"
	} else {
		return strconv.Itoa(node.value)
	}
}

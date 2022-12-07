package day7

import (
	"advent-of-code-2022/aoc_library"
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Directory struct {
	name       string
	fileSizes  int
	totalSizes int
	parent     *Directory
	children   []*Directory
}

func Run(filename string) (int, int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Can't open %v: %v", filename, err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	root := Directory{}
	currentPath := &root
	part1 := 0
	var part2 *Directory
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			currentPath = ParseCommand(line, currentPath)
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1, part2 = ParseTree(&root, &root, &root, []string{})

	return part1, part2.totalSizes
}

func ParseTree(root *Directory, cwd *Directory, smallestMatch *Directory, path []string) (int, *Directory) {
	result := 0
	if cwd.totalSizes <= 100000 {
		result += cwd.totalSizes
	}
	if root.totalSizes-cwd.totalSizes < 40000000 && cwd.totalSizes < smallestMatch.totalSizes {
		smallestMatch = cwd
	}
	for _, v := range cwd.children {
		childTotal := 0
		childTotal, smallestMatch = ParseTree(root, v, smallestMatch, append(path, v.name))
		result += childTotal
	}
	log.Printf("[%v] /%v: size %v, total %v, children %v", result,
		strings.Join(path, "/"), cwd.fileSizes, cwd.totalSizes,
		aoc_library.ArrayTranslate(cwd.children, func(_ int, v *Directory) string { return v.name }))
	return result, smallestMatch
}

func ParseCommand(line string, cwd *Directory) *Directory {
	parts := strings.Split(line, " ")
	if len(parts) < 2 {
		return cwd
	}
	if parts[0] == "$" {
		// It's a command
		if parts[1] == "cd" && len(parts) > 2 {
			// Change directory
			if parts[2] == ".." {
				// Go up if we can
				if cwd.parent != nil {
					return cwd.parent
				}
			} else {
				// Go into a directory
				for _, value := range cwd.children {
					if value.name == parts[2] {
						return value
					}
				}
			}
		}
		// We don't care about any other commands
	} else if parts[0] == "dir" {
		// It's a directory
		child := Directory{name: parts[1], parent: cwd}
		cwd.children = append(cwd.children, &child)
	} else {
		// It's a file
		size, _ := strconv.Atoi(parts[0])
		cwd.fileSizes += size
		for dir := cwd; dir != nil; dir = dir.parent {
			dir.totalSizes += size
		}
	}
	return cwd
}

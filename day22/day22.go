package day22

import (
	"advent-of-code-2022/aoc_library"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	Right int = 0
	Down  int = 1
	Left  int = 2
	Up    int = 3
)

type Position struct {
	row    int
	column int
	facing int
}

type Direction struct {
	steps     int
	clockwise bool
}

type Cube struct {
	width    int
	height   int
	faceSize int
	faceGrid [][]int
	faces    []Face
}

type FaceTransform struct {
	id     int
	matrix int
}

type Face struct {
	id          int
	row         int
	column      int
	above       FaceTransform
	below       FaceTransform
	left        FaceTransform
	right       FaceTransform
	orientation int
}

const (
	Move        int = 0
	RotateRight int = 1
	Flip        int = 2
	RotateLeft  int = 3
)

var orderedFaceTransforms = [][][]int{
	// Cube face map
	// ... 0
	// ... |
	// 1 - 2 - 3 - 4
	// ... |
	// ... 5
	// Above, right, below, left
	{{4, Flip}, {3, RotateLeft}, {2, Move}, {1, RotateRight}}, // 0
	{{0, RotateLeft}, {2, Move}, {5, RotateRight}, {4, Move}}, // 1
	{{0, Move}, {3, Move}, {5, Move}, {1, Move}},              // 2
	{{0, RotateRight}, {4, Move}, {5, RotateLeft}, {2, Move}}, // 3
	{{0, Flip}, {1, Move}, {5, Flip}, {3, Move}},              // 4
	{{2, Move}, {3, RotateRight}, {4, Flip}, {1, RotateLeft}}, // 5
}

func Run(filename string) (int, int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Can't open %v: %v", filename, err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var grid [][]rune
	var directions []Direction
	var doneGrid bool

	position := Position{
		row:    0,
		column: 0,
		facing: Right,
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			doneGrid = true
		} else if doneGrid {
			directions = Parse(line)
		} else {
			if len(grid) == 0 {
				position.column = strings.Index(line, ".")
			}
			grid = append(grid, []rune(line))
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	cube := FindCube(grid)
	part1 := Walk(grid, cube, position, directions, false)
	part2 := Walk(grid, cube, position, directions, true)

	return part1, part2
}

func Parse(line string) []Direction {
	var directions []Direction
	next := line
	for next != "" {
		var steps int
		res, _ := fmt.Sscanf(next, "%d", &steps)
		if res == 1 {
			directions = append(directions, Direction{steps: steps})
			next = next[len(fmt.Sprint(steps)):]
		} else {
			directions = append(directions, Direction{clockwise: next[0] == 'R'})
			next = next[1:]
		}
	}
	return directions
}

func FindCube(grid [][]rune) Cube {
	// TODO Fix this. This is all horrible
	var faceGrid [][]int
	max := len(grid)
	width := 4
	height := 4
	for _, v := range grid {
		if max < len(v) {
			max = len(v)
		}
	}
	faceSize := max / 4
	faceId := 0
	var faces []Face
	for i := 0; i < height; i++ {
		var row []int
		for j := 0; j < width; j++ {
			top := faceSize * i
			left := faceSize * j
			empty := len(grid) <= top || len(grid[top]) <= left || grid[top][left] == ' '
			if empty {
				row = append(row, -1)
				//log.Printf("Area %v,%v is empty.", i, j)
			} else {
				row = append(row, faceId)
				faces = append(faces, Face{id: faceId, row: i, column: j})
				faceId++
				//log.Printf("Area %v,%v is a cube face.", i, j)
			}
		}
		faceGrid = append(faceGrid, row)
	}

	// Totally ad hoc logic to build the map that won't work for anything other than the input and test data
	topFace := faces[0]
	orderedFaces := make(map[int]Face)
	orderedFaces[0] = topFace
	below := faceGrid[aoc_library.Mod(topFace.row+1, 4)][topFace.column]
	if below >= 0 {
		//log.Printf("[2] %v is below 0.", below)
		orderedFaces[2] = faces[below]
		left := faceGrid[faces[below].row][aoc_library.Mod(topFace.column-1, 4)]
		if left >= 0 {
			//log.Printf("[1] %v is to the left of %v.", left, below)
			orderedFaces[1] = faces[left]
			left2 := faceGrid[faces[left].row][aoc_library.Mod(faces[left].column-1, 4)]
			if left2 >= 0 {
				//log.Printf("[4] %v is to the left of %v.", left2, left)
				orderedFaces[4] = faces[left2]
			}
		}
		right := faceGrid[faces[below].row][aoc_library.Mod(topFace.column+1, 4)]
		if right >= 0 {
			//log.Printf("[3] %v is to the right of %v.", right, below)
			orderedFaces[3] = faces[right]
		}
		below2 := faceGrid[aoc_library.Mod(faces[below].row+1, 4)][faces[below].column]
		if below2 >= 0 {
			//log.Printf("[5] %v is below %v.", below2, below)
			orderedFaces[5] = faces[below2]
			left2 := faceGrid[faces[below2].row][aoc_library.Mod(faces[below2].column-1, 4)]
			if left2 >= 0 {
				//log.Printf("[1] %v is to the left of %v (rotated right).", left2, below2)
				faces[left2].orientation = RotateRight
				orderedFaces[1] = faces[left2]
				below3 := faceGrid[aoc_library.Mod(faces[left2].row+1, 4)][faces[left2].column]
				if below3 >= 0 {
					//log.Printf("[4] %v is below %v (rotated right).", below3, left2)
					faces[below3].orientation = RotateRight
					orderedFaces[4] = faces[below3]
				}
			}
			right2 := faceGrid[faces[below2].row][aoc_library.Mod(faces[below2].column+1, 4)]
			if right2 >= 0 {
				//log.Printf("[3] %v is to the right of %v (rotated left).", right2, below2)
				faces[right2].orientation = RotateLeft
				orderedFaces[3] = faces[right2]
			}
			below3 := faceGrid[aoc_library.Mod(faces[below2].row+1, 4)][faces[below2].column]
			if below3 >= 0 {
				//log.Printf("[4] %v is below %v (flipped).", below3, below2)
				faces[below3].orientation = Flip
				orderedFaces[4] = faces[below3]
			}
		}
	}
	right := faceGrid[topFace.row][aoc_library.Mod(topFace.column+1, 4)]
	if right >= 0 {
		//log.Printf("%v is to the right of 0.", right)
		faces[right].orientation = RotateRight
		orderedFaces[3] = faces[right]
		if below < 0 {
			rightBelow := faceGrid[aoc_library.Mod(topFace.row+1, 4)][faces[right].column]
			if rightBelow >= 0 {
				//log.Printf("%v is below %v.", rightBelow, right)
				faces[rightBelow].orientation = RotateRight
				orderedFaces[2] = faces[rightBelow]
			}
		}
	}
	left := faceGrid[topFace.row][aoc_library.Mod(topFace.column-1, 4)]
	if left >= 0 {
		//log.Printf("%v is to the left of 0.", left)
		faces[left].orientation = RotateLeft
		orderedFaces[1] = faces[left]
		if below < 0 {
			leftBelow := faceGrid[aoc_library.Mod(topFace.row+1, 4)][faces[left].column]
			if leftBelow >= 0 {
				//log.Printf("%v is below %v.", leftBelow, left)
				faces[leftBelow].orientation = RotateLeft
				orderedFaces[2] = faces[leftBelow]
			}
		}
	}

	for i, v := range orderedFaces {
		transforms := orderedFaceTransforms[i]
		orientation := faces[v.id].orientation
		faces[v.id].above = Transform(orderedFaces, transforms[aoc_library.Mod(0+orientation, 4)], orientation)
		faces[v.id].right = Transform(orderedFaces, transforms[aoc_library.Mod(1+orientation, 4)], orientation)
		faces[v.id].below = Transform(orderedFaces, transforms[aoc_library.Mod(2+orientation, 4)], orientation)
		faces[v.id].left = Transform(orderedFaces, transforms[aoc_library.Mod(3+orientation, 4)], orientation)
		log.Printf("Face %v (index %v) rotate %v: %v,%v A:%v B:%v L:%v R:%v",
			v.id, i, v.orientation, v.row, v.column, faces[v.id].above, faces[v.id].below, faces[v.id].left, faces[v.id].right)
	}

	return Cube{width, height, faceSize, faceGrid, faces}
}

func Transform(orderedFaces map[int]Face, transform []int, sourceOrientation int) FaceTransform {
	targetId := orderedFaces[transform[0]].id
	targetOrientation := orderedFaces[transform[0]].orientation
	//log.Printf("Source orientation: %v Target face: %v Target orientation: %v Transform: %v",
	//	sourceOrientation, targetId, targetOrientation, transform)
	return FaceTransform{
		id:     targetId,
		matrix: aoc_library.Mod(transform[1]+targetOrientation-sourceOrientation, 4),
	}
}

func Walk(grid [][]rune, cube Cube, originalPosition Position, directions []Direction, part2 bool) int {
	position := originalPosition
	for _, dir := range directions {
		if dir.steps == 0 {
			if dir.clockwise {
				position.facing++
			} else {
				position.facing--
			}
			position.facing = aoc_library.Mod(position.facing, 4)
		} else {
			for count := 1; count <= dir.steps; count++ {
				position = CalculateNextPosition(grid, cube, position, part2)
			}
		}
	}
	return 1000*(position.row+1) + 4*(position.column+1) + position.facing
}

var offset = []int{0, 1, 0, -1, 0}

func CalculateNextPosition(grid [][]rune, cube Cube, position Position, part2 bool) Position {
	dx := offset[position.facing+1]
	dy := offset[position.facing]
	var next Position
	if !part2 {
		next = MoveWithWrap(grid, cube, position, dx, dy)
	} else {
		next = MoveByCubeFace(grid, cube, position, dx, dy)
	}
	if grid[next.row][next.column] != '.' {
		return position
	}
	return next
}

func MoveWithWrap(grid [][]rune, cube Cube, originalPosition Position, dx int, dy int) Position {
	position := originalPosition
	nextRow := aoc_library.Mod(position.row+dy, cube.faceSize*cube.height)
	nextColumn := aoc_library.Mod(position.column+dx, cube.faceSize*cube.width)
	for len(grid) <= nextRow || len(grid[nextRow]) <= nextColumn || grid[nextRow][nextColumn] == ' ' {
		nextRow = aoc_library.Mod(nextRow+dy, cube.faceSize*cube.height)
		nextColumn = aoc_library.Mod(nextColumn+dx, cube.faceSize*cube.width)
	}
	position.row = nextRow
	position.column = nextColumn
	return position
}

func MoveByCubeFace(grid [][]rune, cube Cube, originalPosition Position, dx int, dy int) Position {
	position := originalPosition
	nextRow := aoc_library.Mod(position.row+dy, cube.faceSize*cube.height)
	nextColumn := aoc_library.Mod(position.column+dx, cube.faceSize*cube.width)
	if len(grid) <= nextRow || len(grid[nextRow]) <= nextColumn || grid[nextRow][nextColumn] == ' ' {
		faceGridRow := position.row / cube.faceSize
		faceGridColumn := position.column / cube.faceSize
		faceRow := nextRow % cube.faceSize
		faceColumn := nextColumn % cube.faceSize
		face := cube.faces[cube.faceGrid[faceGridRow][faceGridColumn]]
		var transform FaceTransform
		switch position.facing {
		case Up:
			transform = face.above
		case Down:
			transform = face.below
		case Left:
			transform = face.left
		case Right:
			transform = face.right
		}
		log.Printf("Moving from (%d,%d) on face %d to (%d,%d) [%d:%d,%d:%d] on face %v via transform %v, facing %d",
			position.row, position.column, face.id,
			nextRow, nextColumn, faceGridRow, faceRow, faceGridColumn, faceColumn, transform.id, transform.matrix, position.facing)
		switch transform.matrix {
		case Move:
			nextRow = faceRow
			nextColumn = faceColumn
			log.Printf("Moving: row %v col %v", nextRow, nextColumn)
		case RotateLeft:
			nextRow = faceColumn
			nextColumn = cube.faceSize - 1 - faceRow
			log.Printf("Left: row %v col %v", nextRow, nextColumn)
		case RotateRight:
			nextRow = cube.faceSize - 1 - faceColumn
			nextColumn = faceRow
			log.Printf("Right: row %v col %v", nextRow, nextColumn)
		case Flip:
			nextRow = cube.faceSize - 1 - faceRow
			nextColumn = cube.faceSize - 1 - faceColumn
			log.Printf("Flip: row %v col %v", nextRow, nextColumn)
		}
		faceGridRow = cube.faces[transform.id].row
		faceGridColumn = cube.faces[transform.id].column
		position.facing = aoc_library.Mod(position.facing-transform.matrix, 4)
		nextRow += faceGridRow * cube.faceSize
		nextColumn += faceGridColumn * cube.faceSize
		log.Printf("Actually moved to (%d,%d) [%d:%d,%d:%d], now facing %d",
			nextRow, nextColumn, faceGridRow, faceRow, faceGridColumn, faceColumn, position.facing)
	}
	position.row = nextRow
	position.column = nextColumn
	//log.Printf("Normal: %v -> %v", originalPosition, position)
	return position
}

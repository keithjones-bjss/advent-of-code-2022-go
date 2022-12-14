package aoc_library

type Point struct {
	X int
	Y int
}

var directions = map[string]Point{
	"U":   {0, -1},
	"D":   {0, 1},
	"L":   {-1, 0},
	"R":   {1, 0},
	"D+L": {-1, 1},
	"D+R": {1, 1},
}

func (p Point) Move(dir string) Point {
	return Point{p.X + directions[dir].X, p.Y + directions[dir].Y}
}

func (p Point) IsBetween(p1 Point, p2 Point) bool {
	return IsBetween(p.X, p1.X, p2.X) && IsBetween(p.Y, p1.Y, p2.Y)
}

func IsBetween(i int, i1 int, i2 int) bool {
	if i1 <= i2 {
		return i >= i1 && i <= i2
	}
	return i >= i2 && i <= i1
}

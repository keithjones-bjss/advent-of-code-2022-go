package aoc_library

type Point struct {
	X int
	Y int
}

var directions = map[string]Point{
	"U": {0, -1},
	"D": {0, 1},
	"L": {-1, 0},
	"R": {1, 0},
}

func (p Point) Move(dir string) Point {
	return Point{p.X + directions[dir].X, p.Y + directions[dir].Y}
}

func (p Point) DistanceFrom(q Point) int {
	return Abs(p.X-q.X) + Abs(p.Y-q.Y)
}

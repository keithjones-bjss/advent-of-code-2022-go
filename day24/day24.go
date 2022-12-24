package day24

import (
	"advent-of-code-2022/aoc_library"
	"bufio"
	"log"
	"os"
	"strings"
)

func Run(filename string) (int, int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Can't open %v: %v", filename, err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var grid [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			grid = append(grid, []rune(line))
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	initialState := InitialState(grid)
	stateMap := make(map[aoc_library.Point]State)
	stateMap[initialState.expedition] = initialState
	state2, moves := Move(stateMap, 0)

	part1 := moves

	state2.start = initialState.goal
	state2.goal = initialState.start
	state2Map := make(map[aoc_library.Point]State)
	state2Map[state2.start] = state2
	state3, moves2 := Move(state2Map, moves)

	state3.start = initialState.start
	state3.goal = initialState.goal
	state3Map := make(map[aoc_library.Point]State)
	state3Map[state3.start] = state3
	_, moves3 := Move(state3Map, moves2)

	part2 := moves3

	return part1, part2
}

type Blizzard struct {
	position  aoc_library.Point
	direction aoc_library.Point
}

type State struct {
	height     int
	width      int
	start      aoc_library.Point
	goal       aoc_library.Point
	expedition aoc_library.Point
	blizzards  []Blizzard
}

var directions = []aoc_library.Point{
	{X: 0, Y: -1},
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: -1, Y: 0},
	{X: 0, Y: 0},
}

func InitialState(grid [][]rune) State {
	state := State{
		height: len(grid),
		width:  len(grid[0]),
	}
	for y, row := range grid {
		for x, terrain := range row {
			if terrain != '#' {
				if terrain == '.' {
					if y == 0 {
						state.start = aoc_library.Point{X: x, Y: y}
						state.expedition = state.start
					} else if y == len(grid)-1 {
						state.goal = aoc_library.Point{X: x, Y: y}
					}
				} else {
					blizzard := Blizzard{
						position:  aoc_library.Point{X: x, Y: y},
						direction: directions[strings.Index("^>v<", string(terrain))],
					}
					state.blizzards = append(state.blizzards, blizzard)
				}
			}
		}
	}
	return state
}

func Move(states map[aoc_library.Point]State, moves int) (State, int) {
	log.Printf("Moves: %v Possible states: %v", moves, len(states))
	futureStates := make(map[aoc_library.Point]State)
	for _, state := range states {
		if state.expedition == state.goal {
			return state, moves
		}
		nextState := UpdateState(state)
		for _, direction := range directions {
			if TryMove(nextState, direction) {
				futureState := DoMove(nextState, direction)
				futureStates[futureState.expedition] = futureState
				//log.Printf("Move %v: at %v,%v possible move to %v,%v", moves,
				//	state.expedition.X, state.expedition.Y,
				//	nextState.expedition.X, nextState.expedition.Y)
			}
		}
	}
	if len(futureStates) == 0 {
		return State{}, 0
	}
	return Move(futureStates, moves+1)
}

func UpdateState(state State) State {
	newState := State{
		height:     state.height,
		width:      state.width,
		start:      state.start,
		goal:       state.goal,
		expedition: state.expedition,
	}
	for _, blizzard := range state.blizzards {
		x := blizzard.position.X + blizzard.direction.X
		if x >= state.width-1 {
			x = 1
		}
		if x <= 0 {
			x = state.width - 2
		}
		y := blizzard.position.Y + blizzard.direction.Y
		if y >= state.height-1 {
			y = 1
		}
		if y <= 0 {
			y = state.height - 2
		}
		nextBlizzard := Blizzard{
			position:  aoc_library.Point{X: x, Y: y},
			direction: blizzard.direction,
		}
		newState.blizzards = append(newState.blizzards, nextBlizzard)
	}
	return newState
}

func TryMove(state State, direction aoc_library.Point) bool {
	x := state.expedition.X + direction.X
	if x <= 0 || x >= state.width-1 {
		return false
	}
	y := state.expedition.Y + direction.Y
	if y < 0 || y >= state.height {
		return false
	}
	if (y == 0 || y == state.height-1) && x != state.start.X && x != state.goal.X {
		return false
	}
	for _, blizzard := range state.blizzards {
		if x == blizzard.position.X && y == blizzard.position.Y {
			return false
		}
	}
	return true
}

func DoMove(state State, direction aoc_library.Point) State {
	return State{
		height:     state.height,
		width:      state.width,
		start:      state.start,
		goal:       state.goal,
		expedition: aoc_library.Point{X: state.expedition.X + direction.X, Y: state.expedition.Y + direction.Y},
		blizzards:  state.blizzards,
	}
}

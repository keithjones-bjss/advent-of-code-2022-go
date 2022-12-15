package day10

import (
	"advent-of-code-2022/aoc_library"
	"bufio"
	"log"
	"os"
	"strconv"
)

type vm struct {
	clock  int
	x      int
	signal int
	beam   [240]bool
}

func (machine vm) AddX(value int) vm {
	return vm{clock: machine.clock + 2, x: machine.x + value}
}

func (machine vm) Nop() vm {
	return vm{clock: machine.clock + 1, x: machine.x}
}

func (machine vm) WithSignalStrength(old vm) vm {
	monitor1 := (old.clock + 20) / 40
	monitor2 := (machine.clock + 20) / 40
	var signalStrength int
	if monitor1 == monitor2 {
		signalStrength = 0
	} else {
		cycle := (monitor2 * 40) - 20
		signalStrength = old.x * cycle
	}
	return vm{clock: machine.clock, x: machine.x, signal: old.signal + signalStrength}
}

func (machine vm) WithBeam(old vm) vm {
	var beam = [240]bool{}
	for n := 0; n < 240; n++ {
		beam[n] = old.beam[n]
	}
	for cycle := old.clock; cycle < machine.clock; cycle++ {
		position := cycle % 240
		horizontal := position % 40
		beam[position] = aoc_library.Abs(old.x-horizontal) <= 1
	}
	return vm{clock: machine.clock, x: machine.x, signal: machine.signal, beam: beam}
}

func (machine vm) Interpret(line string) vm {
	var next vm
	if line[:4] == "addx" {
		value, _ := strconv.Atoi(line[5:])
		next = machine.AddX(value)
	} else {
		next = machine.Nop()
	}
	return next.WithSignalStrength(machine).WithBeam(machine)
}

func Run(filename string) (int, [240]bool) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Can't open %v: %v", filename, err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	machine := vm{clock: 0, x: 1}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			machine = machine.Interpret(line)
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return machine.signal, machine.beam
}

func Stringify(beam [240]bool) string {
	var display = ""
	for n := 0; n < 240; n++ {
		if beam[n] {
			display += "#"
		} else {
			display += "."
		}
		if n%40 == 39 {
			display += "\n"
		}
	}
	return display
}

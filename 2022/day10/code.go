package day10

import (
	"adventofcode/helper"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type cpu struct {
	registerX int
}

type instruction interface {
	cycle(*cpu) bool
}

type noop struct {
}

func (n *noop) cycle(_ *cpu) bool {
	return true
}

type addX struct {
	toAdd          int
	observedCycles int
}

func (a *addX) cycle(c *cpu) bool {
	a.observedCycles++
	if a.observedCycles == 1 {
		return false
	}
	c.registerX += a.toAdd
	return true
}

var _ instruction = (*noop)(nil)
var _ instruction = (*addX)(nil)

func parseInstructions(inputFilePath string) ([]instruction, error) {
	instructions := []instruction{}
	for fileLine := range helper.FileLineReader(inputFilePath) {
		switch {
		case strings.HasPrefix(fileLine, "addx"):
			toAdd, err := strconv.Atoi(strings.TrimPrefix(fileLine, "addx "))
			if err != nil {
				return nil, errors.Wrap(err, "")
			}
			instructions = append(instructions, &addX{toAdd: toAdd})
		case strings.HasPrefix(fileLine, "noop"):
			instructions = append(instructions, &noop{})
		default:
			return nil, fmt.Errorf("unknown instruction %q", fileLine)
		}
	}
	return instructions, nil
}

func Solve1(inputFilePath string) (int, error) {
	instructions, err := parseInstructions(inputFilePath)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	c := &cpu{registerX: 1}
	currentInstruction := instructions[0]
	instructions = instructions[1:]
	signalStrength := 0
	for i := 1; i <= 220; i++ {

		switch i {
		case 20, 60, 100, 140, 180, 220:
			signalStrength += (i * c.registerX)
		}
		instructionFinished := currentInstruction.cycle(c)

		if instructionFinished {
			if len(instructions) == 0 {
				break
			}
			currentInstruction = instructions[0]
			instructions = instructions[1:]
		}
	}
	return signalStrength, nil
}

func Solve2(inputFilePath string) (int, error) {
	instructions, err := parseInstructions(inputFilePath)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	c := &cpu{registerX: 1}
	currentInstruction := instructions[0]
	instructions = instructions[1:]

	screen := make([][]string, 6)
	for rowIdx := range screen {
		screen[rowIdx] = make([]string, 40)
		for colIdx := range screen[rowIdx] {
			screen[rowIdx][colIdx] = "."
		}
	}

	for cycle := 0; cycle < 240; cycle++ {
		relativeCyclePosition := cycle % 40
		row := cycle / 40

		delta := relativeCyclePosition - c.registerX
		switch delta {
		case -1, 0, 1:
			screen[row][c.registerX+delta] = "#"
		}

		instructionFinished := currentInstruction.cycle(c)

		if instructionFinished {
			if len(instructions) == 0 {
				break
			}
			currentInstruction = instructions[0]
			instructions = instructions[1:]
		}
	}

	for _, row := range screen {
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println("")
	}
	return 0, nil
}

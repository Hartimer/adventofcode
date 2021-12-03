package day8

import (
	"adventofcode/helper"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type instruction struct {
	command  string
	argument int
}

func (i instruction) execute() (int, int) {
	switch i.command {
	case "acc":
		return i.argument, 1
	case "nop":
		return 0, 1
	case "jmp":
		return 0, i.argument
	default:
		panic(fmt.Sprintf("Unknown command %s", i.command))
	}
}

func (i instruction) String() string {
	return fmt.Sprintf("%s %d", i.command, i.argument)
}

type program []instruction

func (p program) flipInstructionAfterIndex(idx int) (program, int, error) {
	p2 := make(program, len(p))
	for idx := range p {
		p2[idx] = p[idx]
	}
	for i := idx + 1; i < len(p2); i++ {
		it := p2[i]
		if it.command == "jmp" {
			it.command = "nop"
			p2[i] = it
			return p2, i, nil
		} else if it.command == "nop" {
			it.command = "jmp"
			p2[i] = it
			return p2, i, nil
		}
	}
	return p, 0, errors.New(fmt.Sprintf("Program %v has no other instructions after index %d", p, idx))
}

func Solve1(inputFilePath string) (int, error) {
	p := program{}
	var err error
	for fileLine := range helper.FileLineReader(inputFilePath) {
		p, err = parseProgramInstruction(p, fileLine)
		if err != nil {
			return 0, errors.Wrap(err, "")
		}
	}
	acc, _, err := solve(p)
	return acc, errors.Wrap(err, "")
}

func Solve2(inputFilePath string) (int, error) {
	originalP := program{}
	var err error
	for fileLine := range helper.FileLineReader(inputFilePath) {
		originalP, err = parseProgramInstruction(originalP, fileLine)
		if err != nil {
			return 0, errors.Wrap(err, "")
		}
	}
	lastInstructionChangedIdx := -1
	acc, finished, err := solve(originalP)
	var p program
	for {
		if err != nil {
			return 0, errors.Wrap(err, "")
		}
		if finished {
			return acc, nil
		}
		p, lastInstructionChangedIdx, err = originalP.flipInstructionAfterIndex(lastInstructionChangedIdx)
		if err != nil {
			return 0, errors.Wrap(err, "")
		}
		log.Printf("Flipping instruction at index %d produces %v", lastInstructionChangedIdx, p)
		acc, finished, err = solve(p)
	}
}

func solve(p program) (int, bool, error) {
	var visitedInstructions = map[int]struct{}{}
	acc, programLine := 0, 0
	for {
		if _, executed := visitedInstructions[programLine]; executed || programLine == len(p) {
			return acc, programLine == len(p), nil
		}
		visitedInstructions[programLine] = struct{}{}
		accOffset, lineOffset := p[programLine].execute()
		acc += accOffset
		programLine += lineOffset
	}
}

func parseProgramInstruction(p program, rawInstruction string) (program, error) {
	i, err := parseInstruction(rawInstruction)
	if err != nil {
		return program{}, errors.Wrap(err, "")
	}
	return append(p, i), nil
}

func parseInstruction(rawInstruction string) (instruction, error) {
	instructionParts := strings.Split(rawInstruction, " ")
	if len(instructionParts) != 2 {
		return instruction{}, errors.New(fmt.Sprintf("Instruction %s should have 2 parts by has %d", rawInstruction, len(instructionParts)))
	}

	i := instruction{command: instructionParts[0]}

	var err error
	i.argument, err = strconv.Atoi(instructionParts[1])
	if err != nil {
		return instruction{}, errors.Wrap(err, "")
	}
	// log.Printf("Raw command %q translates to %s", rawInstruction, i)
	return i, nil
}

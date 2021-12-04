package day14

import (
	"adventofcode/helper"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

const ignoreBit = 'X'

type binaryNumber string

func (b binaryNumber) pad(targetLength int) binaryNumber {
	if len(b) >= targetLength {
		return b
	}
	result := binaryNumber(strings.Repeat("0", targetLength-len(b))) + b
	if len(result) != targetLength {
		panic(fmt.Sprintf("expected %q to have length 36 but has %d", result, len(result)))
	}
	return result
}

func (b binaryNumber) toDecimal() (int64, error) {
	return strconv.ParseInt(string(b), 2, 64)
}

func (b binaryNumber) expand(mask binaryNumber, idx int) []binaryNumber {
	if idx >= len(b) {
		return nil
	}

	var possibilities []binaryNumber
	switch mask[idx] {
	case '0':
		possibilities = append(possibilities, binaryNumber(b[idx]))
	case '1':
		possibilities = append(possibilities, "1")
	default:
		possibilities = append(possibilities, "0", "1")
	}

	nextPossibilities := b.expand(mask, idx+1)
	if len(nextPossibilities) == 0 {
		return possibilities
	}

	var result []binaryNumber
	for _, nextPossibility := range nextPossibilities {
		for _, thisPossibility := range possibilities {
			result = append(result, thisPossibility+nextPossibility)
		}
	}

	return result
}

func (b binaryNumber) trimmedString() string {
	result := string(b)
	for {
		newResult := strings.TrimPrefix(result, "0")
		if len(newResult) == len(result) {
			return result
		}
		result = newResult
	}
}

type assignment struct {
	memoryAddress int64
	value         binaryNumber
}

func (a *assignment) mask(m binaryNumber) {
	for bitIdx, val := range m {
		if val != ignoreBit {
			// originalValue := a.value
			valArr := []rune(a.value)
			valArr[bitIdx] = val
			a.value = binaryNumber(valArr)
			// log.Printf("Original number %s, changing index %d to %s produces %s", originalValue, bitIdx, string(val), a.value)
		}
	}
}

type program struct {
	assignments []*assignment
	memory      map[int]int64
}

func (p *program) run() error {
	var err error
	for _, a := range p.assignments {
		p.memory[int(a.memoryAddress)], err = a.value.toDecimal()
		if err != nil {
			return errors.Wrap(err, "")
		}
	}
	return nil
}
func parseMask(line string) binaryNumber {
	return binaryNumber(strings.TrimPrefix(line, "mask = "))
}

func parseAssignment(line string) (*assignment, error) {
	a := &assignment{}
	lineParts := strings.Split(line, " = ")
	if len(lineParts) != 2 {
		return nil, errors.New(fmt.Sprintf("Expected %s to have 2 parts, but has %d", line, len(lineParts)))
	}
	addressStr := strings.TrimSuffix(strings.TrimPrefix(lineParts[0], "mem["), "]")
	var err error
	a.memoryAddress, err = strconv.ParseInt(addressStr, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	decNumber, err := strconv.ParseInt(lineParts[1], 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	binNumber := binaryNumber(strconv.FormatInt(decNumber, 2))

	a.value = binNumber.pad(36)
	return a, nil
}

func parseAssignmentV2(line string, mask binaryNumber) ([]*assignment, error) {
	var result []*assignment
	lineParts := strings.Split(line, " = ")
	if len(lineParts) != 2 {
		return nil, errors.New(fmt.Sprintf("Expected %s to have 2 parts, but has %d", line, len(lineParts)))
	}

	decNumber, err := strconv.ParseInt(lineParts[1], 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	binNumber := binaryNumber(strconv.FormatInt(decNumber, 2))

	addressStr := strings.TrimSuffix(strings.TrimPrefix(lineParts[0], "mem["), "]")
	addressInt, err := strconv.ParseInt(addressStr, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	addressBinaryStr := binaryNumber(strconv.FormatInt(addressInt, 2)).pad(36)
	expandedAddresses := binaryNumber(addressBinaryStr).expand(mask, 0)

	for _, address := range expandedAddresses {
		a := &assignment{value: binNumber.pad(36)}
		a.memoryAddress, err = strconv.ParseInt(string(address), 2, 64)
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		result = append(result, a)
	}
	return result, nil
}

func Solve1(inputFilePath string) (int, error) {
	fileLineCh := helper.FileLineReader(inputFilePath)

	p := &program{
		memory: map[int]int64{},
	}
	var currentMask binaryNumber
	for fileLine := range fileLineCh {
		if strings.HasPrefix(fileLine, "mask") {
			currentMask = parseMask(fileLine)
			continue
		}
		a, err := parseAssignment(fileLine)
		if err != nil {
			return 0, errors.Wrap(err, "")
		}
		a.mask(currentMask)
		p.assignments = append(p.assignments, a)
	}

	if err := p.run(); err != nil {
		return 0, errors.Wrap(err, "")
	}
	memory := p.memory
	var sum int64 = 0
	for _, m := range memory {
		sum += m
	}
	return int(sum), nil
}

func Solve2(inputFilePath string) (int, error) {
	fileLineCh := helper.FileLineReader(inputFilePath)

	p := &program{
		memory: map[int]int64{},
	}
	var currentMask binaryNumber
	for fileLine := range fileLineCh {
		if strings.HasPrefix(fileLine, "mask") {
			currentMask = parseMask(fileLine)
			continue
		}
		a, err := parseAssignmentV2(fileLine, currentMask)
		if err != nil {
			return 0, errors.Wrap(err, "")
		}
		p.assignments = append(p.assignments, a...)
	}

	if err := p.run(); err != nil {
		return 0, errors.Wrap(err, "")
	}
	memory := p.memory
	var sum int64 = 0
	for _, m := range memory {
		sum += m
	}
	return int(sum), nil
}

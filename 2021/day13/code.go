package day13

import (
	"adventofcode/helper"
	"log"
	"strconv"

	"github.com/pkg/errors"
)

type cellState int

const (
	dot cellState = iota
	empty
)

type coordinate struct {
	x int
	y int
}

type matrix [][]cellState

type foldInstruction struct {
	axis       string
	coordinate int
}

func (m matrix) String() string {
	res := "\n"
	for _, row := range m {
		for _, cell := range row {
			if cell == dot {
				res += "#"
			} else {
				res += "."
			}
		}
		res += "\n"
	}
	return res
}

func (m matrix) foldOnY(y int) matrix {
	var m2 matrix

	topHalf := m[:y]
	bottomHalf := m[y+1:]

	for i, j := 0, len(bottomHalf)-1; i < len(topHalf) && j >= 0; i, j = i+1, j-1 {
		topRow := topHalf[i]
		bottomRow := bottomHalf[j]
		var newRow []cellState
		for idx := range topRow {
			if topRow[idx] == dot || bottomRow[idx] == dot {
				newRow = append(newRow, dot)
			} else {
				newRow = append(newRow, empty)
			}
		}
		m2 = append(m2, newRow)
	}

	return m2
}

func (m matrix) foldOnX(xCut int) matrix {
	var m2 matrix

	for _, row := range m {
		leftHalf := row[:xCut]
		rightHalf := row[xCut+1:]
		var newRow []cellState
		for i, j := 0, len(leftHalf)-1; i < len(rightHalf) && j >= 0; i, j = i+1, j-1 {
			if leftHalf[i] == dot || rightHalf[j] == dot {
				newRow = append(newRow, dot)
			} else {
				newRow = append(newRow, empty)
			}
		}
		m2 = append(m2, newRow)
	}
	return m2
}

func parseInputs(inputFilePath string) (matrix, []foldInstruction, error) {
	var inputs matrix
	var folds []foldInstruction
	coordinates := map[coordinate]struct{}{}
	maxX := 0
	maxY := 0
	parsingFolds := false
	for fileLine := range helper.FileLineReader(inputFilePath) {
		if len(fileLine) == 0 {
			parsingFolds = true
			continue
		}

		if parsingFolds {
			parts, err := helper.SplitAndCheck(fileLine, " ", 3)
			if err != nil {
				return nil, nil, errors.Wrap(err, "")
			}
			foldParts, err := helper.SplitAndCheck(parts[2], "=", 2)
			if err != nil {
				return nil, nil, errors.Wrap(err, "")
			}
			f := foldInstruction{axis: foldParts[0]}
			f.coordinate, err = strconv.Atoi(foldParts[1])
			if err != nil {
				return nil, nil, errors.Wrap(err, "")
			}
			folds = append(folds, f)
		} else {
			coordinateParts, err := helper.SplitAndCheck(fileLine, ",", 2)
			if err != nil {
				return nil, nil, errors.Wrap(err, "")
			}
			c := coordinate{}
			c.x, err = strconv.Atoi(coordinateParts[1])
			if err != nil {
				return nil, nil, errors.Wrap(err, "")
			}
			if c.x > maxX {
				maxX = c.x
			}
			c.y, err = strconv.Atoi(coordinateParts[0])
			if err != nil {
				return nil, nil, errors.Wrap(err, "")
			}
			if c.y > maxY {
				maxY = c.y
			}
			coordinates[c] = struct{}{}
		}
	}

	inputs = make(matrix, maxX+1)
	for x := 0; x <= maxX; x++ {
		row := make([]cellState, maxY+1)
		for y := 0; y <= maxY; y++ {
			c := coordinate{x, y}
			if _, isDot := coordinates[c]; isDot {
				row[y] = dot
			} else {
				row[y] = empty
			}
		}
		inputs[x] = row
	}
	return inputs, folds, nil
}

func Solve1(inputFilePath string) (int, error) {
	inputs, folds, err := parseInputs(inputFilePath)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	firstFold := folds[0]
	if firstFold.axis == "x" {
		inputs = inputs.foldOnX(firstFold.coordinate)
	} else {
		inputs = inputs.foldOnY(firstFold.coordinate)
	}
	count := 0
	for _, row := range inputs {
		for _, cell := range row {
			if cell == dot {
				count++
			}
		}
	}
	return count, nil
}

// Solve2 actually requires printing the matrix in order to get the answer
func Solve2(inputFilePath string) (int, error) {
	inputs, folds, err := parseInputs(inputFilePath)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	for _, fold := range folds {
		if fold.axis == "x" {
			inputs = inputs.foldOnX(fold.coordinate)
		} else {
			inputs = inputs.foldOnY(fold.coordinate)
		}
	}
	log.Printf("Ended up with %s", inputs)
	count := 0
	for _, row := range inputs {
		for _, cell := range row {
			if cell == dot {
				count++
			}
		}
	}
	return count, nil
}

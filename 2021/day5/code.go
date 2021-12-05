package day5

import (
	"adventofcode/helper"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type coordinate struct {
	x int
	y int
}

func newCoordinate(rawCoordinate string) (coordinate, error) {
	cParts := strings.Split(rawCoordinate, ",")
	if len(cParts) != 2 {
		return coordinate{}, errors.New(fmt.Sprintf("Expected %s to have 2 parts but has %d", rawCoordinate, len(cParts)))
	}
	var err error
	c := coordinate{}
	c.x, err = strconv.Atoi(cParts[0])
	if err != nil {
		return coordinate{}, errors.Wrap(err, "")
	}
	c.y, err = strconv.Atoi(cParts[1])
	if err != nil {
		return coordinate{}, errors.Wrap(err, "")
	}
	return c, nil
}

type line struct {
	start coordinate
	end   coordinate
}

var gt = func(a, b int) bool {
	return a >= b
}

var lt = func(a, b int) bool {
	return a <= b
}

func (l line) coordinatesCovered() []coordinate {
	var result []coordinate
	var xIncrementer, yIncrementer int
	xCompFunc, yCompFunc := lt, lt
	if l.isHorizontal() {
		xIncrementer = 0
	} else if l.start.x > l.end.x {
		xIncrementer = -1
		xCompFunc = gt
	} else {
		xIncrementer = 1
	}
	if l.isVertical() {
		yIncrementer = 0
	} else if l.start.y > l.end.y {
		yIncrementer = -1
		yCompFunc = gt
	} else {
		yIncrementer = 1
	}

	if xIncrementer == 0 {
		for y := l.start.y; yCompFunc(y, l.end.y); y += yIncrementer {
			result = append(result, coordinate{x: l.start.x, y: y})
		}
	} else if yIncrementer == 0 {
		for x := l.start.x; xCompFunc(x, l.end.x); x += xIncrementer {
			result = append(result, coordinate{x: x, y: l.start.y})
		}
	} else {
		for x, y := l.start.x, l.start.y; xCompFunc(x, l.end.x) && yCompFunc(y, l.end.y); x, y = x+xIncrementer, y+yIncrementer {
			result = append(result, coordinate{x: x, y: y})
		}
	}
	return result
}

func (l line) isHorizontal() bool {
	return l.start.x == l.end.x
}

func (l line) isVertical() bool {
	return l.start.y == l.end.y
}

func newLine(expression string) (line, error) {
	lParts := strings.Split(expression, " -> ")
	if len(lParts) != 2 {
		return line{}, errors.New(fmt.Sprintf("Expected %s to have 2 parts but has %d", expression, len(lParts)))
	}
	l := line{}
	var err error
	l.start, err = newCoordinate(lParts[0])
	if err != nil {
		return line{}, errors.Wrap(err, "")
	}

	l.end, err = newCoordinate(lParts[1])
	if err != nil {
		return line{}, errors.Wrap(err, "")
	}
	return l, nil
}

type matrix [][]int

func newMatrix(maxX, maxY int) matrix {
	m := make([][]int, maxX+1)
	for x := range m {
		m[x] = make([]int, maxY+1)
	}
	return m
}

func (m matrix) incrementCoordinates(coordinates []coordinate) {
	for _, c := range coordinates {
		m[c.x][c.y]++
	}
}

func (m matrix) countCellsAboveorEqual(n int) int {
	count := 0
	for _, row := range m {
		for _, cell := range row {
			if cell >= n {
				count++
			}
		}
	}
	return count
}

type lineFilter func(line) bool

var horizontalOrVerticalLine lineFilter = func(l line) bool {
	return l.isHorizontal() || l.isVertical()
}

var noop lineFilter = func(_ line) bool {
	return true
}

func solve(inputFilePath string, filter lineFilter) (int, error) {
	maxX, maxY := 0, 0
	var lines []line
	for fileLine := range helper.FileLineReader(inputFilePath) {
		l, err := newLine(fileLine)
		if err != nil {
			return 0, errors.Wrap(err, "")
		}
		maxX = maxInt(maxX, l.start.x)
		maxX = maxInt(maxX, l.end.x)
		maxY = maxInt(maxY, l.start.y)
		maxY = maxInt(maxY, l.end.y)

		lines = append(lines, l)
	}
	m := newMatrix(maxX, maxY)
	for _, l := range lines {
		if filter(l) {
			m.incrementCoordinates(l.coordinatesCovered())
		}
	}

	return m.countCellsAboveorEqual(2), nil
}

func Solve1(inputFilePath string) (int, error) {
	result, err := solve(inputFilePath, horizontalOrVerticalLine)
	return result, errors.Wrap(err, "")
}

func Solve2(inputFilePath string) (int, error) {
	result, err := solve(inputFilePath, noop)
	return result, errors.Wrap(err, "")
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package day11

import (
	"adventofcode/helper"
	"strconv"

	"github.com/pkg/errors"
)

type coordinate struct {
	x int
	y int
}

type matrix [][]int

func (m matrix) get(c coordinate) (int, bool) {
	if c.x < 0 || c.x >= len(m) || c.y < 0 || c.y >= len(m[c.x]) {
		return 0, false
	}
	return m[c.x][c.y], true
}

func (m matrix) getNeighbors(c coordinate) []coordinate {
	var result []coordinate
	offsets := []int{-1, 0, 1}
	for _, xOffset := range offsets {
		for _, yOffset := range offsets {
			if xOffset == 0 && yOffset == 0 {
				continue
			}

			newC := coordinate{x: c.x + xOffset, y: c.y + yOffset}
			if _, isValid := m.get(newC); isValid {
				result = append(result, newC)
			}
		}
	}
	return result
}

func (m matrix) step() (matrix, int) {
	var m2 matrix = make([][]int, len(m))
	flashedCells := map[coordinate]struct{}{}

	var flashQueue []coordinate
	for x, row := range m {
		newRow := make([]int, len(row))
		for y, cell := range row {
			c := coordinate{x, y}
			if _, alreadyFlashed := flashedCells[c]; alreadyFlashed {
				continue
			}
			newEnergy := cell + 1
			if newEnergy > 9 {
				newEnergy = 0
				flashQueue = append(flashQueue, c)
				flashedCells[c] = struct{}{}
			}
			newRow[y] = newEnergy
		}
		m2[x] = newRow
	}

	for len(flashQueue) > 0 {
		q := flashQueue[0]
		flashQueue = flashQueue[1:]
		for _, n := range m2.getNeighbors(q) {
			if _, alreadyFlashed := flashedCells[n]; alreadyFlashed {
				continue
			}
			newEnergy, isValid := m2.get(n)
			if !isValid {
				panic("Why")
			}
			newEnergy++
			if newEnergy > 9 {
				newEnergy = 0
				flashQueue = append(flashQueue, n)
				flashedCells[n] = struct{}{}
			}
			m2[n.x][n.y] = newEnergy
		}
	}

	return m2, len(flashedCells)
}

func parseMatrix(inputFilePath string) (matrix, error) {
	m := matrix{}
	for fileLine := range helper.FileLineReader(inputFilePath) {
		var row []int
		for _, e := range fileLine {
			n, err := strconv.Atoi(string(e))
			if err != nil {
				return nil, errors.Wrap(err, "")
			}
			row = append(row, n)
		}
		m = append(m, row)
	}
	return m, nil
}

func Solve1(inputFilePath string) (int, error) {
	m, err := parseMatrix(inputFilePath)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	flashCount := 0
	for i := 0; i < 100; i++ {
		var partial int
		m, partial = m.step()
		flashCount += partial
	}
	return flashCount, nil
}

func Solve2(inputFilePath string) (int, error) {
	m, err := parseMatrix(inputFilePath)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	var flashCount int
	for i := 1; ; i++ {
		if m, flashCount = m.step(); flashCount == len(m)*len(m[0]) {
			return i, nil
		}
	}
}

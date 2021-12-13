package day18

import (
	"adventofcode/helper"

	"github.com/pkg/errors"
)

type lightState int

const (
	On lightState = iota
	Off
)

type matrix [][]lightState

func (m matrix) get(x, y int) (lightState, bool) {
	if x < 0 || x >= len(m) || y < 0 || y >= len(m[x]) {
		return 0, false
	}
	return m[x][y], true
}

func (m matrix) litNeighbors(x, y int) int {
	var result int
	offsets := []int{-1, 0, 1}
	for _, xOffset := range offsets {
		for _, yOffset := range offsets {
			if xOffset == 0 && yOffset == 0 {
				continue
			}
			newX := x + xOffset
			newY := y + yOffset
			if l, valid := m.get(newX, newY); valid && l == On {
				result++
			}
		}
	}
	return result
}

func (m matrix) isCorder(x, y int) bool {
	return (x == 0 || x == len(m)-1) && (y == 0 || y == len(m[x])-1)
}

func (m matrix) step(cornersAlwaysOn bool) matrix {
	m2 := make(matrix, len(m))
	for x, row := range m {
		newRow := make([]lightState, len(row))
		for y, light := range row {
			if cornersAlwaysOn && m.isCorder(x, y) {
				newRow[y] = On
				continue
			}
			litN := m.litNeighbors(x, y)
			if light == On && (litN == 2 || litN == 3) {
				newRow[y] = On
			} else if light == Off && litN == 3 {
				newRow[y] = On
			} else {
				newRow[y] = Off
			}
		}
		m2[x] = newRow
	}
	return m2
}

func parseMatrix(inputFilePath string) (matrix, error) {
	var m matrix
	for fileLine := range helper.FileLineReader(inputFilePath) {
		row := make([]lightState, len(fileLine))
		for idx, c := range fileLine {
			if c == '.' {
				row[idx] = Off
			} else {
				row[idx] = On
			}
		}
		m = append(m, row)
	}
	return m, nil
}

func Solve1(inputFilePath string, iterations int) (int, error) {
	m, err := parseMatrix(inputFilePath)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	for i := 0; i < iterations; i++ {
		m = m.step(false)
	}
	count := 0
	for _, row := range m {
		for _, l := range row {
			if l == On {
				count++
			}
		}
	}
	return count, nil
}

func Solve2(inputFilePath string, iterations int) (int, error) {
	m, err := parseMatrix(inputFilePath)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	m[0][0] = On
	m[0][len(m[0])-1] = On
	m[len(m)-1][0] = On
	m[len(m)-1][len(m[0])-1] = On
	for i := 0; i < iterations; i++ {
		m = m.step(true)
	}
	count := 0
	for _, row := range m {
		for _, l := range row {
			if l == On {
				count++
			}
		}
	}
	return count, nil
}

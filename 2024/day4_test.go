package _test

import (
	"adventofcode/helper"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type matriz [][]string

func (m matriz) get(x, y int) (string, bool) {
	if x < 0 || x >= len(m) || y < 0 || y >= len(m[x]) {
		return "", false
	}
	return m[x][y], true
}

func (m matriz) search(x, y int, xOffset, yOffset int) bool {
	if m[x][y] != "X" {
		return false
	}

	x += xOffset
	y += yOffset
	if l, valid := m.get(x, y); !valid || l != "M" {
		return false
	}

	x += xOffset
	y += yOffset
	if l, valid := m.get(x, y); !valid || l != "A" {
		return false
	}

	x += xOffset
	y += yOffset
	if l, valid := m.get(x, y); !valid || l != "S" {
		return false
	}
	return true
}

func validPart2(l string) bool {
	return l == "S" || l == "M"
}

func (m matriz) search2(x, y int) bool {
	if m[x][y] != "A" {
		return false
	}
	mCount, sCount := 0, 0
	l, valid := m.get(x-1, y-1)
	if !valid || !validPart2(l) {
		return false
	} else if l == "M" {
		mCount++
	} else if l == "S" {
		sCount++
	}
	if l2, valid := m.get(x+1, y+1); valid && l2 == l {
		return false
	}

	l, valid = m.get(x+1, y-1)
	if !valid || !validPart2(l) {
		return false
	} else if l == "M" {
		mCount++
	} else if l == "S" {
		sCount++
	}
	if l2, valid := m.get(x-1, y+1); valid && l2 == l {
		return false
	}

	if l, valid := m.get(x-1, y+1); !valid || !validPart2(l) {
		return false
	} else if l == "M" {
		mCount++
	} else if l == "S" {
		sCount++
	}
	if l, valid := m.get(x+1, y+1); !valid || !validPart2(l) {
		return false
	} else if l == "M" {
		mCount++
	} else if l == "S" {
		sCount++
	}
	return mCount == 2 && sCount == 2
}

func TestDay4_1(t *testing.T) {
	inputs := []struct {
		filename string
		expected int
	}{
		{
			filename: "day4.1.1.input",
			expected: 18,
		},
		{
			filename: "day4.1.input",
			expected: 2517,
		},
	}

	for _, input := range inputs {
		t.Run(fmt.Sprintf("%s produces %d", input.filename, input.expected), func(t *testing.T) {
			matrix := [][]string{}
			row := 0
			for fileLine := range helper.FileLineReader(input.filename) {
				matrix = append(matrix, make([]string, len(fileLine)))
				for idx, l := range fileLine {
					matrix[row][idx] = string(l)
				}
				row++
			}
			require.NotEmpty(t, matrix)
			total := 0
			possibleOffsets := []int{-1, 0, 1}
			for x, row := range matrix {
				for y := range row {
					for _, xOffset := range possibleOffsets {
						for _, yOffset := range possibleOffsets {
							if xOffset == 0 && yOffset == 0 {
								continue
							}
							valid := matriz(matrix).search(x, y, xOffset, yOffset)
							if valid {
								total++
							}
						}
					}
				}
			}
			require.Equal(t, input.expected, total)
		})
	}
}

func TestDay4_2(t *testing.T) {
	inputs := []struct {
		filename string
		expected int
	}{
		{
			filename: "day4.1.1.input",
			expected: 9,
		},
		{
			filename: "day4.1.input",
			expected: 1960, // not 2018
		},
	}

	for _, input := range inputs {
		t.Run(fmt.Sprintf("%s produces %d", input.filename, input.expected), func(t *testing.T) {
			matrix := [][]string{}
			row := 0
			for fileLine := range helper.FileLineReader(input.filename) {
				matrix = append(matrix, make([]string, len(fileLine)))
				for idx, l := range fileLine {
					matrix[row][idx] = string(l)
				}
				row++
			}
			require.NotEmpty(t, matrix)
			total := 0
			for x, row := range matrix {
				for y := range row {
					valid := matriz(matrix).search2(x, y)
					if valid {
						total++
					}
				}
			}
			require.Equal(t, input.expected, total)
		})
	}
}

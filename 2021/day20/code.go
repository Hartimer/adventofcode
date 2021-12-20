package day20

import (
	"adventofcode/helper"
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

type floorMap [][]int

func (f floorMap) String() string {
	str := "\n"
	for _, row := range f {
		for _, cell := range row {
			if cell == 1 {
				str += "#"
			} else {
				str += "."
			}
		}
		str += "\n"
	}
	return str + "\n"
}

func (f floorMap) TruncateN(n int) floorMap {
	for i := 0; i < n; i++ {
		f = f.Truncate()
	}
	return f
}

func (f floorMap) Truncate() floorMap {
	f = f[1 : len(f)-2]
	for x, row := range f {
		f[x] = row[1 : len(row)-2]
	}
	return f
}

func (f floorMap) ExpandN(n int) floorMap {
	for i := 0; i < n; i++ {
		f = f.Expand()
	}
	return f
}

func (f floorMap) Expand() floorMap {
	emptyRow := make([]int, len(f[0]))
	f = append(f, emptyRow)
	f = append([][]int{emptyRow}, f...)
	for x, row := range f {
		f[x] = append([]int{0}, append(row, 0)...)
	}
	return f
}

func (f floorMap) get(x, y int) int {
	if x < 0 || x >= len(f) || y < 0 || y >= len(f[x]) {
		return 0
	}
	return f[x][y]
}

var offsets = []int{-1, 0, 1}

func (f floorMap) screen(x, y int) (int64, error) {
	str := ""
	for _, xOffset := range offsets {
		for _, yOffset := range offsets {
			newX, newY := x+xOffset, y+yOffset
			str += fmt.Sprint(f.get(newX, newY))
		}
	}
	return strconv.ParseInt(str, 2, 64)
}

func (f floorMap) enhance(formula string) (floorMap, error) {
	f2 := make(floorMap, len(f))
	for x, row := range f {
		newRow := make([]int, len(row))
		for y := range row {
			res, err := f.screen(x, y)
			if err != nil {
				return nil, errors.Wrap(err, "")
			}
			if string(formula[res]) == "#" {
				newRow[y] = 1
			}
		}
		f2[x] = newRow
	}
	return f2, nil
}

func (f floorMap) countLit() int {
	count := 0
	for _, row := range f {
		for _, cell := range row {
			count += cell
		}
	}
	return count
}

func parseInputs(inputFilePath string) (string, floorMap, error) {
	parsingFormula := true
	var formula string
	f := floorMap{}
	for fileLine := range helper.FileLineReader(inputFilePath) {
		if len(fileLine) == 0 {
			parsingFormula = false
			continue
		}
		if parsingFormula {
			formula = fileLine
		} else {
			row := make([]int, len(fileLine))
			for x, c := range fileLine {
				if c == '.' {
					row[x] = 0
				} else {
					row[x] = 1
				}
			}
			f = append(f, row)
		}
	}
	return formula, f, nil
}

func solve(inputFilePath string, iteractions int) (int, error) {
	formula, f, err := parseInputs(inputFilePath)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	f = f.ExpandN(iteractions * 3)
	for i := 0; i < iteractions; i++ {
		f, err = f.enhance(formula)
		if err != nil {
			return 0, errors.Wrap(err, "")
		}
	}
	return f.TruncateN(iteractions).countLit(), nil
}

func Solve1(inputFilePath string) (int, error) {
	return solve(inputFilePath, 2)
}

func Solve2(inputFilePath string) (int, error) {
	return solve(inputFilePath, 50)
}

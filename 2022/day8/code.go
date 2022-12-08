package day8

import (
	"adventofcode/helper"
	"math"
	"strconv"

	"github.com/pkg/errors"
)

type treeArea [][]int

func (t treeArea) isVisible(x, y int) bool {
	if x < 0 || x >= len(t) || y < 0 || y >= len(t[x]) {
		return false
	}

	// Edges are always visible
	if x == 0 || y == 0 || x == len(t)-1 || y == len(t[x])-1 {
		return true
	}

	treeHeight := t[x][y]

	// check left
	neighborY := y - 1
	for ; neighborY >= 0; neighborY-- {
		if t[x][neighborY] >= treeHeight {
			break
		}
	}
	if neighborY == -1 {
		return true
	}
	// check right
	neighborY = y + 1
	for ; neighborY < len(t[x]); neighborY++ {
		if t[x][neighborY] >= treeHeight {
			break
		}
	}
	if neighborY == len(t[x]) {
		return true
	}
	// check top
	neighborX := x - 1
	for ; neighborX >= 0; neighborX-- {
		if t[neighborX][y] >= treeHeight {
			break
		}
	}
	if neighborX == -1 {
		return true
	}
	// check bottom
	neighborX = x + 1
	for ; neighborX < len(t); neighborX++ {
		if t[neighborX][y] >= treeHeight {
			break
		}
	}
	return neighborX == len(t)
}

func (t treeArea) scenicScore(x, y int) int {
	if x < 0 || x >= len(t) || y < 0 || y >= len(t[x]) {
		return -1
	}

	// Edges are always zero
	if x == 0 || y == 0 || x == len(t)-1 || y == len(t[x])-1 {
		return 0
	}

	treeHeight := t[x][y]

	leftScore, rightScore, topScore, bottomScore := 0, 0, 0, 0
	// check left
	for neighborY := y - 1; neighborY >= 0; neighborY-- {
		leftScore++
		if t[x][neighborY] >= treeHeight {
			break
		}
	}
	// check right
	for neighborY := y + 1; neighborY < len(t[x]); neighborY++ {
		rightScore++
		if t[x][neighborY] >= treeHeight {
			break
		}
	}
	// check top
	for neighborX := x - 1; neighborX >= 0; neighborX-- {
		topScore++
		if t[neighborX][y] >= treeHeight {
			break
		}
	}
	// check bottom
	for neighborX := x + 1; neighborX < len(t); neighborX++ {
		bottomScore++
		if t[neighborX][y] >= treeHeight {
			break
		}
	}
	return rightScore * leftScore * topScore * bottomScore
}

func (t treeArea) countVisible() int {
	counter := 0
	for x := 0; x < len(t); x++ {
		for y := 0; y < len(t[x]); y++ {
			if t.isVisible(x, y) {
				counter++
			}
		}
	}
	return counter
}

func (t treeArea) maxScore() int {
	maxScore := math.MinInt
	for x := 0; x < len(t); x++ {
		for y := 0; y < len(t[x]); y++ {
			treeScore := t.scenicScore(x, y)
			if treeScore > maxScore {
				maxScore = treeScore
			}
		}
	}
	return maxScore
}

func parseTreeArea(inputFilePath string) (treeArea, error) {
	t := treeArea{}
	for fileLine := range helper.FileLineReader(inputFilePath) {
		t = append(t, []int{})
		for _, heightStr := range fileLine {
			height, err := strconv.Atoi(string(heightStr))
			if err != nil {
				return treeArea{}, errors.Wrap(err, "")
			}
			t[len(t)-1] = append(t[len(t)-1], height)
		}
	}
	return t, nil
}

func Solve1(inputFilePath string) (int, error) {
	t, err := parseTreeArea(inputFilePath)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	return t.countVisible(), nil
}

func Solve2(inputFilePath string) (int, error) {
	t, err := parseTreeArea(inputFilePath)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	return t.maxScore(), nil
}

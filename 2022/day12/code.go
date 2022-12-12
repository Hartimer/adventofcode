package day12

import (
	"adventofcode/helper"
	"fmt"
	"math"
)

type position struct {
	x, y int
}

type heightMap [][]rune

func NewHeightMap(inputFilePath string) heightMap {
	h := heightMap{}
	for fileLine := range helper.FileLineReader(inputFilePath) {
		row := make([]rune, len(fileLine))
		for idx, c := range fileLine {
			row[idx] = c
		}
		h = append(h, row)
	}
	return h
}

func (h heightMap) heightOf(p position) (rune, bool) {
	if p.x < 0 || p.y < 0 || p.x >= len(h) || p.y >= len(h[p.x]) {
		return 0, false
	}
	return h[p.x][p.y], true
}

func (h heightMap) nextStepsFrom(p position) ([]position, bool) {
	deltas := []position{
		{x: 0, y: -1},
		{x: 0, y: 1},
		{x: -1, y: 0},
		{x: 1, y: 0},
	}

	currentHeight := h[p.x][p.y]
	if currentHeight == 'S' {
		currentHeight = 'a'
	}

	possibleNextSteps := []position{}
	for _, delta := range deltas {
		newPosition := position{
			x: p.x + delta.x,
			y: p.y + delta.y,
		}

		nextHeight, isValid := h.heightOf(newPosition)
		if isValid {
			if nextHeight == 'E' {
				return nil, true
			}
			if math.Abs(float64(nextHeight-currentHeight)) <= 1 {
				possibleNextSteps = append(possibleNextSteps, newPosition)
			}
		}
	}

	return possibleNextSteps, false
}

type visited map[position]struct{}

func (v visited) contains(p position) bool {
	_, contains := v[p]
	return contains
}

func (v visited) add(p position) visited {
	newVisited := map[position]struct{}{p: {}}
	for k, val := range v {
		newVisited[k] = val
	}
	return newVisited
}

func findEnd(h heightMap, currentPosition position, v visited) (int, bool) {
	if v.contains(currentPosition) {
		return 0, false
	}
	v = v.add(currentPosition)
	nextSteps, foundEnd := h.nextStepsFrom(currentPosition)
	if foundEnd {
		return 1, true
	}
	if len(nextSteps) == 0 {
		return 0, false
	}

	lowerCount := math.MaxInt
	for _, nextStep := range nextSteps {
		count, found := findEnd(h, nextStep, v)
		if found && count < lowerCount {
			lowerCount = count
		}
	}
	return lowerCount + 1, lowerCount < math.MaxInt
}

func Solve1(inputFilePath string) (int, error) {
	h := NewHeightMap(inputFilePath)
	shortestPath, found := findEnd(h, position{x: 0, y: 0}, visited{})
	if !found {
		return 0, fmt.Errorf("no path found")
	}

	return shortestPath, nil
}

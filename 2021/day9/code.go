package day9

import (
	"adventofcode/helper"
	"sort"
	"strconv"

	"github.com/pkg/errors"
)

type coordinate struct {
	x int
	y int
}

type floorMap [][]int

func (f floorMap) get(c coordinate) (int, bool) {
	if c.x < 0 || c.y < 0 || c.x >= len(f) || c.y >= len(f[c.x]) {
		return 0, false
	}
	return f[c.x][c.y], true
}

func (f floorMap) neighbors(c coordinate) []int {
	var result []int
	for _, c := range f.neighborCoordinates(c) {
		result = append(result, f[c.x][c.y])
	}
	return result
}

func (f floorMap) neighborCoordinates(c coordinate) []coordinate {
	var result []coordinate
	for _, offset := range []int{-1, 1} {
		if _, valid := f.get(coordinate{x: c.x, y: c.y + offset}); valid {
			result = append(result, coordinate{x: c.x, y: c.y + offset})
		}
		if _, valid := f.get(coordinate{x: c.x + offset, y: c.y}); valid {
			result = append(result, coordinate{x: c.x + offset, y: c.y})
		}
	}
	return result
}

func (f floorMap) getBasinSize(lowersPoint coordinate) int {
	seenCoordinates := map[coordinate]struct{}{lowersPoint: {}}
	toSeeCoordinates := map[coordinate]struct{}{}

	neighborCoordinates := f.neighborCoordinates(lowersPoint)
	for idx := range neighborCoordinates {
		neighborCoordinate := neighborCoordinates[idx]
		if _, seen := seenCoordinates[neighborCoordinate]; seen {
			continue
		}
		if n, valid := f.get(neighborCoordinate); valid && n != 9 {
			toSeeCoordinates[neighborCoordinate] = struct{}{}
		}
	}

	for len(toSeeCoordinates) > 0 {
		for nextCoordinate := range toSeeCoordinates {
			delete(toSeeCoordinates, nextCoordinate)
			if _, seen := seenCoordinates[nextCoordinate]; seen {
				break
			}
			seenCoordinates[nextCoordinate] = struct{}{}
			neighborCoordinates := f.neighborCoordinates(nextCoordinate)
			for idx := range neighborCoordinates {
				neighborCoordinate := neighborCoordinates[idx]
				if _, seen := seenCoordinates[neighborCoordinate]; seen {
					continue
				}
				if n, valid := f.get(neighborCoordinate); valid && n != 9 {
					toSeeCoordinates[neighborCoordinate] = struct{}{}
				}
			}
		}
	}

	return len(seenCoordinates)
}

func parseMap(inputFilePath string) (floorMap, error) {
	var m floorMap
	for fileLine := range helper.FileLineReader(inputFilePath) {
		var row []int
		for _, c := range fileLine {
			height, err := strconv.Atoi(string(c))
			if err != nil {
				return nil, errors.Wrap(err, "")
			}
			row = append(row, height)
		}
		m = append(m, row)
	}

	return m, nil
}

func Solve1(inputFilePath string) (int, error) {
	m, err := parseMap(inputFilePath)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	sum := 0
	for x, row := range m {
		for y, height := range row {
			isLowestPoint := true
			for _, neighborHeight := range m.neighbors(coordinate{x, y}) {
				if neighborHeight <= height {
					isLowestPoint = false
					break
				}
			}
			if isLowestPoint {
				sum += height + 1
			}
		}
	}
	return sum, nil
}

func Solve2(inputFilePath string) (int, error) {
	m, err := parseMap(inputFilePath)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	var lowestPoints []coordinate
	for x, row := range m {
		for y, height := range row {
			isLowestPoint := true
			for _, neighborHeight := range m.neighbors(coordinate{x, y}) {
				if neighborHeight <= height {
					isLowestPoint = false
					break
				}
			}
			if isLowestPoint {
				lowestPoints = append(lowestPoints, coordinate{x: x, y: y})
			}
		}
	}
	basins := map[coordinate]int{}
	for _, lowestPoint := range lowestPoints {
		basins[lowestPoint] = m.getBasinSize(lowestPoint)
	}
	var sizes []int
	for _, size := range basins {
		sizes = append(sizes, size)
	}
	sort.Ints(sizes)
	sum := 1
	for _, n := range sizes[len(sizes)-3:] {
		sum *= n
	}
	return sum, nil
}

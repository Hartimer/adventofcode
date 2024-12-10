package _test

import (
	"adventofcode/helper"
	"fmt"
	"maps"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	North = Coordinate{X: -1, Y: 0}
	South = Coordinate{X: 1, Y: 0}
	East  = Coordinate{X: 0, Y: 1}
	West  = Coordinate{X: 0, Y: -1}
)

type map10 [][]int

func (m map10) get(c Coordinate) (int, bool) {
	if c.X < 0 || c.X >= len(m) ||
		c.Y < 0 || c.Y >= len(m[c.X]) {
		return 0, false
	}
	return m[c.X][c.Y], true
}

func (m map10) getTrailHeads() []Coordinate {
	var result []Coordinate
	for x, row := range m {
		for y, height := range row {
			if height == 0 {
				result = append(result, Coordinate{
					X: x,
					Y: y,
				})
			}
		}
	}
	return result
}

func (m map10) stepToPeak(currentPosition Coordinate) map[Coordinate]struct{} {
	peaksReached := map[Coordinate]struct{}{}
	currentHeight := m[currentPosition.X][currentPosition.Y]
	if currentHeight == 9 {
		return map[Coordinate]struct{}{currentPosition: {}}
	}
	for _, direction := range []Coordinate{North, South, East, West} {
		nextPosition := Coordinate{
			X: currentPosition.X + direction.X,
			Y: currentPosition.Y + direction.Y,
		}
		nextStep, isValid := m.get(nextPosition)
		if !isValid || nextStep != currentHeight+1 {
			continue
		}
		subPeaksReached := m.stepToPeak(nextPosition)
		for peak := range subPeaksReached {
			peaksReached[peak] = struct{}{}
		}
	}

	return peaksReached
}

func (m map10) stepToPeakDistinct(currentPosition Coordinate) map[Coordinate]int {
	currentHeight := m[currentPosition.X][currentPosition.Y]
	if currentHeight == 9 {
		return map[Coordinate]int{currentPosition: 1}
	}
	peaksReached := map[Coordinate]int{}
	for _, direction := range []Coordinate{North, South, East, West} {
		nextPosition := Coordinate{
			X: currentPosition.X + direction.X,
			Y: currentPosition.Y + direction.Y,
		}
		nextStep, isValid := m.get(nextPosition)
		if !isValid || nextStep != currentHeight+1 {
			continue
		}
		subPeaksReached := m.stepToPeakDistinct(nextPosition)
		for peak, subReachCount := range subPeaksReached {
			reachCount, exists := peaksReached[peak]
			if !exists {
				reachCount = 0
			}
			peaksReached[peak] = (reachCount + subReachCount)
		}
	}

	return peaksReached
}

func TestDay10_1(t *testing.T) {
	inputs := []struct {
		filename string
		expected int
	}{
		{
			filename: "day10.1.1.input",
			expected: 36,
		},
		{
			filename: "day10.1.input",
			expected: 496,
		},
	}

	for _, input := range inputs {
		t.Run(fmt.Sprintf("%s produces %d", input.filename, input.expected), func(t *testing.T) {
			theMap := map10{}
			for fileLine := range helper.FileLineReader(input.filename) {
				var row []int
				for _, cell := range fileLine {
					height := int(cell - '0')
					row = append(row, height)
				}
				theMap = append(theMap, row)
			}
			total := 0
			trailHeads := theMap.getTrailHeads()
			for _, trailHead := range trailHeads {
				peaksReached := theMap.stepToPeak(trailHead)
				total += len(peaksReached)
			}
			require.Equal(t, input.expected, total)
		})
	}
}

func TestDay10_2(t *testing.T) {
	inputs := []struct {
		filename string
		expected int
	}{
		{
			filename: "day10.1.1.input",
			expected: 81,
		},
		{
			filename: "day10.1.input",
			expected: 1120,
		},
	}

	for _, input := range inputs {
		t.Run(fmt.Sprintf("%s produces %d", input.filename, input.expected), func(t *testing.T) {
			theMap := map10{}
			for fileLine := range helper.FileLineReader(input.filename) {
				var row []int
				for _, cell := range fileLine {
					height := int(cell - '0')
					row = append(row, height)
				}
				theMap = append(theMap, row)
			}
			total := 0
			trailHeads := theMap.getTrailHeads()
			for _, trailHead := range trailHeads {
				peaksReached := theMap.stepToPeakDistinct(trailHead)
				for distinctPaths := range maps.Values(peaksReached) {
					total += distinctPaths
				}
			}
			require.Equal(t, input.expected, total)
		})
	}
}

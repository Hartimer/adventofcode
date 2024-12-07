package _test

import (
	"adventofcode/helper"
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

type Direction struct {
	OffsetX, OffsetY int
}

func (d Direction) rotateRight() Direction {
	switch d {
	case Up:
		return Right
	case Right:
		return Down
	case Down:
		return Left
	case Left:
		return Up
	}
	panic("should not happen")
}

type Coordinate struct {
	X, Y int
}

var (
	Up    = Direction{OffsetX: -1, OffsetY: 0}
	Down  = Direction{OffsetX: 1, OffsetY: 0}
	Right = Direction{OffsetX: 0, OffsetY: 1}
	Left  = Direction{OffsetX: 0, OffsetY: -1}
)

type Floor5 [][]rune

func (f Floor5) move(startPosition Coordinate, direction Direction) (Coordinate, bool, bool) {
	newPosition := Coordinate{
		X: startPosition.X + direction.OffsetX,
		Y: startPosition.Y + direction.OffsetY,
	}
	if newPosition.X < 0 || newPosition.X >= len(f) ||
		newPosition.Y < 0 || newPosition.Y >= len(f[newPosition.X]) {
		return newPosition, true, true
	}
	if f[newPosition.X][newPosition.Y] == '.' {
		f[startPosition.X][startPosition.Y] = '.'
		return newPosition, true, false
	}
	return startPosition, false, false
}

func (f Floor5) moveCheckBlock(startPosition Coordinate,
	direction Direction, visitedPositions map[Coordinate][]Direction) (Coordinate, bool, bool, bool) {
	newPosition := Coordinate{
		X: startPosition.X + direction.OffsetX,
		Y: startPosition.Y + direction.OffsetY,
	}
	if newPosition.X < 0 || newPosition.X >= len(f) ||
		newPosition.Y < 0 || newPosition.Y >= len(f[newPosition.X]) {
		return newPosition, true, false, true
	}
	// turnRightDirection := direction.rotateRight()
	// turnRightPosition := Coordinate{
	// 	X: startPosition.X + turnRightDirection.OffsetX,
	// 	Y: startPosition.Y + turnRightDirection.OffsetY,
	// }
	// wouldLoop := f.hitsObstacle(turnRightPosition, turnRightDirection, visitedPositions)
	wouldLoop := f.couldLoop(startPosition, direction.rotateRight(), visitedPositions)
	if f[newPosition.X][newPosition.Y] == '.' {
		f[startPosition.X][startPosition.Y] = '.'
		return newPosition, true, wouldLoop, false
	}
	return startPosition, false, wouldLoop, false
}

func (f Floor5) hitsObstacle(startPosition Coordinate, direction Direction, visitedPositions map[Coordinate][]Direction) bool {
	if observedDirections, hasBeenVisited := visitedPositions[startPosition]; !hasBeenVisited || !slices.Contains(observedDirections, direction) {
		return false
	}
	newPosition := Coordinate{
		X: startPosition.X + direction.OffsetX,
		Y: startPosition.Y + direction.OffsetY,
	}
	if newPosition.X < 0 || newPosition.X >= len(f) ||
		newPosition.Y < 0 || newPosition.Y >= len(f[newPosition.X]) {
		return false
	}
	if f[newPosition.X][newPosition.Y] == '#' {
		return true
	}
	return f.hitsObstacle(newPosition, direction, visitedPositions)
}

func (f Floor5) couldLoop(original Coordinate, currentDirection Direction, originalVisitedPositons map[Coordinate][]Direction) bool {
	currentPosition := original
	visitedPositons := make(map[Coordinate][]Direction)
	for k, v := range originalVisitedPositons {
		newV := make([]Direction, len(v))
		copy(newV, v)
		visitedPositons[k] = v
	}
	newPosition, didMove, movedOutside := f.move(currentPosition, currentDirection)
	for !movedOutside && newPosition != original {
		if didMove {
			visitedPositons[newPosition] = append(visitedPositons[newPosition], currentDirection)
			currentPosition = newPosition
		} else {
			currentDirection = currentDirection.rotateRight()
			visitedPositons[currentPosition] = append(visitedPositons[currentPosition], currentDirection)
		}
		newPosition, didMove, movedOutside = f.move(currentPosition, currentDirection)
	}
	return newPosition == original
}

func TestDay6_1(t *testing.T) {
	inputs := []struct {
		filename string
		expected int
	}{
		{
			filename: "day6.1.1.input",
			expected: 41,
		},
		{
			filename: "day6.1.input",
			expected: 4939,
		},
	}

	for _, input := range inputs {
		t.Run(fmt.Sprintf("%s produces %d", input.filename, input.expected), func(t *testing.T) {
			floor := Floor5{}
			var currentPosition Coordinate
			for fileLine := range helper.FileLineReader(input.filename) {
				row := make([]rune, 0, len(fileLine))
				for idx, cell := range fileLine {
					row = append(row, cell)
					if cell == '^' {
						currentPosition.X = len(floor)
						currentPosition.Y = idx
					}
				}
				floor = append(floor, row)
			}

			visitedPositons := map[Coordinate]struct{}{
				currentPosition: {},
			}

			currentDirection := Up
			newPosition, didMove, movedOutside := floor.move(currentPosition, currentDirection)
			for !movedOutside {
				if didMove {
					visitedPositons[newPosition] = struct{}{}
					currentPosition = newPosition
				} else {
					currentDirection = currentDirection.rotateRight()
				}
				newPosition, didMove, movedOutside = floor.move(currentPosition, currentDirection)
			}
			require.Equal(t, input.expected, len(visitedPositons))
		})
	}
}
func TestDay6_2(t *testing.T) {
	inputs := []struct {
		filename string
		expected int
	}{
		// {
		// 	filename: "day6.1.1.input",
		// 	expected: 6,
		// },
		{
			filename: "day6.1.input",
			expected: 0,
		},
	}

	for _, input := range inputs {
		t.Run(fmt.Sprintf("%s produces %d", input.filename, input.expected), func(t *testing.T) {
			floor := Floor5{}
			var currentPosition Coordinate
			for fileLine := range helper.FileLineReader(input.filename) {
				row := make([]rune, 0, len(fileLine))
				for idx, cell := range fileLine {
					row = append(row, cell)
					if cell == '^' {
						currentPosition.X = len(floor)
						currentPosition.Y = idx
					}
				}
				floor = append(floor, row)
			}

			visitedPositons := map[Coordinate][]Direction{
				currentPosition: {Up},
			}

			currentDirection := Up
			possibleBlocks := 0
			var couldBlock bool
			newPosition, didMove, movedOutside := floor.move(currentPosition, currentDirection)
			for !movedOutside {
				if didMove {
					visitedPositons[newPosition] = append(visitedPositons[newPosition], currentDirection)
					currentPosition = newPosition
				} else {
					currentDirection = currentDirection.rotateRight()
					visitedPositons[currentPosition] = append(visitedPositons[currentPosition], currentDirection)
				}
				newPosition, didMove, couldBlock, movedOutside = floor.moveCheckBlock(currentPosition, currentDirection, visitedPositons)
				if couldBlock {
					possibleBlocks++
				}
			}
			require.Equal(t, input.expected, possibleBlocks)
		})
	}
}

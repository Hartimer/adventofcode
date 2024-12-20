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
		return newPosition, true, false
	}
	return startPosition, false, false
}

func (f Floor5) couldLoop(original Coordinate, originalDirection Direction, originalVisitedPositons map[Coordinate][]Direction) (Coordinate, bool) {
	visitedPositons := make(map[Coordinate][]Direction)
	for k, v := range originalVisitedPositons {
		newV := make([]Direction, len(v))
		copy(newV, v)
		visitedPositons[k] = newV
	}
	blockPosition, didMove, movedOutside := f.move(original, originalDirection)
	if !didMove || movedOutside {
		return original, false
	}
	originalBlockContent := f[blockPosition.X][blockPosition.Y]
	f[blockPosition.X][blockPosition.Y] = '#'
	defer func() {
		f[blockPosition.X][blockPosition.Y] = originalBlockContent
	}()
	currentDirection := originalDirection

	newPosition, didMove, movedOutside := f.move(original, currentDirection)
	for !movedOutside {
		if didMove {
			directionsInPosition, seenPosition := visitedPositons[newPosition]
			if seenPosition && slices.Contains(directionsInPosition, currentDirection) {
				return blockPosition, true
			}
		} else {
			currentDirection = currentDirection.rotateRight()
		}
		visitedPositons[newPosition] = append(visitedPositons[newPosition], currentDirection)
		newPosition, didMove, movedOutside = f.move(newPosition, currentDirection)
	}

	// newPosition, didMove, movedOutside := f.move(original, currentDirection)
	// for !movedOutside {
	// 	if didMove {
	// 		directionsInPosition, seenPosition := visitedPositons[newPosition]
	// 		if seenPosition && slices.Contains(directionsInPosition, currentDirection) {
	// 			return blockPosition, true
	// 		}
	// 	} else {
	// 		currentDirection = currentDirection.rotateRight()
	// 	}
	// 	visitedPositons[newPosition] = append(visitedPositons[newPosition], currentDirection)
	// 	newPosition, didMove, movedOutside = f.move(newPosition, currentDirection)
	// }
	return blockPosition, false
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
					if cell == '^' {
						currentPosition.X = len(floor)
						currentPosition.Y = idx
						cell = '.'
					}
					row = append(row, cell)
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
		{
			filename: "day6.1.1.input",
			expected: 6,
		},
		{
			filename: "day6.1.input",
			expected: 1434, // not 1456 (too high), got the result elsewhere
		},
	}

	for _, input := range inputs {
		t.Run(fmt.Sprintf("%s produces %d", input.filename, input.expected), func(t *testing.T) {
			floor := Floor5{}
			var currentPosition Coordinate
			for fileLine := range helper.FileLineReader(input.filename) {
				row := make([]rune, 0, len(fileLine))
				for idx, cell := range fileLine {
					if cell == '^' {
						currentPosition.X = len(floor)
						currentPosition.Y = idx
						cell = '.'
					}
					row = append(row, cell)
				}
				floor = append(floor, row)
			}

			visitedPositons := map[Coordinate][]Direction{
				currentPosition: {Up},
			}

			currentDirection := Up
			possibleBlocks := map[Coordinate]struct{}{}
			newPosition, didMove, movedOutside := floor.move(currentPosition, currentDirection)
			for !movedOutside {
				blockPosition, couldBlock := floor.couldLoop(newPosition, currentDirection, visitedPositons)
				if couldBlock {
					possibleBlocks[blockPosition] = struct{}{}
				}
				if !didMove {
					currentDirection = currentDirection.rotateRight()
				}
				visitedPositons[newPosition] = append(visitedPositons[newPosition], currentDirection)
				newPosition, didMove, movedOutside = floor.move(newPosition, currentDirection)
			}
			require.Equal(t, input.expected, len(possibleBlocks))
		})
	}
}

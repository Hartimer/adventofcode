package day3

import (
	"adventofcode/helper"

	"github.com/pkg/errors"
)

type house struct {
	x int
	y int
}

func Solve1(inputFilePath string) (int, error) {
	currentPosition := house{x: 0, y: 0}
	visitedHouses := map[house]struct{}{
		currentPosition: {},
	}
	directions := <-helper.FileLineReader(inputFilePath)
	for _, direction := range directions {
		switch direction {
		case '<':
			currentPosition.x--
		case '>':
			currentPosition.x++
		case '^':
			currentPosition.y++
		case 'v':
			currentPosition.y--
		default:
			return 0, errors.New("Unknown direction " + string(direction))
		}
		visitedHouses[currentPosition] = struct{}{}
	}
	return len(visitedHouses), nil
}

func Solve2(inputFilePath string) (int, error) {
	currentSantaPosition := house{x: 0, y: 0}
	currentRobotPosition := currentSantaPosition
	visitedHouses := map[house]struct{}{
		currentSantaPosition: {},
	}
	directions := <-helper.FileLineReader(inputFilePath)
	for idx, direction := range directions {
		var newPosition house
		var err error
		if idx%2 == 0 {
			newPosition, err = move(currentSantaPosition, direction)
			currentSantaPosition = newPosition
		} else {
			newPosition, err = move(currentRobotPosition, direction)
			currentRobotPosition = newPosition
		}
		if err != nil {
			return 0, errors.Wrap(err, "")
		}
		visitedHouses[newPosition] = struct{}{}
	}
	return len(visitedHouses), nil
}

func move(currentPosition house, direction rune) (house, error) {
	switch direction {
	case '<':
		currentPosition.x--
	case '>':
		currentPosition.x++
	case '^':
		currentPosition.y++
	case 'v':
		currentPosition.y--
	default:
		return house{}, errors.New("Unknown direction " + string(direction))
	}
	return currentPosition, nil
}

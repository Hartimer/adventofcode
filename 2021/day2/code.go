package day2

import (
	"adventofcode/helper"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type position struct {
	horizontal int
	vertical   int
	aim        int
}

func (p *position) move(horizontalOffset int, verticalOffset int) {
	p.horizontal += horizontalOffset
	p.vertical += verticalOffset
}

func (p *position) newMove(horizontalOffset int, aimOffset int) {
	p.horizontal += horizontalOffset
	p.aim += aimOffset
	p.vertical += horizontalOffset * p.aim
}

func parseMove(move string) (int, int, error) {
	moveParts := strings.Split(move, " ")
	if len(moveParts) != 2 {
		return 0, 0, errors.New(fmt.Sprintf("%s produced %d parts instead of 2", move, len(moveParts)))
	}
	offset, err := strconv.ParseInt(moveParts[1], 10, 32)
	if err != nil {
		return 0, 0, errors.Wrap(err, "")
	}
	offsetInt := int(offset)
	switch moveParts[0] {
	case "forward":
		return offsetInt, 0, nil
	case "up":
		return 0, -offsetInt, nil
	case "down":
		return 0, offsetInt, nil
	default:
		return 0, 0, errors.New(fmt.Sprintf("Unknown move %s", moveParts[0]))
	}
}

func Solve1(inputFilePath string) (int, error) {
	currentPosition := &position{}

	for fileLine := range helper.FileLineReader(inputFilePath) {
		horizontalOffset, verticalOffset, err := parseMove(fileLine)
		if err != nil {
			return 0, errors.Wrap(err, "")
		}
		currentPosition.move(horizontalOffset, verticalOffset)
	}
	return currentPosition.horizontal * currentPosition.vertical, nil
}

func Solve2(inputFilePath string) (int, error) {
	currentPosition := &position{}
	for fileLine := range helper.FileLineReader(inputFilePath) {
		horizontalOffset, verticalOffset, err := parseMove(fileLine)
		if err != nil {
			return 0, errors.Wrap(err, "")
		}
		currentPosition.newMove(horizontalOffset, verticalOffset)
	}
	return currentPosition.horizontal * currentPosition.vertical, nil
}

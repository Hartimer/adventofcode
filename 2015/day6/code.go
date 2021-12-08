package day6

import (
	"adventofcode/helper"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type lightState int

const (
	On  lightState = 1
	Off lightState = 0
)

type lightSwitch func(lightAction, lightState) lightState

var switch1 = func(action lightAction, l lightState) lightState {
	switch action {
	case turnOn:
		return On
	case turnOff:
		return Off
	default:
		if l == On {
			return Off
		}
		return On
	}
}

var switch2 = func(action lightAction, l lightState) lightState {
	switch action {
	case turnOn:
		return l + 1
	case turnOff:
		if l == 0 {
			return l
		}
		return l - 1
	default:
		return l + 2
	}
}

type lightAction int

const (
	turnOn lightAction = iota
	turnOff
	toggle
)

type board [1000][1000]lightState

func (b *board) do(action lightAction, s lightSwitch, rawStart, rawEnd string) error {
	startX, startY, err := parseCoordinates(rawStart)
	if err != nil {
		return errors.Wrap(err, "")
	}
	endX, endY, err := parseCoordinates(rawEnd)
	if err != nil {
		return errors.Wrap(err, "")
	}

	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			b[x][y] = s(action, b[x][y])
		}
	}
	return nil
}

func parseCoordinates(rawCoordinate string) (int, int, error) {
	parts := strings.Split(rawCoordinate, ",")
	if len(parts) != 2 {
		return 0, 0, errors.New(fmt.Sprintf("Expeted %s to have 2 parts by has %d", rawCoordinate, len(parts)))
	}
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, errors.Wrap(err, "")
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, errors.Wrap(err, "")
	}
	return x, y, nil
}

func Solve1(inputFilePath string) (int, error) {
	b := board{}
	for rowIdx, row := range b {
		for colIdx := range row {
			b[rowIdx][colIdx] = Off
		}
	}
	for fileLine := range helper.FileLineReader(inputFilePath) {
		var action lightAction
		if strings.HasPrefix(fileLine, "turn on ") {
			action = turnOn
			fileLine = strings.TrimPrefix(fileLine, "turn on ")
		} else if strings.HasPrefix(fileLine, "turn off ") {
			action = turnOff
			fileLine = strings.TrimPrefix(fileLine, "turn off ")
		} else {
			action = toggle
			fileLine = strings.TrimPrefix(fileLine, "toggle ")
		}

		coordinateParts := strings.Split(fileLine, " through ")
		if len(coordinateParts) != 2 {
			return 0, errors.New(fmt.Sprintf("Expected %s to have 2 parts bu has %d", fileLine, len(coordinateParts)))
		}

		if err := b.do(action, switch1, coordinateParts[0], coordinateParts[1]); err != nil {
			return 0, errors.Wrap(err, "")
		}
	}

	count := 0
	for _, row := range b {
		for _, light := range row {
			if light == On {
				count++
			}
		}
	}
	return count, nil
}

func Solve2(inputFilePath string) (int, error) {
	b := board{}
	for rowIdx, row := range b {
		for colIdx := range row {
			b[rowIdx][colIdx] = Off
		}
	}
	for fileLine := range helper.FileLineReader(inputFilePath) {
		var action lightAction
		if strings.HasPrefix(fileLine, "turn on ") {
			action = turnOn
			fileLine = strings.TrimPrefix(fileLine, "turn on ")
		} else if strings.HasPrefix(fileLine, "turn off ") {
			action = turnOff
			fileLine = strings.TrimPrefix(fileLine, "turn off ")
		} else {
			action = toggle
			fileLine = strings.TrimPrefix(fileLine, "toggle ")
		}

		coordinateParts := strings.Split(fileLine, " through ")
		if len(coordinateParts) != 2 {
			return 0, errors.New(fmt.Sprintf("Expected %s to have 2 parts bu has %d", fileLine, len(coordinateParts)))
		}

		if err := b.do(action, switch2, coordinateParts[0], coordinateParts[1]); err != nil {
			return 0, errors.Wrap(err, "")
		}
	}

	count := 0
	for _, row := range b {
		for _, light := range row {
			count += int(light)
		}
	}
	return count, nil
}

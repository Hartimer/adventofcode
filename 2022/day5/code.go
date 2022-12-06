package day5

import (
	"adventofcode/helper"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type stack []byte

type ship []stack

func parseContainer(s ship, line string) (ship, error) {
	for stackID := 0; len(line) > 0; stackID++ {
		containerLetter := line[1]
		if len(line) < 4 {
			line = ""
		} else {
			line = line[4:]
		}
		if len(s)-1 < stackID {
			s = append(s, stack{})
		}
		if containerLetter != ' ' {
			if _, err := strconv.Atoi(string(containerLetter)); err == nil {
				continue
			}
			s[stackID] = append([]byte{containerLetter}, s[stackID]...)
		}
	}
	return s, nil
}

func crateMover9000(s ship, sourceStackID, targetStackID, containerCount int) ship {
	for ; containerCount > 0; containerCount-- {
		s[targetStackID] = append(s[targetStackID], s[sourceStackID][len(s[sourceStackID])-1])
		s[sourceStackID] = s[sourceStackID][:len(s[sourceStackID])-1]
	}
	return s
}

func crateMover9001(s ship, sourceStackID, targetStackID, containerCount int) ship {
	s[targetStackID] = append(s[targetStackID], s[sourceStackID][len(s[sourceStackID])-containerCount:]...)
	s[sourceStackID] = s[sourceStackID][:len(s[sourceStackID])-containerCount]
	return s
}

func parseLine(s ship, line string, moveMode bool, crane func(ship, int, int, int) ship) (ship, error) {
	if moveMode {
		moveParts := strings.Split(line, " ")
		sourceStackID, err := strconv.Atoi(moveParts[3])
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		sourceStackID--
		targetStackID, err := strconv.Atoi(moveParts[5])
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		targetStackID--
		containerCount, err := strconv.Atoi(moveParts[1])
		if err != nil {
			return nil, errors.Wrap(err, "")
		}

		return crane(s, sourceStackID, targetStackID, containerCount), nil
	}
	return parseContainer(s, line)
}

func solve(inputFilePath string, crane func(ship, int, int, int) ship) (string, error) {
	var s ship
	moveMode := false
	for fileLine := range helper.FileLineReader(inputFilePath) {
		if len(fileLine) == 0 {
			moveMode = true
			continue
		}
		var err error
		s, err = parseLine(s, fileLine, moveMode, crane)
		if err != nil {
			return "", errors.Wrap(err, "")
		}
	}
	result := make([]byte, len(s))
	for stackID, stack := range s {
		result[stackID] = stack[len(stack)-1]
	}
	return string(result), nil
}

func Solve1(inputFilePath string) (string, error) {
	return solve(inputFilePath, crateMover9000)
}

func Solve2(inputFilePath string) (string, error) {
	return solve(inputFilePath, crateMover9001)
}

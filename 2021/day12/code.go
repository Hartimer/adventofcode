package day12

import (
	"adventofcode/helper"
	"unicode"

	"github.com/pkg/errors"
)

type cave struct {
	name        string
	connections map[string]cave
}

func parseInputs(inputFilePath string) (map[string]cave, error) {
	caves := map[string]cave{}
	for fileLine := range helper.FileLineReader(inputFilePath) {
		parts, err := helper.SplitAndCheck(fileLine, "-", 2)
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		from, exists := caves[parts[0]]
		if !exists {
			from = cave{name: parts[0], connections: map[string]cave{}}
			caves[parts[0]] = from
		}
		to, exists := caves[parts[1]]
		if !exists {
			to = cave{name: parts[1], connections: map[string]cave{}}
			caves[parts[1]] = to
		}

		from.connections[to.name] = to
		to.connections[from.name] = from
	}
	return caves, nil
}

func navigate(current cave, visitedCaves map[string]struct{}) int {
	if current.name == "end" {
		return 1
	}

	if unicode.IsLower(rune(current.name[0])) {
		if _, alreadyVisited := visitedCaves[current.name]; alreadyVisited {
			return 0
		}
	}
	newVisitedCaves := clone(visitedCaves)
	newVisitedCaves[current.name] = struct{}{}

	validCount := 0
	for name := range current.connections {
		validCount += navigate(current.connections[name], newVisitedCaves)
	}
	return validCount
}

func navigate2(current cave, visitedCaves map[string]struct{}, hasVisitedSmallTwice bool) int {
	if current.name == "end" {
		return 1
	}

	if unicode.IsLower(rune(current.name[0])) {
		if _, alreadyVisited := visitedCaves[current.name]; alreadyVisited {
			if hasVisitedSmallTwice || current.name == "start" {
				return 0
			}
			hasVisitedSmallTwice = true
		}
	}
	newVisitedCaves := clone(visitedCaves)
	newVisitedCaves[current.name] = struct{}{}

	validCount := 0
	for name := range current.connections {
		validCount += navigate2(current.connections[name], newVisitedCaves, hasVisitedSmallTwice)
	}
	return validCount
}

func clone(m1 map[string]struct{}) map[string]struct{} {
	m2 := map[string]struct{}{}
	for k := range m1 {
		m2[k] = m1[k]
	}
	return m2
}

func Solve1(inputFilePath string) (int, error) {
	caves, err := parseInputs(inputFilePath)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	return navigate(caves["start"], map[string]struct{}{}), nil
}

func Solve2(inputFilePath string) (int, error) {
	caves, err := parseInputs(inputFilePath)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	return navigate2(caves["start"], map[string]struct{}{}, false), nil
}

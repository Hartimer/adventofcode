package day9

import (
	"adventofcode/helper"
	"math"
	"strconv"

	"github.com/pkg/errors"
)

type path struct {
	destiny  string
	distance int
}

type mapz map[string][]path

func (m mapz) add(from, to string, distance int) {
	if _, exists := m[from]; !exists {
		m[from] = []path{}
	}
	m[from] = append(m[from], path{destiny: to, distance: distance})
	if _, exists := m[to]; !exists {
		m[to] = []path{}
	}
	m[to] = append(m[to], path{destiny: from, distance: distance})
}

func parsePath(raw string) (string, string, int, error) {
	pathParts, err := helper.SplitAndCheck(raw, " = ", 2)
	if err != nil {
		return "", "", 0, errors.Wrap(err, "")
	}
	distance, err := strconv.Atoi(pathParts[1])
	if err != nil {
		return "", "", 0, errors.Wrap(err, "")
	}

	locationParts, err := helper.SplitAndCheck(pathParts[0], " to ", 2)
	if err != nil {
		return "", "", 0, errors.Wrap(err, "")
	}

	return locationParts[0], locationParts[1], distance, nil
}
func Solve1(inputFilePath string) (int, error) {
	return solve(inputFilePath, comparator{f: lt, initialCounter: math.MaxInt})
}

func Solve2(inputFilePath string) (int, error) {
	return solve(inputFilePath, comparator{f: gt, initialCounter: math.MinInt})
}

func solve(inputFilePath string, comp comparator) (int, error) {
	m := mapz{}
	for fileLine := range helper.FileLineReader(inputFilePath) {
		from, to, distance, err := parsePath(fileLine)
		if err != nil {
			return 0, errors.Wrap(err, "")
		}
		m.add(from, to, distance)
	}
	smallestPath := comp.initialCounter
	for startLocation := range m {
		thisCost, valid := cost(m, map[string]struct{}{}, startLocation, 0, comp)
		if !valid {
			continue
		}
		if comp.f(thisCost, smallestPath) {
			smallestPath = thisCost
		}
	}
	return smallestPath, nil
}

func cost(m mapz, visited map[string]struct{}, currentLocation string, currentCost int, comp comparator) (int, bool) {
	if _, alreadySeen := visited[currentLocation]; alreadySeen {
		return 0, false
	}
	visited[currentLocation] = struct{}{}
	if len(m) == len(visited) {
		return currentCost, true
	}
	smallestPath := comp.initialCounter
	for _, nextLocation := range m[currentLocation] {
		visitedClone := cloneMap(visited)
		pathCost, valid := cost(m, visitedClone, nextLocation.destiny, currentCost+nextLocation.distance, comp)
		if !valid {
			continue
		}
		if comp.f(pathCost, smallestPath) {
			smallestPath = pathCost
		}
	}
	return smallestPath, true
}

func cloneMap(s map[string]struct{}) map[string]struct{} {
	s2 := map[string]struct{}{}
	for k := range s {
		s2[k] = s[k]
	}
	return s2
}

type comparator struct {
	f              func(int, int) bool
	initialCounter int
}

var gt = func(a, b int) bool {
	return a > b
}

var lt = func(a, b int) bool {
	return a < b
}

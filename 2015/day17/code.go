package day17

import (
	"adventofcode/helper"
	"math"
	"strconv"

	"github.com/pkg/errors"
)

type container struct {
	id       int
	capacity int
}

func parseContainers(inputFilePath string) ([]container, error) {
	var containers []container

	id := 0
	for fileLine := range helper.FileLineReader(inputFilePath) {
		capacity, err := strconv.Atoi(fileLine)
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		containers = append(containers, container{id: id, capacity: capacity})
		id++
	}

	return containers, nil
}

func diffentCombos(currentFill, targetFill int, remainingContainers []container) [][]int {
	combos := [][]int{}
	for idx, nextContainer := range remainingContainers {
		nextFill := currentFill + nextContainer.capacity
		if nextFill == targetFill {
			combos = append(combos, []int{nextContainer.capacity})
		} else if nextFill > targetFill {
			continue
		} else {
			nextCombos := diffentCombos(nextFill, targetFill, remainingContainers[idx+1:])
			for idx := range nextCombos {
				nextCombos[idx] = append(nextCombos[idx], nextContainer.capacity)
			}
			combos = append(combos, nextCombos...)
		}
	}
	return combos
}

func Solve1(inputFilePath string, target int) (int, error) {
	containers, err := parseContainers(inputFilePath)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	return len(diffentCombos(0, target, containers)), nil
}

func Solve2(inputFilePath string, target int) (int, error) {
	containers, err := parseContainers(inputFilePath)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	combos := diffentCombos(0, target, containers)
	minSize := math.MaxInt
	minCount := 0
	for _, c := range combos {
		if len(c) < minSize {
			minSize = len(c)
			minCount = 1
		} else if len(c) == minSize {
			minCount++
		}
	}
	return minCount, nil
}

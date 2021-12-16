package day15

import (
	"adventofcode/helper"
	"math"
	"strconv"

	"github.com/pkg/errors"
)

type coordinate struct {
	x int
	y int
}

type node struct {
	c    coordinate
	best int
}

type caveMap [][]int

func (c caveMap) valid(coord coordinate) bool {
	return !(coord.x < 0 || coord.x >= len(c) || coord.y < 0 || coord.y >= len(c[coord.x]))
}

func (c caveMap) get(coord coordinate) (int, bool) {
	if !c.valid(coord) {
		return 0, false
	}
	return c[coord.x][coord.y], true
}

func (c caveMap) neigbhors(coord coordinate) []coordinate {
	result := []coordinate{
		{x: coord.x + 1, y: coord.y},
		{x: coord.x, y: coord.y + 1},
		{x: coord.x - 1, y: coord.y},
		{x: coord.x, y: coord.y - 1},
	}
	var finalResult []coordinate
	for idx := range result {
		coord := result[idx]
		if c.valid(coord) {
			finalResult = append(finalResult, coord)
		}
	}
	return finalResult
}

func (c caveMap) expand() caveMap {
	expansionFactor := 5
	c2 := make(caveMap, len(c)*5)
	for rowExpansion := 0; rowExpansion < expansionFactor; rowExpansion++ {
		for x, row := range c {
			newRow := make([]int, len(row)*5)
			for y, cell := range row {
				for colExpansion := 0; colExpansion < expansionFactor; colExpansion++ {
					newRow[y+colExpansion*len(row)] = safeIncr(cell, rowExpansion+colExpansion)
				}
			}
			c2[x+rowExpansion*len(c)] = newRow
		}
	}
	return c2
}

func safeIncr(n, p int) int {
	for i := 0; i < p; i++ {
		if n == 9 {
			n = 1
		} else {
			n++
		}
	}
	return n
}

func parseInputs(inputFilePath string) (caveMap, error) {
	cave := caveMap{}
	for fileLine := range helper.FileLineReader(inputFilePath) {
		newRow := make([]int, len(fileLine))
		var err error
		for idx, c := range fileLine {
			newRow[idx], err = strconv.Atoi((string(c)))
			if err != nil {
				return nil, errors.Wrap(err, "")
			}
		}
		cave = append(cave, newRow)
	}
	return cave, nil
}

func navigate(m caveMap) int {
	start := coordinate{x: 0, y: 0}
	unsettled := map[coordinate]*node{start: {c: start, best: 0}}
	settled := map[coordinate]*node{}

	for len(unsettled) > 0 {
		currentNode := lowestDistanceNode(unsettled)
		delete(unsettled, currentNode.c)
		for _, neighbor := range m.neigbhors(currentNode.c) {
			if _, alreadySettled := settled[neighbor]; alreadySettled {
				continue
			}

			neighborNode, exists := unsettled[neighbor]
			if !exists {
				neighborNode = &node{c: neighbor, best: math.MaxInt}
			}
			calculateMinimumDistance(currentNode, neighborNode, m)
			unsettled[neighbor] = neighborNode
		}
		settled[currentNode.c] = currentNode
	}
	target := settled[coordinate{x: len(m) - 1, y: len(m[0]) - 1}]
	return target.best
}

func lowestDistanceNode(unsettled map[coordinate]*node) *node {
	var result *node
	lowestDistance := math.MaxInt
	for c := range unsettled {
		n := unsettled[c]
		if n.best < lowestDistance {
			lowestDistance = n.best
			result = n
		}
	}
	return result
}

func calculateMinimumDistance(from, to *node, m caveMap) {
	weight, _ := m.get(to.c)
	if from.best+weight < to.best {
		to.best = from.best + weight
	}
}

func Solve1(inputFilePath string) (int, error) {
	inputs, err := parseInputs(inputFilePath)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	return navigate(inputs), nil
}

func Solve2(inputFilePath string) (int, error) {
	inputs, err := parseInputs(inputFilePath)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	return navigate(inputs.expand()), nil
}

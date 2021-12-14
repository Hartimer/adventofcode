package day14

import (
	"adventofcode/helper"
	"math"

	"github.com/pkg/errors"
)

type pair struct {
	p string
	l string
}

func parseInputs(inputFilePath string) (string, map[string]string, error) {
	var formula string
	pairs := map[string]string{}
	parsingFormula := true
	for fileLine := range helper.FileLineReader(inputFilePath) {
		if len(fileLine) == 0 {
			parsingFormula = false
		} else if parsingFormula {
			formula = fileLine
		} else {
			parts, err := helper.SplitAndCheck(fileLine, " -> ", 2)
			if err != nil {
				return "", nil, errors.Wrap(err, "")
			}
			pairs[parts[0]] = parts[1]
		}

	}
	return formula, pairs, nil
}

func solve(formula []byte, pairs map[string]string, targetIterations int) int {
	var result []byte
	for iterations := 0; iterations < targetIterations; iterations, formula, result = iterations+1, result, make([]byte, 0, len(result)*2) {
		for i := 0; i < len(formula)-1; i++ {
			pair := formula[i : i+2]
			if letter, exists := pairs[string(pair)]; exists {
				if len(result) == 0 || result[len(result)-1] != pair[0] {
					result = append(result, pair[0], letter[0], pair[1])
				} else {
					result = append(result, letter[0], pair[1])
				}
			}
		}
	}

	stats := map[byte]int{}
	for _, r := range formula {
		if _, exists := stats[r]; !exists {
			stats[r] = 0
		}
		stats[r]++
	}

	max, min := math.MinInt, math.MaxInt
	for _, count := range stats {
		if count > max {
			max = count
		}
		if count < min {
			min = count
		}
	}
	return max - min
}

func Solve1(inputFilePath string) (int, error) {
	inputs, pairs, err := parseInputs(inputFilePath)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	return solve([]byte(inputs), pairs, 10), nil
}

func Solve2(inputFilePath string) (int, error) {
	inputs, pairs, err := parseInputs(inputFilePath)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	return solve([]byte(inputs), pairs, 40), nil
}

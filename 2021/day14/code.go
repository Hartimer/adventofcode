package day14

import (
	"adventofcode/helper"
	"log"
	"math"

	"github.com/pkg/errors"
)

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

func solve(formula []byte, rules map[string]string, targetIterations int) int {
	charCount := map[byte]int{}
	pairCount := map[string]int{}
	for i := 0; i < len(formula)-1; i++ {
		pair := formula[i : i+2]
		if _, exists := pairCount[string(string(pair))]; !exists {
			pairCount[string(string(pair))] = 0
		}
		pairCount[string(pair)]++
		if _, exists := charCount[formula[i]]; !exists {
			charCount[formula[i]] = 0
		}
		charCount[formula[i]]++
	}
	if _, exists := charCount[formula[len(formula)-1]]; !exists {
		charCount[formula[len(formula)-1]] = 1
	} else {
		charCount[formula[len(formula)-1]]++
	}

	for iterations := 0; iterations < targetIterations; iterations++ {
		newPairCount := map[string]int{}
		for pair, originalPairCount := range pairCount {
			letter, exists := rules[string(pair)]
			if !exists {
				continue
			}
			leftPair := append([]byte(pair[:1]), letter[0])
			rightPair := append([]byte{letter[0]}, []byte(pair[1:])...)

			newPairCount[string(leftPair)] += originalPairCount
			newPairCount[string(rightPair)] += originalPairCount
			if _, exists := charCount[letter[0]]; !exists {
				charCount[letter[0]] = 0
			}
			charCount[letter[0]] += originalPairCount
		}
		pairCount = newPairCount
	}

	max, min := math.MinInt, math.MaxInt
	for c, count := range charCount {
		log.Printf("%s: %d", string(c), count)
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

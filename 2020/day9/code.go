package day9

import (
	"adventofcode/helper"
	"math"
	"strconv"

	"github.com/pkg/errors"
)

func parseNumbers(inputFilePath string) ([]int, error) {
	var numbers []int

	for fileLine := range helper.FileLineReader(inputFilePath) {
		n, err := strconv.Atoi(fileLine)
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		numbers = append(numbers, n)
	}
	return numbers, nil
}

func Solve1(inputFilePath string, preambleSize int) (int, error) {
	numbers, err := parseNumbers(inputFilePath)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	for len(numbers) > preambleSize+1 {
		if !isValidNumber(numbers[:preambleSize], numbers[preambleSize]) {
			return numbers[preambleSize], nil
		}
		numbers = numbers[1:]
	}
	return 0, errors.New("No invalid number found")
}

func Solve2(inputFilePath string, invalidNumber int) (int, error) {
	numbers, err := parseNumbers(inputFilePath)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	for idx := range numbers {
		sumElements, success := recursiveSum(numbers[idx:], invalidNumber, 0)
		if !success {
			continue
		}
		return min(sumElements) + max(sumElements), nil
	}
	return 0, errors.New("No good vulnerability found")
}

func recursiveSum(numbers []int, target int, acc int) ([]int, bool) {
	if len(numbers) == 0 {
		return nil, false
	}
	n := numbers[0]
	sum := acc + n
	if sum == target {
		return []int{n}, true
	}
	if sum > target {
		return nil, false
	}
	if furtherSum, success := recursiveSum(numbers[1:], target, sum); success {
		return append([]int{n}, furtherSum...), true
	}
	return nil, false
}

func min(n []int) int {
	i := math.MaxInt32
	for _, newI := range n {
		if newI < i {
			i = newI
		}
	}
	return i
}

func max(n []int) int {
	i := 0
	for _, newI := range n {
		if newI > i {
			i = newI
		}
	}
	return i
}

func isValidNumber(preamble []int, number int) bool {
	for i := 0; i < len(preamble); i++ {
		for j := i + 1; j < len(preamble); j++ {
			if preamble[i]+preamble[j] == number {
				return true
			}
		}
	}
	return false
}

package day7

import (
	"adventofcode/helper"
	"math"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func Solve1(inputFilePath string) (int, error) {
	return solve(inputFilePath, func(n float64) float64 {
		return n
	})
}

func Solve2(inputFilePath string) (int, error) {
	return solve(inputFilePath, func(n float64) float64 {
		var res float64 = 0
		for i := n; i > 0; i-- {
			res += i
		}
		return res
	})
}

func solve(inputFilePath string, fuelFunc func(float64) float64) (int, error) {
	input := <-helper.FileLineReader(inputFilePath)
	var hPositions []int
	minH := math.MaxInt
	maxH := -1

	for _, rawNumber := range strings.Split(input, ",") {
		n, err := strconv.Atoi(rawNumber)
		if err != nil {
			return 0, errors.Wrap(err, "")
		}
		hPositions = append(hPositions, n)
		minH = minInt(minH, n)
		maxH = maxInt(maxH, n)
	}
	minFuel := math.MaxFloat64
	for i := minH; i <= maxH; i++ {
		var localCost float64 = 0
		for _, n := range hPositions {
			localCost += fuelFunc(math.Abs(float64(n - i)))
		}
		minFuel = math.Min(minFuel, localCost)
	}
	return int(minFuel), nil
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

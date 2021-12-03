package day1

import (
	"adventofcode/helper"
	"log"
	"strconv"

	"github.com/pkg/errors"
)

func Solve1(inputFilePath string) (int, error) {
	increaseCounter := 0
	var previousMeasurement int64 = -1
	for fileLine := range helper.FileLineReader(inputFilePath) {
		currentMeasurement, err := strconv.ParseInt(fileLine, 10, 64)
		if err != nil {
			return 0, errors.Wrap(err, "")
		}
		if previousMeasurement > 0 && currentMeasurement > previousMeasurement {
			increaseCounter++
		}
		previousMeasurement = currentMeasurement
	}

	return increaseCounter, nil
}

const windowSize int = 3
const windowCount int = 3

func Solve2(inputFilePath string) (int, error) {
	var slidingWindows [][]int64
	var previousWindowMeasurement int64 = -1
	increaseCounter := 0

	for fileLine := range helper.FileLineReader(inputFilePath) {
		currentMeasurement, err := strconv.ParseInt(fileLine, 10, 64)
		if err != nil {
			return 0, errors.Wrap(err, "")
		}

		if len(slidingWindows) < windowCount {
			slidingWindows = append(slidingWindows, []int64{})
		}

		removeFirstElement := false

		// Add measurement to all active sliding windows
		for windowIdx := range slidingWindows {
			slidingWindows[windowIdx] = append(slidingWindows[windowIdx], currentMeasurement)
			if len(slidingWindows[windowIdx]) == windowSize {
				log.Printf("Complete window: %v", slidingWindows[windowIdx])
				currentWindowMeasurement := sumSlice(slidingWindows[windowIdx])
				if previousWindowMeasurement > 0 && currentWindowMeasurement > previousWindowMeasurement {
					increaseCounter++
				}
				previousWindowMeasurement = currentWindowMeasurement
				// Whenever a window reaches its max capacity, we can remove it
				removeFirstElement = true
			}
		}

		if removeFirstElement {
			slidingWindows = slidingWindows[1:]
		}
	}

	return increaseCounter, nil
}

func sumSlice(slice []int64) int64 {
	var sum int64 = 0
	for _, val := range slice {
		sum += val
	}
	return sum
}

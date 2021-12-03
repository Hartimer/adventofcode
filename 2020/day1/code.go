package day1

import (
	"adventofcode/helper"
	"strconv"

	"github.com/pkg/errors"
)

const targetSum = 2020

func parseCoins(inputFilePath string) ([]int64, error) {
	var coinValues []int64

	for fileLine := range helper.FileLineReader(inputFilePath) {
		currentCoin, err := strconv.ParseInt(fileLine, 10, 64)
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		if currentCoin > targetSum {
			continue
		}

		coinValues = append(coinValues, currentCoin)
	}
	return coinValues, nil
}

func Solve1(inputFilePath string) (int, error) {
	coinValues, err := parseCoins(inputFilePath)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	for i := 0; i < len(coinValues); i++ {
		firstCoin := coinValues[i]
		for j := i + 1; j < len(coinValues); j++ {
			secondCoin := coinValues[j]
			if firstCoin+secondCoin == targetSum {
				return int(firstCoin * secondCoin), nil
			}
		}
	}

	return 0, errors.New("No pair found")
}

func Solve2(inputFilePath string) (int, error) {
	coinValues, err := parseCoins(inputFilePath)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	for i := 0; i < len(coinValues); i++ {
		firstCoin := coinValues[i]
		for j := i + 1; j < len(coinValues); j++ {
			secondCoin := coinValues[j]
			if firstCoin+secondCoin >= targetSum {
				continue
			}
			for k := j + 1; k < len(coinValues); k++ {
				thirdCoin := coinValues[k]
				if firstCoin+secondCoin+thirdCoin == targetSum {
					return int(firstCoin * secondCoin * thirdCoin), nil
				}
			}
		}
	}

	return 0, errors.New("No pair found")
}

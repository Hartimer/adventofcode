package day6

import (
	"adventofcode/helper"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func solveImproved(inputFilePath string, days int) (int, error) {
	inputs := <-helper.FileLineReader(inputFilePath)

	timers := make([]int, 9)
	for _, rawFish := range strings.Split(inputs, ",") {
		timer, err := strconv.Atoi(rawFish)
		if err != nil {
			return 0, errors.Wrap(err, "")
		}
		timers[timer]++
	}

	for i := 0; i < days; i++ {
		breedingFish := timers[0]
		timers = append(timers[1:], breedingFish)
		timers[6] += breedingFish
	}

	count := 0
	for _, c := range timers {
		count += c
	}
	return count, nil
}

func solve(inputFilePath string, days int) (int, error) {
	inputs := <-helper.FileLineReader(inputFilePath)

	timers := map[int]int{}
	for _, rawFish := range strings.Split(inputs, ",") {
		timer, err := strconv.Atoi(rawFish)
		if err != nil {
			return 0, errors.Wrap(err, "")
		}
		if _, exists := timers[timer]; !exists {
			timers[timer] = 0
		}
		timers[timer]++
	}

	for i := 0; i < days; i++ {
		newTimers := map[int]int{}
		for timer, count := range timers {
			if timer == 0 {
				newTimers[6] = count
				newTimers[8] = count
				if sevens, exist := timers[7]; exist {
					newTimers[6] += sevens
					delete(timers, 7)
				}
			} else {
				newTimers[timer-1] = count
			}
		}
		timers = newTimers
	}

	count := 0
	for _, c := range timers {
		count += c
	}
	return count, nil
}

func Solve1(inputFilePath string) (int, error) {
	return solveImproved(inputFilePath, 80)
}

func Solve2(inputFilePath string) (int, error) {
	return solveImproved(inputFilePath, 256)
}

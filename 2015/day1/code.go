package day1

import (
	"adventofcode/helper"

	"github.com/pkg/errors"
)

func Solve1(inputFilePath string) (int, error) {
	input := <-helper.FileLineReader(inputFilePath)
	level := 0
	for _, instruction := range input {
		if instruction == '(' {
			level++
		} else {
			level--
		}
	}
	return level, nil
}

func Solve2(inputFilePath string) (int, error) {
	input := <-helper.FileLineReader(inputFilePath)
	level := 0
	for idx, instruction := range input {
		if instruction == '(' {
			level++
		} else {
			level--
		}
		if level < 0 {
			return idx + 1, nil
		}
	}
	return 0, errors.New("Never entered basement")
}

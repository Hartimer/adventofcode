package day2_test

import (
	"adventofcode/2015/day2"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	inputs := map[string]int{
		"./input_test.txt": 101,
		"./input.txt":      1598415,
	}
	testhelper.Runner(t, day2.Solve1, inputs)
}

func TestSolve2(t *testing.T) {
	inputs := map[string]int{
		"./input_test.txt": 48,
		"./input.txt":      3812909,
	}
	testhelper.Runner(t, day2.Solve2, inputs)
}

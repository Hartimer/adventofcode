package day7_test

import (
	"adventofcode/2015/day7"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	inputs := map[string]int{
		"./input_test.txt": 114,
		"./input.txt":      3176,
	}
	testhelper.Runner(t, day7.Solve1, inputs)
}

func TestSolve2(t *testing.T) {
	inputs := map[string]int{
		"./input.txt": 14710,
	}
	testhelper.Runner(t, day7.Solve2, inputs)
}

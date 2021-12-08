package day5_test

import (
	"adventofcode/2015/day5"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	inputs := map[string]int{
		"./input_test.txt": 2,
		"./input.txt":      255,
	}
	testhelper.Runner(t, day5.Solve1, inputs)
}

func TestSolve2(t *testing.T) {
	inputs := map[string]int{
		"./input_test2.txt": 3,
		"./input.txt":       55,
	}
	testhelper.Runner(t, day5.Solve2, inputs)
}

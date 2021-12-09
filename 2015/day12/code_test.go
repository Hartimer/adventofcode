package day12_test

import (
	"adventofcode/2015/day12"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	inputs := map[string]int{
		"./input_test.txt": 18,
		"./input.txt":      111754,
	}
	testhelper.Runner(t, day12.Solve1, inputs)
}

func TestSolve2(t *testing.T) {
	inputs := map[string]int{
		"./input_test.txt": 18,
		"./input.txt":      65402,
	}
	testhelper.Runner(t, day12.Solve2, inputs)
}

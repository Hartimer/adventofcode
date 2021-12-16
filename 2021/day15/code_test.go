package day15_test

import (
	"adventofcode/2021/day15"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 40,
		"./input.txt":      441,
	}
	testhelper.Runner(t, day15.Solve1, input)
}

func TestSolve2(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 315,
		"./input.txt":      2849,
	}
	testhelper.Runner(t, day15.Solve2, input)
}

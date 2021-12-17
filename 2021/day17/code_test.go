package day17_test

import (
	"adventofcode/2021/day17"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 45,
		"./input.txt":      5253,
	}
	testhelper.Runner(t, day17.Solve1, input)
}

func TestSolve2(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 112,
		"./input.txt":      1770,
	}
	testhelper.Runner(t, day17.Solve2, input)
}

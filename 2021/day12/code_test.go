package day12_test

import (
	"adventofcode/2021/day12"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 10,
		"./input.txt":      5756,
	}
	testhelper.Runner(t, day12.Solve1, input)
}

func TestSolve2(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 36,
		"./input.txt":      144603,
	}
	testhelper.Runner(t, day12.Solve2, input)
}

package day14_test

import (
	"adventofcode/2021/day14"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 1588,
		"./input.txt":      2703,
	}
	testhelper.Runner(t, day14.Solve1, input)
}

func TestSolve2(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 2188189693529,
		"./input.txt":      2984946368465,
	}
	testhelper.Runner(t, day14.Solve2, input)
}

package day3_test

import (
	"adventofcode/2020/day3"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 7,
		"./input.txt":      247,
	}
	testhelper.Runner(t, day3.Solve1, input)
}

func TestSolve2(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 336,
		"./input.txt":      2983070376,
	}
	testhelper.Runner(t, day3.Solve2, input)
}

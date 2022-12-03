package day3_test

import (
	"adventofcode/2022/day3"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 157,
		"./input.txt":      7674,
	}
	testhelper.Runner(t, day3.Solve1, input)
}

func TestSolve2(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 70,
		"./input.txt":      0,
	}
	testhelper.Runner(t, day3.Solve2, input)
}

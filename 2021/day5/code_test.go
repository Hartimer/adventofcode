package day5_test

import (
	"adventofcode/2021/day5"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 5,
		"./input.txt":      8111,
	}
	testhelper.Runner(t, day5.Solve1, input)
}

func TestSolve2(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 12,
		"./input.txt":      22088,
	}
	testhelper.Runner(t, day5.Solve2, input)
}

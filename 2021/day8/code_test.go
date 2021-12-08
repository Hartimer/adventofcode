package day8_test

import (
	"adventofcode/2021/day8"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 26,
		"./input.txt":      0,
	}
	testhelper.Runner(t, day8.Solve1, input)
}

func TestSolve2(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 61229,
		"./input.txt":      1091609,
	}
	testhelper.Runner(t, day8.Solve2, input)
}

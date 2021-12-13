package day13_test

import (
	"adventofcode/2021/day13"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 17,
		"./input.txt":      689,
	}
	testhelper.Runner(t, day13.Solve1, input)
}

func TestSolve2(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 16,
		"./input.txt":      91,
	}
	testhelper.Runner(t, day13.Solve2, input)
}

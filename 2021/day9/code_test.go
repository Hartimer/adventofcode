package day9_test

import (
	"adventofcode/2021/day9"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 15,
		"./input.txt":      452,
	}
	testhelper.Runner(t, day9.Solve1, input)
}

func TestSolve2(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 1134,
		"./input.txt":      1263735,
	}
	testhelper.Runner(t, day9.Solve2, input)
}

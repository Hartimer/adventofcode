package day10_test

import (
	"adventofcode/2021/day10"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 26397,
		"./input.txt":      469755,
	}
	testhelper.Runner(t, day10.Solve1, input)
}

func TestSolve2(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 288957,
		"./input.txt":      2762335572,
	}
	testhelper.Runner(t, day10.Solve2, input)
}

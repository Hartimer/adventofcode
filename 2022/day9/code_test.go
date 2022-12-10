package day9_test

import (
	"adventofcode/2022/day9"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 13,
		"./input.txt":      6026,
	}
	testhelper.Runner(t, day9.Solve1, input)
}

func TestSolve2(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt":  1,
		"./input_test2.txt": 36,
		"./input.txt":       2273, // 2527 is too high
	}
	testhelper.Runner(t, day9.Solve2, input)
}

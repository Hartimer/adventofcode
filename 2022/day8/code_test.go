package day8_test

import (
	"adventofcode/2022/day8"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 21,
		"./input.txt":      1787, // 9392 is too high
	}
	testhelper.Runner(t, day8.Solve1, input)
}

func TestSolve2(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 8,
		"./input.txt":      440640,
	}
	testhelper.Runner(t, day8.Solve2, input)
}

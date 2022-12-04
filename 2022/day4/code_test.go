package day4_test

import (
	"adventofcode/2022/day4"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 2,
		"./input.txt":      483,
	}
	testhelper.Runner(t, day4.Solve1, input)
}

func TestSolve2(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 4,
		"./input.txt":      874,
	}
	testhelper.Runner(t, day4.Solve2, input)
}

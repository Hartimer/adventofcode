package day1_test

import (
	"adventofcode/2021/day1"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 7,
		"./input.txt":      1502,
	}
	testhelper.Runner(t, day1.Solve1, input)
}

func TestSolve2(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 5,
		"./input.txt":      1538,
	}
	testhelper.Runner(t, day1.Solve2, input)
}

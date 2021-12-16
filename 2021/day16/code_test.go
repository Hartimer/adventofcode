package day16_test

import (
	"adventofcode/2021/day16"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 16,
		"./input.txt":      977,
	}
	testhelper.Runner(t, day16.Solve1, input)
}

func TestSolve2(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 15,
		"./input.txt":      101501020883,
	}
	testhelper.Runner(t, day16.Solve2, input)
}

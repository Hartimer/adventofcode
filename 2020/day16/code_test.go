package day16_test

import (
	"adventofcode/2020/day16"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 71,
		"./input.txt":      0,
	}
	testhelper.Runner(t, day16.Solve1, input)
}

func TestSolve2(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test2.txt": 0,
		// "./input.txt": 0,
	}
	testhelper.Runner(t, day16.Solve2, input)
}

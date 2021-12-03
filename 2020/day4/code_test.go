package day4_test

import (
	"adventofcode/2020/day4"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 2,
		"./input.txt":      250,
	}
	testhelper.Runner(t, day4.Solve1, input)
}

func TestSolve2(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test2.txt": 4,
		// "./input.txt":      0,
	}
	testhelper.Runner(t, day4.Solve2, input)
}

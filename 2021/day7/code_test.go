package day7_test

import (
	"adventofcode/2021/day7"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 37,
		"./input.txt":      333755,
	}
	testhelper.Runner(t, day7.Solve1, input)
}

func TestSolve2(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 168,
		"./input.txt":      94017638,
	}
	testhelper.Runner(t, day7.Solve2, input)
}

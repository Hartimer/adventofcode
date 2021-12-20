package day19_test

import (
	"adventofcode/2021/day19"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 0,
		// "./input.txt":      0,
	}
	testhelper.Runner(t, day19.Solve1, input)
}

func TestSolve2(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 0,
		// "./input.txt":      0,
	}
	testhelper.Runner(t, day19.Solve2, input)
}

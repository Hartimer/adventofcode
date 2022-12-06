package day5_test

import (
	"adventofcode/2022/day5"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	input := testhelper.InputsString{
		"./input_test.txt": "CMZ",
		"./input.txt":      "TQRFCBSJJ",
	}
	testhelper.RunnerString(t, day5.Solve1, input)
}

func TestSolve2(t *testing.T) {
	input := testhelper.InputsString{
		"./input_test.txt": "MCD",
		"./input.txt":      "RMHFJNVFP",
	}
	testhelper.RunnerString(t, day5.Solve2, input)
}

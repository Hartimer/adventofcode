package day19_test

import (
	"adventofcode/2020/day19"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 2,
		"./input.txt":      168,
		"./input2.txt":     0,
	}
	testhelper.Runner(t, day19.Solve1, input)
}

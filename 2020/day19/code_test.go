package day19_test

import (
	"adventofcode/2020/day19"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	input := testhelper.Inputs{
		// "./input_test.txt": 2,
		// "./input.txt":      168,
		"./input_test3.txt": 1,
		// "./input_test4.txt": 1,
	}
	testhelper.Runner(t, day19.Solve1, input)
}

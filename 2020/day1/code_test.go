package day1_test

import (
	"adventofcode/2020/day1"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 514579,
		"./input.txt":      1007331,
	}
	testhelper.Runner(t, day1.Solve1, input)
}

func TestSolve2(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 241861950,
		"./input.txt":      48914340,
	}
	testhelper.Runner(t, day1.Solve2, input)
}

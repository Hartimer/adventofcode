package day10_test

import (
	"adventofcode/2022/day10"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 13140,
		"./input.txt":      13680,
	}
	testhelper.Runner(t, day10.Solve1, input)
}

func TestSolve2(t *testing.T) {
	t.Skip("The answer is written to the console")
	input := testhelper.Inputs{
		"./input_test.txt": 1,
		"./input.txt":      1, // answer is PZGPKPEB
	}
	testhelper.Runner(t, day10.Solve2, input)
}

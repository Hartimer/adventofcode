package day22_test

import (
	"adventofcode/2021/day22"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt":  590784,
		"./input.txt":       642125,
		"./input_test2.txt": 474140,
	}
	testhelper.Runner(t, day22.Solve1, input)
}

func TestSolve2(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test2.txt": 2758514936282235,
	}
	testhelper.Runner(t, day22.Solve2, input)
}

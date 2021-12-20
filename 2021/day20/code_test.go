package day20_test

import (
	"adventofcode/2021/day20"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 35,
		"./input.txt":      5275,
	}
	testhelper.Runner(t, day20.Solve1, input)
}

func TestSolve2(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 3351,
		"./input.txt":      16482, // 17222 is too high, 3638 is too low
	}
	testhelper.Runner(t, day20.Solve2, input)
}

package day11_test

import (
	"adventofcode/2021/day11"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 1656,
		"./input.txt":      1640,
	}
	testhelper.Runner(t, day11.Solve1, input)
}

func TestSolve2(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 195,
		"./input.txt":      312,
	}
	testhelper.Runner(t, day11.Solve2, input)
}

package day2_test

import (
	"adventofcode/2020/day2"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 2,
		"./input.txt":      474,
	}
	testhelper.Runner(t, day2.Solve1, input)
}

func TestSolve2(t *testing.T) {
	input := map[string]int{
		"./input_test.txt": 1,
		"./input.txt":      745,
	}
	testhelper.Runner(t, day2.Solve2, input)
}

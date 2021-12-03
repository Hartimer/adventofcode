package day8_test

import (
	"adventofcode/2020/day8"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 5,
		"./input.txt":      1654,
	}
	testhelper.Runner(t, day8.Solve1, input)
}

func TestSolve2(t *testing.T) {
	input := map[string]int{
		"./input_test.txt": 8,
		"./input.txt":      0,
	}
	testhelper.Runner(t, day8.Solve1, input)
}

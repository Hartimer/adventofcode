package day1_test

import (
	"adventofcode/2015/day1"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	inputs := map[string]int{
		"./input_test.txt": -3,
		"./input.txt":      138,
	}
	testhelper.Runner(t, day1.Solve1, inputs)
}

func TestSolve2(t *testing.T) {
	inputs := map[string]int{
		"./input_test2.txt": 5,
		"./input.txt":       1771,
	}
	testhelper.Runner(t, day1.Solve2, inputs)
}

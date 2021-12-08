package day3_test

import (
	"adventofcode/2015/day3"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	inputs := map[string]int{
		"./input_test.txt": 4,
		"./input.txt":      2572,
	}
	testhelper.Runner(t, day3.Solve1, inputs)
}

func TestSolve2(t *testing.T) {
	inputs := map[string]int{
		"./input_test.txt": 3,
		"./input.txt":      2631,
	}
	testhelper.Runner(t, day3.Solve2, inputs)
}

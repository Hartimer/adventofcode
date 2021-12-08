package day6_test

import (
	"adventofcode/2015/day6"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	inputs := map[string]int{
		"./input_test.txt": 4,
		"./input.txt":      400410,
	}
	testhelper.Runner(t, day6.Solve1, inputs)
}

func TestSolve2(t *testing.T) {
	inputs := map[string]int{
		"./input_test.txt":  4,
		"./input_test2.txt": 8,
		"./input.txt":       15343601,
	}
	testhelper.Runner(t, day6.Solve2, inputs)
}

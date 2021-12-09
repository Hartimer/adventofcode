package day9_test

import (
	"adventofcode/2015/day9"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	inputs := map[string]int{
		"./input_test.txt": 605,
		"./input.txt":      207,
	}
	testhelper.Runner(t, day9.Solve1, inputs)
}

func TestSolve2(t *testing.T) {
	inputs := map[string]int{
		"./input_test.txt": 982,
		"./input.txt":      804,
	}
	testhelper.Runner(t, day9.Solve2, inputs)
}

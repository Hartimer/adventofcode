package day15_test

import (
	"adventofcode/2015/day15"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	inputs := map[string]int{
		"./input_test.txt": 62842880,
		// "./input.txt":      664,
	}
	testhelper.Runner(t, day15.Solve1, inputs)
}

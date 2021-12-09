package day13_test

import (
	"adventofcode/2015/day13"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	inputs := map[string]int{
		"./input_test.txt": 330,
		// "./input.txt":      0,
	}
	testhelper.Runner(t, day13.Solve1, inputs)
}

package day15_test

import (
	"adventofcode/2015/day15"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	inputs := map[string]int{
		"./input_test.txt": 62842880,
		"./input.txt":      18965440,
	}
	testhelper.Runner(t, day15.Solve1, inputs)
}

func TestSolve2(t *testing.T) {
	inputs := map[string]int{
		"./input_test.txt": 57600000,
		"./input.txt":      15862900,
	}
	testhelper.Runner(t, day15.Solve2, inputs)
}

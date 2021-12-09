package day8_test

import (
	"adventofcode/2015/day8"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	inputs := map[string]int{
		"./input_test.txt": 12,
		"./input.txt":      1350,
	}
	testhelper.Runner(t, day8.Solve1, inputs)
}

func TestSolve2(t *testing.T) {
	inputs := map[string]int{
		"./input_test2.txt": 19,
		"./input.txt":       2085,
	}
	testhelper.Runner(t, day8.Solve2, inputs)
}

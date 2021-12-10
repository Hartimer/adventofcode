package day16_test

import (
	"adventofcode/2015/day16"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	inputs := map[string]int{
		"./input.txt": 103,
	}
	testhelper.Runner(t, day16.Solve1, inputs)
}

func TestSolve2(t *testing.T) {
	inputs := map[string]int{
		"./input.txt": 405,
	}
	testhelper.Runner(t, day16.Solve2, inputs)
}

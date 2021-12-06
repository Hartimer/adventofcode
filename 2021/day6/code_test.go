package day6_test

import (
	"adventofcode/2021/day6"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	inputs := map[string]int{
		"./input_test.txt": 5934,
		"./input.txt":      365131,
	}
	testhelper.Runner(t, day6.Solve1, inputs)
}

func TestSolve2(t *testing.T) {
	inputs := map[string]int{
		"./input_test.txt": 26984457539,
		"./input.txt":      1650309278600,
	}
	testhelper.Runner(t, day6.Solve2, inputs)
}

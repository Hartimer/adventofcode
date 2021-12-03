package day2_test

import (
	"adventofcode/2021/day2"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	inputs := map[string]int{
		"./input_test.txt": 150,
		"./input.txt":      2039912,
	}
	testhelper.Runner(t, day2.Solve1, inputs)
}

func TestSolve2(t *testing.T) {
	inputs := map[string]int{
		"./input_test.txt": 900,
		"./input.txt":      1942068080,
	}
	testhelper.Runner(t, day2.Solve2, inputs)
}

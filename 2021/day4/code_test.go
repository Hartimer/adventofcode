package day4_test

import (
	"adventofcode/2021/day4"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	inputs := map[string]int{
		"./input_test.txt": 4512,
		"./input.txt":      34506,
	}
	testhelper.Runner(t, day4.Solve1, inputs)
}

func TestSolve2(t *testing.T) {
	inputs := map[string]int{
		"./input_test.txt": 1924,
		"./input.txt":      7686,
	}
	testhelper.Runner(t, day4.Solve2, inputs)
}

package day3_test

import (
	"adventofcode/2021/day3"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	inputs := map[string]int{
		"./input_test.txt": 198,
		"./input.txt":      3309596,
	}
	testhelper.Runner(t, day3.Solve1, inputs)
}

func TestSolve2(t *testing.T) {
	inputs := map[string]int{
		"./input_test.txt": 230,
		"./input.txt":      2981085,
	}
	testhelper.Runner(t, day3.Solve2, inputs)
}

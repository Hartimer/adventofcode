package day11_test

import (
	"adventofcode/2022/day11"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 10605,
		"./input.txt":      50172,
	}
	testhelper.Runner(t, day11.Solve1, input)
}

func TestSolve2(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 2713310158, // This fails?!
		"./input.txt":      11614682178,
	}
	testhelper.Runner(t, day11.Solve2, input)
}

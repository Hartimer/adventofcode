package day12_test

import (
	"adventofcode/2022/day12"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 31,
		// "./input.txt":      50172,
	}
	testhelper.Runner(t, day12.Solve1, input)
}

// func TestSolve2(t *testing.T) {
// 	input := testhelper.Inputs{
// 		"./input_test.txt": 2713310158, // This fails?!
// 		"./input.txt":      11614682178,
// 	}
// 	testhelper.Runner(t, day12.Solve2, input)
// }

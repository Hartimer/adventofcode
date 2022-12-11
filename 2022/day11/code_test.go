package day11_test

import (
	"adventofcode/2022/day11"
	"adventofcode/testhelper"
	"testing"
)

func TestSolve1(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 10605,
		// "./input.txt":      13680,
	}
	testhelper.Runner(t, day11.Solve1, input)
}

// func TestSolve2(t *testing.T) {
// 	input := testhelper.Inputs{
// 		"./input_test.txt": 1,
// 		"./input.txt":      1,
// 	}
// 	testhelper.Runner(t, day11.Solve2, input)
// }

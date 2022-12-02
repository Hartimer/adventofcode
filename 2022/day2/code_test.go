package day2_test

import (
	"adventofcode/2022/day2"
	"adventofcode/testhelper"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolve1(t *testing.T) {
	inputs := testhelper.Inputs{
		"./input_test.txt": 15,
		"./input.txt":      14297,
	}
	for filename, expected := range inputs {
		t.Run(fmt.Sprintf("%s should output %d", filename, expected), func(t *testing.T) {
			rounds := day2.ParseInputs(filename)
			result := day2.Solve1(rounds)
			require.Equal(t, expected, result)
		})
	}
}

func TestSolve2(t *testing.T) {
	inputs := testhelper.Inputs{
		"./input_test.txt": 12,
		"./input.txt":      14297,
	}
	for filename, expected := range inputs {
		t.Run(fmt.Sprintf("%s should output %d", filename, expected), func(t *testing.T) {
			rounds := day2.ParseInputs(filename)
			result := day2.Solve2(rounds)
			require.Equal(t, expected, result)
		})
	}
}

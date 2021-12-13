package day18_test

import (
	"adventofcode/2015/day18"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolve1(t *testing.T) {
	inputs := []struct {
		file       string
		iterations int
		expected   int
	}{
		{
			file:       "./input_test.txt",
			iterations: 4,
			expected:   4,
		},
		{
			file:       "./input.txt",
			iterations: 100,
			expected:   821,
		},
	}

	for _, input := range inputs {
		t.Run(fmt.Sprintf("%s with iterations %d should produce %d", input.file, input.iterations, input.expected), func(t *testing.T) {
			result, err := day18.Solve1(input.file, input.iterations)
			require.NoError(t, err)
			require.Equal(t, input.expected, result)
		})
	}
}

func TestSolve2(t *testing.T) {
	inputs := []struct {
		file       string
		iterations int
		expected   int
	}{
		{
			file:       "./input_test.txt",
			iterations: 5,
			expected:   17,
		},
		{
			file:       "./input.txt",
			iterations: 100,
			expected:   886,
		},
	}

	for _, input := range inputs {
		t.Run(fmt.Sprintf("%s with iterations %d should produce %d", input.file, input.iterations, input.expected), func(t *testing.T) {
			result, err := day18.Solve2(input.file, input.iterations)
			require.NoError(t, err)
			require.Equal(t, input.expected, result)
		})
	}
}

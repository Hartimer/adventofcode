package day17_test

import (
	"adventofcode/2015/day17"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolve1(t *testing.T) {
	inputs := []struct {
		file     string
		target   int
		expected int
	}{
		{
			file:     "./input_test.txt",
			target:   25,
			expected: 4,
		},
		{
			file:     "./input.txt",
			target:   150,
			expected: 1638,
		},
	}

	for _, input := range inputs {
		t.Run(fmt.Sprintf("%s with target %d should produce %d", input.file, input.target, input.expected), func(t *testing.T) {
			result, err := day17.Solve1(input.file, input.target)
			require.NoError(t, err)
			require.Equal(t, input.expected, result)
		})
	}
}

func TestSolve2(t *testing.T) {
	inputs := []struct {
		file     string
		target   int
		expected int
	}{
		{
			file:     "./input_test.txt",
			target:   25,
			expected: 3,
		},
		{
			file:     "./input.txt",
			target:   150,
			expected: 17,
		},
	}

	for _, input := range inputs {
		t.Run(fmt.Sprintf("%s with target %d should produce %d", input.file, input.target, input.expected), func(t *testing.T) {
			result, err := day17.Solve2(input.file, input.target)
			require.NoError(t, err)
			require.Equal(t, input.expected, result)
		})
	}
}

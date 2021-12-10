package day14_test

import (
	"adventofcode/2015/day14"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolve1(t *testing.T) {
	inputs := []struct {
		file     string
		duration int
		expected int
	}{
		{
			file:     "./input_test.txt",
			duration: 1000,
			expected: 1120,
		},
		{
			file:     "./input.txt",
			duration: 2503,
			expected: 2655,
		},
	}
	for _, input := range inputs {
		t.Run(fmt.Sprintf("File %s with duration %d should produce %d", input.file, input.duration, input.expected), func(t *testing.T) {
			result, err := day14.Solve1(input.file, input.duration)
			require.NoError(t, err)
			require.Equal(t, input.expected, result)
		})
	}
}

func TestSolve2(t *testing.T) {
	inputs := []struct {
		file     string
		duration int
		expected int
	}{
		{
			file:     "./input_test.txt",
			duration: 1000,
			expected: 689,
		},
		{
			file:     "./input.txt",
			duration: 2503,
			expected: 1059,
		},
	}
	for _, input := range inputs {
		t.Run(fmt.Sprintf("File %s with duration %d should produce %d", input.file, input.duration, input.expected), func(t *testing.T) {
			result, err := day14.Solve2(input.file, input.duration)
			require.NoError(t, err)
			require.Equal(t, input.expected, result)
		})
	}
}

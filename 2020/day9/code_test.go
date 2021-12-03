package day9_test

import (
	"adventofcode/2020/day9"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolve1(t *testing.T) {
	inputs := []struct {
		file         string
		expected     int
		preambleSize int
	}{
		{
			file:         "./input_test.txt",
			expected:     127,
			preambleSize: 5,
		},
		{
			file:         "./input.txt",
			expected:     144381670,
			preambleSize: 25,
		},
	}
	for _, input := range inputs {
		t.Run(fmt.Sprintf("%s with preamble %d should output %d", input.file, input.preambleSize, input.expected), func(t *testing.T) {
			result, err := day9.Solve1(input.file, input.preambleSize)
			require.NoError(t, err)
			require.Equal(t, input.expected, result)
		})
	}
}

func TestSolve2(t *testing.T) {
	inputs := []struct {
		file          string
		expected      int
		invalidNumber int
	}{
		{
			file:          "./input_test.txt",
			expected:      62,
			invalidNumber: 127,
		},
		{
			file:          "./input.txt",
			expected:      20532569,
			invalidNumber: 144381670,
		},
	}
	for _, input := range inputs {
		t.Run(fmt.Sprintf("%s with invalid number %d should output %d", input.file, input.invalidNumber, input.expected), func(t *testing.T) {
			result, err := day9.Solve2(input.file, input.invalidNumber)
			require.NoError(t, err)
			require.Equal(t, input.expected, result)
		})
	}
}

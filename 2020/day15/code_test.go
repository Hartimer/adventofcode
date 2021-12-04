package day15_test

import (
	"adventofcode/2020/day15"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolve1(t *testing.T) {
	inputs := []struct {
		startingNumbers []int
		expected        int
	}{
		{
			startingNumbers: []int{0, 3, 6},
			expected:        436,
		},
		{
			startingNumbers: []int{1, 3, 2},
			expected:        1,
		},
		{
			startingNumbers: []int{2, 1, 3},
			expected:        10,
		},
		{
			startingNumbers: []int{1, 2, 3},
			expected:        27,
		},
		{
			startingNumbers: []int{2, 3, 1},
			expected:        78,
		},
		{
			startingNumbers: []int{3, 2, 1},
			expected:        438,
		},
		{
			startingNumbers: []int{3, 1, 2},
			expected:        1836,
		},
		{
			startingNumbers: []int{0, 5, 4, 1, 10, 14, 7},
			expected:        203,
		},
	}
	for _, input := range inputs {
		t.Run(fmt.Sprintf("Starting numbers %v should produce %d", input.startingNumbers, input.expected), func(t *testing.T) {
			require.Equal(t, input.expected, day15.Solve1(input.startingNumbers))
		})
	}
}

func TestSolve2(t *testing.T) {
	inputs := []struct {
		startingNumbers []int
		expected        int
	}{
		{
			startingNumbers: []int{0, 3, 6},
			expected:        175594,
		},
		{
			startingNumbers: []int{1, 3, 2},
			expected:        2578,
		},
		{
			startingNumbers: []int{2, 1, 3},
			expected:        3544142,
		},
		{
			startingNumbers: []int{1, 2, 3},
			expected:        261214,
		},
		{
			startingNumbers: []int{2, 3, 1},
			expected:        6895259,
		},
		{
			startingNumbers: []int{3, 2, 1},
			expected:        18,
		},
		{
			startingNumbers: []int{3, 1, 2},
			expected:        362,
		},
		{
			startingNumbers: []int{0, 5, 4, 1, 10, 14, 7},
			expected:        9007186,
		},
	}
	for _, input := range inputs {
		t.Run(fmt.Sprintf("Starting numbers %v should produce %d", input.startingNumbers, input.expected), func(t *testing.T) {
			require.Equal(t, input.expected, day15.Solve2(input.startingNumbers))
		})
	}
}

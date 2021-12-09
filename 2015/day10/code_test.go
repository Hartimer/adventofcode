package day10_test

import (
	"adventofcode/2015/day10"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolve1(t *testing.T) {
	t.Parallel()
	inputs := []struct {
		str        string
		iterations int
		expected   int
	}{
		{str: "1321131112", iterations: 40, expected: 492982},
		{str: "1321131112", iterations: 50, expected: 6989950},
	}
	for _, input := range inputs {
		t.Run(fmt.Sprintf("%s with %d iterations should produce %d", input.str, input.iterations, input.expected), func(t *testing.T) {
			result := day10.Solve(input.str, input.iterations)
			require.Equal(t, input.expected, result)
		})
	}
}

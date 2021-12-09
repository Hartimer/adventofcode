package day11_test

import (
	"adventofcode/2015/day11"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolve1(t *testing.T) {
	inputs := map[string]string{
		"abcdefgf": "abcdffaa",
		"ghijklmn": "ghjaabcc",
		"cqjxjnds": "cqjxxyzz",
		"cqjxxyzz": "cqkaabcc",
	}
	for password, expected := range inputs {
		t.Run(fmt.Sprintf("Valid password after %s should be %s", password, expected), func(t *testing.T) {
			result := day11.Solve1(password)
			require.Equal(t, expected, result)
		})
	}
}

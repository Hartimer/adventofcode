package day4_test

import (
	"adventofcode/2015/day4"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolve1(t *testing.T) {
	inputs := map[string]int{
		"abcdef":   609043,
		"pqrstuv":  1048970,
		"ckczppom": 117946,
	}
	for prefix, expected := range inputs {
		t.Run(fmt.Sprintf("Prefix %s expects int %d", prefix, expected), func(t *testing.T) {
			require.Equal(t, expected, day4.Solve(prefix, "00000"))
		})
	}
}

func TestSolve2(t *testing.T) {
	inputs := map[string]int{
		"ckczppom": 3938038,
	}
	for prefix, expected := range inputs {
		t.Run(fmt.Sprintf("Prefix %s expects int %d", prefix, expected), func(t *testing.T) {
			require.Equal(t, expected, day4.Solve(prefix, "000000"))
		})
	}
}

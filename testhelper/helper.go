package testhelper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type Solver func(string) (int, error)

type Inputs map[string]int

func Runner(t *testing.T, solver Solver, inputs Inputs) {
	for filename, expected := range inputs {
		t.Run(fmt.Sprintf("%s should output %d", filename, expected), func(t *testing.T) {
			result, err := solver(filename)
			require.NoError(t, err)
			require.Equal(t, expected, result)
		})
	}
}

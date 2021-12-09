package testhelper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// Solver describes a common function signature used in solving challenges
type Solver func(string) (int, error)

// Inputs represent a mapping between a key (filepath usually) and the expected result
type Inputs map[string]int

// Runner runs commonly used test signatures, specifically of type Solver
func Runner(t *testing.T, solver Solver, inputs Inputs) {
	for filename, expected := range inputs {
		t.Run(fmt.Sprintf("%s should output %d", filename, expected), func(t *testing.T) {
			result, err := solver(filename)
			require.NoError(t, err)
			require.Equal(t, expected, result)
		})
	}
}

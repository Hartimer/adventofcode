package day21_test

import (
	"adventofcode/2021/day21"
	"adventofcode/testhelper"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolve1(t *testing.T) {
	input := testhelper.Inputs{
		"./input_test.txt": 739785,
		"./input.txt":      412344,
	}
	testhelper.Runner(t, day21.Solve1, input)
}

func TestSolve2(t *testing.T) {
	inputs := map[string]int64{
		"./input_test.txt": 444356092776315,
		// "./input.txt": 0,
	}
	for filename, expected := range inputs {
		t.Run(fmt.Sprintf("%s should output %d", filename, expected), func(t *testing.T) {
			result, err := day21.Solve2(filename)
			require.NoError(t, err)
			require.Equal(t, expected, result)
		})
	}

}

/*

1
2
3

11
21
22
31
32
33

111
211
221
222
311
321
322
331
332
333

11
21
22
31
32
33

111
211
221
222
311
312
322
331
332
333


*/

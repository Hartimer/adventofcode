package _test

import (
	"adventofcode/helper"
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type day7Op func(int, int) int

func add(a, b int) int { return a + b }
func mul(a, b int) int { return a * b }
func concat(a, b int) int {
	return a*int(math.Pow10(len(strconv.Itoa(b)))) + b
}

func TestDay7_1(t *testing.T) {
	inputs := []struct {
		filename string
		expected int
	}{
		{
			filename: "day7.1.1.input",
			expected: 3749,
		},
		{
			filename: "day7.1.input",
			expected: 1545311493300,
		},
	}

	for _, input := range inputs {
		t.Run(fmt.Sprintf("%s produces %d", input.filename, input.expected), func(t *testing.T) {
			resolver(t, input.filename, input.expected, []day7Op{add, mul})
		})
	}
}
func TestDay7_2(t *testing.T) {
	inputs := []struct {
		filename string
		expected int
	}{
		{
			filename: "day7.1.1.input",
			expected: 11387,
		},
		{
			filename: "day7.1.input",
			expected: 169122112716571,
		},
	}

	for _, input := range inputs {
		t.Run(fmt.Sprintf("%s produces %d", input.filename, input.expected), func(t *testing.T) {
			resolver(t, input.filename, input.expected, []day7Op{add, mul, concat})
		})
	}
}

func resolver(t *testing.T, filename string, expected int, ops []day7Op) {
	accumulator := 0
	for fileLine := range helper.FileLineReader(filename) {
		parts := strings.Split(fileLine, ": ")
		require.Len(t, parts, 2)
		expectedResult, err := strconv.Atoi(parts[0])
		require.NoError(t, err)
		rawParcels := strings.Split(parts[1], " ")
		parcels, err := stringSliceToInt(rawParcels)
		require.NoError(t, err)
		_, isValid := calculate(expectedResult, 0, parcels, ops)
		if isValid {
			accumulator += expectedResult
		}
	}
	require.Equal(t, expected, accumulator)
}

func calculate(targetTotal, runningTotal int, remaining []int, ops []day7Op) (int, bool) {
	if len(remaining) == 0 {
		return runningTotal, runningTotal == targetTotal
	}
	if runningTotal > targetTotal {
		return 0, false
	}
	next := remaining[0]
	for _, op := range ops {
		nextTotal := op(runningTotal, next)
		_, isValid := calculate(targetTotal, nextTotal, remaining[1:], ops)
		if isValid {
			return 0, true
		}
	}
	return 0, false
}

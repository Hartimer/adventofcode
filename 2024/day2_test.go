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

func TestDay2_1(t *testing.T) {
	inputs := []struct {
		filename string
		expected int
	}{
		{
			filename: "day2.1.1.input",
			expected: 2,
		},
		{
			filename: "day2.1.input",
			expected: 321,
		},
	}
	for _, input := range inputs {
		t.Run(fmt.Sprintf("%s produces %d", input.filename, input.expected), func(t *testing.T) {
			safeCount := 0
			for fileLine := range helper.FileLineReader(input.filename) {
				levelStrings := strings.Split(fileLine, " ")
				levels, err := strsToInts(levelStrings)
				require.NoError(t, err)

				if solve(levels) {
					safeCount++
				}
			}
			require.Equal(t, input.expected, safeCount)
		})
	}
}

func TestDay2_2(t *testing.T) {
	inputs := []struct {
		filename string
		expected int
	}{
		{
			filename: "day2.1.1.input",
			expected: 4,
		},
		{
			filename: "day2.1.input",
			expected: 386, // not 375,363,410,460,594,366,367
		},
	}
	for _, input := range inputs {
		t.Run(fmt.Sprintf("%s produces %d", input.filename, input.expected), func(t *testing.T) {
			safeCount := 0
			for fileLine := range helper.FileLineReader(input.filename) {
				levelStrings := strings.Split(fileLine, " ")
				levels, err := strsToInts(levelStrings)
				require.NoError(t, err)

				if solve(levels) {
					safeCount++
					continue
				}

				for i := 0; i < len(levels); i++ {
					cleanedLevels := make([]int, 0, len(levels)-1)
					cleanedLevels = append(cleanedLevels, levels[:i]...)
					cleanedLevels = append(cleanedLevels, levels[i+1:]...)
					if solve(cleanedLevels) {
						safeCount++
						break
					}
				}
			}
			require.Equal(t, input.expected, safeCount)
		})
	}
}

func solve(levels []int) bool {
	increasing := levels[0] < levels[1]

	for i := 1; i < len(levels); i++ {
		currentLevel := levels[i]
		lastLevel := levels[i-1]
		diff := math.Abs(float64(currentLevel - lastLevel))

		if (currentLevel < lastLevel && increasing) ||
			(currentLevel > lastLevel && !increasing) ||
			diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func strsToInts(levelStrs []string) ([]int, error) {
	aa := make([]int, 0, len(levelStrs))
	for _, l := range levelStrs {
		lInt, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		aa = append(aa, lInt)
	}
	return aa, nil
}

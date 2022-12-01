package day1_test

import (
	"adventofcode/2022/day1"
	"adventofcode/helper"
	"adventofcode/testhelper"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolve1(t *testing.T) {
	inputs := testhelper.Inputs{
		"./input_test.txt": 24000,
		"./input.txt":      73211,
	}
	for filename, expected := range inputs {
		t.Run(fmt.Sprintf("%s should output %d", filename, expected), func(t *testing.T) {
			elfCalories := [][]int{}
			currentIdx := 0
			for fileLine := range helper.FileLineReader(filename) {
				if len(fileLine) == 0 {
					currentIdx++
					continue
				}
				if len(elfCalories) <= currentIdx {
					elfCalories = append(elfCalories, []int{})
				}
				calories, err := strconv.Atoi(fileLine)
				require.NoError(t, err)
				elfCalories[currentIdx] = append(elfCalories[currentIdx], calories)

			}
			result := day1.Solve1(elfCalories)
			require.Equal(t, expected, result)
		})
	}
}

func TestSolve2(t *testing.T) {
	inputs := testhelper.Inputs{
		"./input_test.txt": 45000,
		"./input.txt":      213958,
	}
	for filename, expected := range inputs {
		t.Run(fmt.Sprintf("%s should output %d", filename, expected), func(t *testing.T) {
			elfCalories := [][]int{}
			currentIdx := 0
			for fileLine := range helper.FileLineReader(filename) {
				if len(fileLine) == 0 {
					currentIdx++
					continue
				}
				if len(elfCalories) <= currentIdx {
					elfCalories = append(elfCalories, []int{})
				}
				calories, err := strconv.Atoi(fileLine)
				require.NoError(t, err)
				elfCalories[currentIdx] = append(elfCalories[currentIdx], calories)

			}
			result := day1.Solve2(elfCalories)
			require.Equal(t, expected, result)
		})
	}
}

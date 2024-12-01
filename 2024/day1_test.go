package _test

import (
	"adventofcode/helper"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay1_1(t *testing.T) {
	inputs := []struct {
		filename string
		expected int
	}{
		{
			filename: "day1.1.1.input",
			expected: 11,
		},
		{
			filename: "day1.1.input",
			expected: 2057374,
		},
	}
	for _, input := range inputs {
		t.Run(fmt.Sprintf("%s produces %d", input.filename, input.expected), func(t *testing.T) {
			var right, left []int
			for fileLine := range helper.FileLineReader(input.filename) {
				require.NotEmpty(t, fileLine)

				parts := strings.Split(fileLine, " ")
				hasRight := false
				for _, part := range parts {
					if part == " " || part == "" {
						continue
					}
					number, err := strconv.Atoi(part)
					require.NoError(t, err)
					if hasRight {
						left = append(left, number)
					} else {
						right = append(right, number)
						hasRight = true
					}
				}
			}
			sort.Ints(right)
			sort.Ints(left)
			total := 0
			for idx := range right {
				total += int(math.Abs(float64(right[idx] - left[idx])))
			}
			require.Equal(t, input.expected, total)
		})
	}
}

func TestDay1_2(t *testing.T) {
	inputs := []struct {
		filename string
		expected int
	}{
		{
			filename: "day1.1.1.input",
			expected: 31,
		},
		{
			filename: "day1.1.input",
			expected: 23177084,
		},
	}
	for _, input := range inputs {
		t.Run(fmt.Sprintf("%s produces %d", input.filename, input.expected), func(t *testing.T) {
			var right []int
			left := map[int]int{}
			for fileLine := range helper.FileLineReader(input.filename) {
				parts := strings.Split(fileLine, " ")
				hasRight := false
				for _, part := range parts {
					if part == " " || part == "" {
						continue
					}
					number, err := strconv.Atoi(part)
					require.NoError(t, err)
					if hasRight {
						if _, exists := left[number]; !exists {
							left[number] = 1
						} else {
							left[number]++
						}
					} else {
						right = append(right, number)
						hasRight = true
					}
				}
			}
			total := 0
			for _, r := range right {
				if count, exists := left[r]; exists {
					total += (r * count)
				}
			}
			require.Equal(t, input.expected, total)
		})
	}
}

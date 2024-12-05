package _test

import (
	"adventofcode/helper"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay5_1(t *testing.T) {
	inputs := []struct {
		filename string
		expected int
	}{
		{
			filename: "day5.1.1.input",
			expected: 143,
		},
		{
			filename: "day5.1.input",
			expected: 4662,
		},
	}

	for _, input := range inputs {
		t.Run(fmt.Sprintf("%s produces %d", input.filename, input.expected), func(t *testing.T) {
			rules := map[int][]int{}
			ruleMode := true
			total := 0
			for fileLine := range helper.FileLineReader(input.filename) {
				if ruleMode {
					if strings.TrimSpace(fileLine) == "" {
						ruleMode = false
						continue
					}
					parts := strings.Split(fileLine, "|")
					require.Len(t, parts, 2)
					page1, err := strconv.Atoi(parts[0])
					require.NoError(t, err)
					page2, err := strconv.Atoi(parts[1])
					require.NoError(t, err)
					rules[page1] = append(rules[page1], page2)
				} else {
					parts := strings.Split(fileLine, ",")
					seenPages := map[int]struct{}{}
					sequence, err := stringSliceToInt(parts)
					require.NoError(t, err)
					validSequence := true
					for _, page := range sequence {
						validSequence = validPage(rules[page], seenPages)
						if !validSequence {
							break
						}
						seenPages[page] = struct{}{}
					}
					if validSequence {
						total += sequence[len(sequence)/2]
					}
				}
			}
			require.False(t, ruleMode)
			require.Equal(t, input.expected, total)
		})
	}
}

func validPage(rules []int, seen map[int]struct{}) bool {
	for _, rule := range rules {
		if _, exists := seen[rule]; exists {
			return false
		}
	}
	return true
}

func TestDay5_2(t *testing.T) {
	inputs := []struct {
		filename string
		expected int
	}{
		{
			filename: "day5.1.1.input",
			expected: 123,
		},
		{
			filename: "day5.1.input",
			expected: 5900,
		},
	}

	for _, input := range inputs {
		t.Run(fmt.Sprintf("%s produces %d", input.filename, input.expected), func(t *testing.T) {
			var invalidSequences [][]int
			rules := map[int][]int{}
			ruleMode := true
			total := 0
			for fileLine := range helper.FileLineReader(input.filename) {
				if ruleMode {
					if strings.TrimSpace(fileLine) == "" {
						ruleMode = false
						continue
					}
					parts := strings.Split(fileLine, "|")
					require.Len(t, parts, 2)
					page1, err := strconv.Atoi(parts[0])
					require.NoError(t, err)
					page2, err := strconv.Atoi(parts[1])
					require.NoError(t, err)
					rules[page1] = append(rules[page1], page2)
				} else {
					parts := strings.Split(fileLine, ",")
					pages, err := stringSliceToInt(parts)
					require.NoError(t, err)
					seenPages := map[int]struct{}{}
					validSequence := true
					for _, page := range pages {
						validSequence = validPage(rules[page], seenPages)
						if !validSequence {
							invalidSequences = append(invalidSequences, pages)
							break
						}
						seenPages[page] = struct{}{}
					}
				}
			}
			for _, sequence := range invalidSequences {
				slices.SortFunc(sequence, func(a, b int) int {
					dependencies, exists := rules[a]
					if exists && slices.Contains(dependencies, b) {
						return -1
					}
					dependencies, exists = rules[b]
					if exists && slices.Contains(dependencies, a) {
						return 1
					}
					return 0
				})
				total += sequence[len(sequence)/2]
			}

			require.False(t, ruleMode)
			require.Equal(t, input.expected, total)
		})
	}
}

func stringSliceToInt(s []string) ([]int, error) {
	r := make([]int, 0, len(s))
	for _, entry := range s {
		n, err := strconv.Atoi(entry)
		if err != nil {
			return nil, err
		}
		r = append(r, n)
	}
	return r, nil
}

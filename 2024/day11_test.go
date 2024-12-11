package _test

import (
	"adventofcode/helper"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type stone int

type stoneRule func(stone) ([]stone, bool)

func zero(s stone) ([]stone, bool) {
	if s == 0 {
		return []stone{1}, true
	}
	return nil, false
}

func even(s stone) ([]stone, bool) {
	digitStr := strconv.Itoa(int(s))
	if len(digitStr)%2 == 0 {
		part1, part2 := digitStr[:len(digitStr)/2], digitStr[len(digitStr)/2:]
		d1, _ := strconv.Atoi(part1)
		d2, _ := strconv.Atoi(part2)
		return []stone{stone(d1), stone(d2)}, true
	}
	return nil, false
}

func default2024(s stone) ([]stone, bool) {
	s *= 2024
	return []stone{s}, true
}

var rules = []stoneRule{zero, even, default2024}

func TestDay11_1(t *testing.T) {
	inputs := []struct {
		filename string
		rounds   int
		expected int
	}{
		{
			filename: "day11.1.1.input",
			rounds:   25,
			expected: 55312,
		},
		{
			filename: "day11.1.input",
			rounds:   25,
			expected: 175006,
		},
		// {
		// 	filename: "day11.1.input",
		// 	rounds:   75,
		// 	expected: 0,
		// },
	}

	for _, input := range inputs {
		t.Run(fmt.Sprintf("%s %d times produces %d", input.filename, input.rounds, input.expected), func(t *testing.T) {
			for fileLine := range helper.FileLineReader(input.filename) {
				numberStrs := strings.Split(fileLine, " ")
				result, err := strsToStones(numberStrs)
				require.NoError(t, err)
				for i := 0; i < input.rounds; i++ {
					var newStoneLine []stone
					for _, s := range result {
						for _, rule := range rules {
							newStones, applied := rule(s)
							if !applied {
								continue
							}
							newStoneLine = append(newStoneLine, newStones...)
							break
						}
					}
					result = newStoneLine
				}
				require.Equal(t, input.expected, len(result))
			}
		})
	}
}

func traceStone(s stone, history map[stone][]stone) []stone {
	if stoneSplit, isKnown := history[s]; isKnown {
		return stoneSplit
	}
	for _, rule := range rules {
		newStones, applied := rule(s)
		if !applied {
			continue
		}
		history[s] = newStones
		return newStones
	}
	panic("should never happen")
}

func strsToStones(levelStrs []string) ([]stone, error) {
	aa := make([]stone, 0, len(levelStrs))
	for _, l := range levelStrs {
		lInt, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		aa = append(aa, stone(lInt))
	}
	return aa, nil
}

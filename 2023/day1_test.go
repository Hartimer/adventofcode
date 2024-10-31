package _test

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/require"
)

func day1(t *testing.T, filename string, lineChanger func(string) string) int {
	f, err := os.Open(filename)
	require.NoError(t, err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var total int
	for scanner.Scan() {
		line := scanner.Text()
		if lineChanger != nil {
			line = lineChanger(line)
		}
		var firstDigit, lastDigit rune
		for _, r := range line {
			if !unicode.IsDigit(r) {
				continue
			}
			if firstDigit == 0 {
				firstDigit = r
			} else {
				lastDigit = r
			}
		}
		if lastDigit == 0 {
			lastDigit = firstDigit
		}
		number, err := strconv.Atoi(string([]rune{firstDigit, lastDigit}))
		require.NoError(t, err)
		total += number
	}
	require.NoError(t, scanner.Err())
	return total
}

func TestPart1(t *testing.T) {
	inputs := []struct {
		filename string
		expected int
	}{
		{
			filename: "day1.1.1.input",
			expected: 142,
		},
		{
			filename: "day1.input",
			expected: 54940,
		},
	}

	for _, input := range inputs {
		t.Run(input.filename, func(t *testing.T) {
			total := day1(t, input.filename, nil)
			require.Equal(t, input.expected, total)
		})
	}
}

func TestPart2(t *testing.T) {
	inputs := []struct {
		filename string
		expected int
	}{
		{
			filename: "day1.2.1.input",
			expected: 281,
		},
		{
			filename: "day1.input",
			expected: 54208,
		},
	}

	for _, input := range inputs {
		t.Run(input.filename, func(t *testing.T) {
			total := day1(t, input.filename, func(s string) string {
				s = strings.ReplaceAll(s, "one", "on1e")
				s = strings.ReplaceAll(s, "two", "tw2o")
				s = strings.ReplaceAll(s, "three", "thre3e")
				s = strings.ReplaceAll(s, "four", "fou4r")
				s = strings.ReplaceAll(s, "five", "fiv5e")
				s = strings.ReplaceAll(s, "six", "si6x")
				s = strings.ReplaceAll(s, "seven", "seve7n")
				s = strings.ReplaceAll(s, "eight", "eigh8t")
				s = strings.ReplaceAll(s, "nine", "nin9e")
				return s
			})
			require.Equal(t, input.expected, total)
		})
	}
}

package _test

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/exp/maps"
)

func TestDay2Part1(t *testing.T) {
	inputs := []struct {
		filename string
		expected int
		red      int
		green    int
		blue     int
	}{
		{
			filename: "day2.1.1.input",
			red:      12,
			green:    13,
			blue:     14,
			expected: 8,
		},
		{
			filename: "day2.input",
			red:      12,
			green:    13,
			blue:     14,
			expected: 2164,
		},
	}

	for _, input := range inputs {
		t.Run(input.filename, func(t *testing.T) {
			ballExp := regexp.MustCompile(`\d+ (red|green|blue)`)
			f, err := os.Open(input.filename)
			require.NoError(t, err)
			defer f.Close()

			var total int
			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				line := scanner.Text()
				colonIdx := strings.Index(line, ":")
				require.True(t, colonIdx > 0)
				validGame := true

				ballCounts := ballExp.FindAllString(line[colonIdx:], 1000)
				require.NotNil(t, ballCounts)
				for _, ballCount := range ballCounts {
					parts := strings.Split(ballCount, " ")
					require.Len(t, parts, 2)
					count, err := strconv.Atoi(parts[0])
					require.NoError(t, err)
					color := parts[1]
					var maxAllowed int
					switch color {
					case "red":
						maxAllowed = input.red
					case "green":
						maxAllowed = input.green
					case "blue":
						maxAllowed = input.blue
					}
					if validGame = count <= maxAllowed; !validGame {
						break
					}
				}
				if validGame {
					gameID, err := strconv.Atoi(strings.TrimPrefix(line[:colonIdx], "Game "))
					require.NoError(t, err)
					total += gameID
				}
			}
			require.NoError(t, scanner.Err())
			require.Equal(t, input.expected, total)
		})
	}
}

func TestDay2Part2(t *testing.T) {
	inputs := []struct {
		filename string
		expected int
	}{
		{
			filename: "day2.1.1.input",
			expected: 2286,
		},
		{
			filename: "day2.input",
			expected: 69929,
		},
	}

	for _, input := range inputs {
		t.Run(input.filename, func(t *testing.T) {
			ballExp := regexp.MustCompile(`\d+ (red|green|blue)`)
			f, err := os.Open(input.filename)
			require.NoError(t, err)
			defer f.Close()

			var total int
			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				line := scanner.Text()
				colonIdx := strings.Index(line, ":")
				require.True(t, colonIdx > 0)
				gameMinimums := map[string]int{}

				ballCounts := ballExp.FindAllString(line[colonIdx:], 1000)
				require.NotNil(t, ballCounts)
				for _, ballCount := range ballCounts {
					parts := strings.Split(ballCount, " ")
					require.Len(t, parts, 2)
					count, err := strconv.Atoi(parts[0])
					require.NoError(t, err)
					color := parts[1]
					currentMin, exists := gameMinimums[color]
					if !exists || count > currentMin {
						gameMinimums[color] = count
					}
				}
				gamePower := 1
				for _, mins := range maps.Values(gameMinimums) {
					gamePower *= mins
				}
				total += gamePower
			}

			require.NoError(t, scanner.Err())
			require.Equal(t, input.expected, total)
		})
	}
}

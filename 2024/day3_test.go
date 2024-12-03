package _test

import (
	"adventofcode/helper"
	"fmt"
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/require"
)

func TestDay3_1(t *testing.T) {
	inputs := []struct {
		filename string
		expected int
	}{
		{
			filename: "day3.1.1.input",
			expected: 161,
		},
		{
			filename: "day3.1.input",
			expected: 173785482,
		},
	}

	for _, input := range inputs {
		t.Run(fmt.Sprintf("%s produces %d", input.filename, input.expected), func(t *testing.T) {
			total := 0
			for fileLine := range helper.FileLineReader(input.filename) {
				for {
					idx := strings.Index(fileLine, "mul(")
					if idx == -1 {
						break
					}
					fileLine = fileLine[idx+4:]
					firstNumber := 0
					numberDigits := 0
					invalid := false
					for _, n := range fileLine {
						if unicode.IsDigit(n) {
							if firstNumber > 0 {
								firstNumber *= 10
							}
							firstNumber += int(n - '0')
							numberDigits++
						} else {
							invalid = n != ','
							break
						}
					}
					if invalid {
						fileLine = fileLine[numberDigits+1:]
						continue
					}
					secondNumber := 0
					secondNumberDigits := 0
					for _, n := range fileLine[numberDigits+1:] {
						if unicode.IsDigit(n) {
							if secondNumber > 0 {
								secondNumber *= 10
							}
							secondNumber += int(n - '0')
							secondNumberDigits++
						} else {
							invalid = n != ')'
							break
						}
					}
					if invalid {
						fileLine = fileLine[numberDigits+1+secondNumberDigits+1:]
						continue
					}
					total += (firstNumber * secondNumber)
				}
			}
			require.Equal(t, input.expected, total)
		})
	}
}

func TestDay3_2(t *testing.T) {
	inputs := []struct {
		filename string
		expected int
	}{
		{
			filename: "day3.2.1.input",
			expected: 48,
		},
		{
			filename: "day3.1.input",
			expected: 83158140, // not 90500439
		},
	}
	for _, input := range inputs {
		t.Run(fmt.Sprintf("%s produces %d", input.filename, input.expected), func(t *testing.T) {
			total := 0
			enabled := true
			for fileLine := range helper.FileLineReader(input.filename) {
				for {
					mulIdx := strings.Index(fileLine, "mul(")
					if mulIdx == -1 {
						break
					}
					if enabled {
						mulDont := strings.Index(fileLine, "don't()")
						if mulDont != -1 && mulDont < mulIdx {
							enabled = false
							fileLine = fileLine[mulDont+7:]
							continue
						}
					} else {
						mulDo := strings.Index(fileLine, "do()")
						if mulDo != -1 && mulDo > mulIdx {
							enabled = true
							fileLine = fileLine[mulDo+4:]
							continue
						}
					}
					if !enabled {
						fileLine = fileLine[mulIdx+4:]
						continue
					}
					fileLine = fileLine[mulIdx+4:]
					firstNumber := 0
					numberDigits := 0
					invalid := false
					for _, n := range fileLine {
						if unicode.IsDigit(n) {
							if firstNumber > 0 {
								firstNumber *= 10
							}
							firstNumber += int(n - '0')
							numberDigits++
						} else {
							invalid = n != ','
							break
						}
					}
					if invalid {
						fileLine = fileLine[numberDigits+1:]
						continue
					}
					secondNumber := 0
					secondNumberDigits := 0
					for _, n := range fileLine[numberDigits+1:] {
						if unicode.IsDigit(n) {
							if secondNumber > 0 {
								secondNumber *= 10
							}
							secondNumber += int(n - '0')
							secondNumberDigits++
						} else {
							invalid = n != ')'
							break
						}
					}
					if invalid {
						fileLine = fileLine[numberDigits+1+secondNumberDigits+1:]
						continue
					}
					total += (firstNumber * secondNumber)
				}
			}
			require.Equal(t, input.expected, total)
		})
	}
}

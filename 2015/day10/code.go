package day10

import (
	"fmt"
)

func Solve(input string, iterations int) int {
	result := []byte(input)
	for i := 0; i < iterations; i++ {
		result = convert(result)
	}
	return len(string(result))
}

func convert(str []byte) []byte {
	var currentChar byte
	var currentCharCount int
	var result []byte

	for _, c := range str {
		if currentCharCount == 0 {
			currentChar = c
			currentCharCount = 1
			continue
		}

		if currentChar == c {
			currentCharCount++
		} else {
			result = append(result, []byte(fmt.Sprint(currentCharCount))...)
			result = append(result, currentChar)
			currentChar = c
			currentCharCount = 1
		}
	}
	result = append(result, []byte(fmt.Sprint(currentCharCount))...)
	return append(result, currentChar)
}

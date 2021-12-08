package day5

import (
	"adventofcode/helper"
	"strings"
)

type validationRule func(string) bool

var hasVowels = func(s string) bool {
	vowelCount := 0
	for _, c := range s {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			vowelCount++
		}
	}
	return vowelCount >= 3
}

var hasDoubleChar = func(s string) bool {
	var previousChar rune
	for _, c := range s {
		if c == previousChar {
			return true
		}
		previousChar = c
	}
	return false
}

var noDisallowedStrings = func(s string) bool {
	disallowedStrings := []string{"ab", "cd", "pq", "xy"}
	for _, disallowedString := range disallowedStrings {
		if strings.Contains(s, disallowedString) {
			return false
		}
	}
	return true
}

func Solve1(inputFilePath string) (int, error) {
	return solve(inputFilePath, hasVowels, hasDoubleChar, noDisallowedStrings), nil
}

var pairRepeats = func(s string) bool {
	for workingIdx := 2; workingIdx < len(s); workingIdx++ {
		pair := s[workingIdx-2 : workingIdx]
		remainingStr := s[workingIdx:]
		if strings.Contains(remainingStr, pair) {
			return true
		}
	}
	return false
}

var letterRepeatsWithMiddle = func(s string) bool {
	for idx, c := range s {
		if idx+2 < len(s) && rune(s[idx+2]) == c {
			return true
		}
	}
	return false
}

func Solve2(inputFilePath string) (int, error) {
	return solve(inputFilePath, pairRepeats, letterRepeatsWithMiddle), nil
}

func solve(inputFilePath string, rules ...validationRule) int {
	niceStringsCount := 0
	for fileLine := range helper.FileLineReader(inputFilePath) {
		isNice := true
		for _, rule := range rules {
			if !rule(fileLine) {
				isNice = false
			}
		}
		if isNice {
			niceStringsCount++
		}
	}
	return niceStringsCount
}
package day3

import (
	"adventofcode/helper"
)

func Solve1(inputFilePath string) (int, error) {
	total := 0
	for fileLine := range helper.FileLineReader(inputFilePath) {
		firstSack := fileLine[:len(fileLine)/2]
		secondSack := fileLine[len(fileLine)/2:]

		contents := map[rune]struct{}{}
		for _, letter := range firstSack {
			contents[letter] = struct{}{}
		}

		for _, letter := range secondSack {
			if _, exists := contents[letter]; exists {
				total += runeToInt(letter)
				break
			}
		}

	}
	return total, nil
}

func Solve2(inputFilePath string) (int, error) {
	total := 0
	lineNumber := 0

	var firstElf, secondElf map[rune]struct{}
	for fileLine := range helper.FileLineReader(inputFilePath) {
		if lineNumber%3 == 0 {
			firstElf, secondElf = map[rune]struct{}{}, map[rune]struct{}{}
		}

		for _, letter := range fileLine {
			if lineNumber == 0 {
				firstElf[letter] = struct{}{}
			} else if lineNumber == 1 {
				secondElf[letter] = struct{}{}
			} else {
				_, firstExists := firstElf[letter]
				_, secondExists := secondElf[letter]
				if firstExists && secondExists {
					total += runeToInt(letter)
					break
				}
			}
		}
		lineNumber = (lineNumber + 1) % 3
	}
	return total, nil
}

func runeToInt(letter rune) int {
	if letter < 'a' {
		return int(letter-'A') + 27
	}
	return int(letter-'a') + 1
}

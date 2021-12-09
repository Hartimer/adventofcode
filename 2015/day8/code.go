package day8

import (
	"adventofcode/helper"
	"fmt"
	"regexp"
)

var hexRegex = regexp.MustCompile(`\\x[0-9abcdef]{2}`)
var quoteRegex = regexp.MustCompile(`\\"`)
var slashRegex = regexp.MustCompile(`\\\\`)

func Solve1(inputFilePath string) (int, error) {
	count := 0
	for fileLine := range helper.FileLineReader(inputFilePath) {
		count += len(fileLine)
		fileLine = hexRegex.ReplaceAllString(fileLine, "_")
		fileLine = quoteRegex.ReplaceAllString(fileLine, "_")
		fileLine = slashRegex.ReplaceAllString(fileLine, "_")
		count -= len(fileLine) - 2
	}
	return count, nil
}

func Solve2(inputFilePath string) (int, error) {
	count := 0
	for fileLine := range helper.FileLineReader(inputFilePath) {
		count -= len(fileLine)
		count += len(fmt.Sprintf("%q", fileLine))
	}
	return count, nil
}

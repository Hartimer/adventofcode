package day8

import (
	"adventofcode/helper"
	"fmt"
	"sort"
	"strings"

	"github.com/pkg/errors"
)

func Solve1(inputFilePath string) (int, error) {
	uniqueCount := 0
	for fileLine := range helper.FileLineReader(inputFilePath) {
		parts := strings.Split(fileLine, " | ")
		if len(parts) != 2 {
			return 0, errors.New(fmt.Sprintf("Expected %s to have 2 parts but has %d", fileLine, len(parts)))
		}
		for _, resultCode := range strings.Split(parts[1], " ") {
			switch len(resultCode) {
			case 2, 3, 4, 7:
				uniqueCount++
			}
		}

	}
	return uniqueCount, nil
}

func Solve2(inputFilePath string) (int, error) {
	sum := 0
	for fileLine := range helper.FileLineReader(inputFilePath) {
		parts := strings.Split(fileLine, " | ")
		if len(parts) != 2 {
			return 0, errors.New(fmt.Sprintf("Expected %s to have 2 parts but has %d", fileLine, len(parts)))
		}

		referenceParts := strings.Split(parts[0], " ")
		codesByNumber := map[int]string{}
		numbersByCode := map[string]int{}
		oneCode := findWithLength(referenceParts, 2)
		codesByNumber[1] = oneCode
		numbersByCode[oneCode] = 1
		sevenCode := findWithLength(referenceParts, 3)
		codesByNumber[7] = sevenCode
		numbersByCode[sevenCode] = 7
		fourCode := findWithLength(referenceParts, 4)
		codesByNumber[4] = fourCode
		numbersByCode[fourCode] = 4
		eightCode := findWithLength(referenceParts, 7)
		codesByNumber[8] = eightCode
		numbersByCode[eightCode] = 8

		twoThreeFive := findAllWithLength(referenceParts, 5)
		if len(twoThreeFive) != 3 {
			return 0, errors.New(fmt.Sprintf("Expected 3 strings with length 5 but have %d (%v from %v)", len(twoThreeFive), twoThreeFive, referenceParts))
		}

		for _, s := range twoThreeFive {
			if commonCharacterCount(s, codesByNumber[1]) == 2 {
				// 3 is the only one that fully overlaps with 1
				codesByNumber[3] = s
				numbersByCode[s] = 3
			} else if commonCharacterCount(s, codesByNumber[4]) == 3 {
				// 5 has 3 characters in common with 4
				codesByNumber[5] = s
				numbersByCode[s] = 5
			} else {
				codesByNumber[2] = s
				numbersByCode[s] = 2
			}
		}
		if !confirmExists(codesByNumber, 2, 3, 5) {
			return 0, errors.New(fmt.Sprintf("Expected 2, 3 and 5 to be in the map %v", codesByNumber))
		}

		zeroSixNine := findAllWithLength(referenceParts, 6)
		if len(twoThreeFive) != 3 {
			return 0, errors.New(fmt.Sprintf("Expected 3 strings with length 6 but have %d (%v from %v)", len(zeroSixNine), zeroSixNine, referenceParts))
		}

		for _, s := range zeroSixNine {
			if commonCharacterCount(s, codesByNumber[1]) == 1 {
				// 6 does not overlap with one
				codesByNumber[6] = s
				numbersByCode[s] = 6
			} else if commonCharacterCount(s, codesByNumber[4]) == len(codesByNumber[4]) {
				// 9 overlaps with 4
				codesByNumber[9] = s
				numbersByCode[s] = 9
			} else {
				codesByNumber[0] = s
				numbersByCode[s] = 0
			}
		}
		if !confirmExists(codesByNumber, 0, 6, 9) {
			return 0, errors.New(fmt.Sprintf("Expected 0, 6 and 9 to be in the map %v", codesByNumber))
		}

		// calculate result
		var n int
		for _, resultCode := range strings.Split(parts[1], " ") {
			code := sortStringCharacters(resultCode)
			numberForCode, exists := numbersByCode[code]
			if !exists {
				return 0, errors.New(fmt.Sprintf("No number exists for code %s", code))
			}
			n = (n * 10) + numberForCode
		}
		sum += n
	}
	return sum, nil
}

func confirmExists(codesByNumber map[int]string, ns ...int) bool {
	for _, n := range ns {
		if _, exists := codesByNumber[n]; !exists {
			return false
		}
	}
	return true
}

func commonCharacterCount(s1, s2 string) int {
	runesSeen := map[rune]int{}
	for _, r := range s1 {
		if _, exists := runesSeen[r]; !exists {
			runesSeen[r] = 0
		}
		runesSeen[r]++
	}
	for _, r := range s2 {
		if _, exists := runesSeen[r]; !exists {
			runesSeen[r] = 0
		}
		runesSeen[r]++
	}
	count := 0
	for _, c := range runesSeen {
		if c == 2 {
			count++
		}
	}
	return count
}

func findWithLength(ref []string, l int) string {
	for _, s := range ref {
		if len(s) == l {
			return sortStringCharacters(s)
		}
	}
	return ""
}

func findAllWithLength(ref []string, l int) []string {
	var result []string
	for _, s := range ref {
		if len(s) == l {
			result = append(result, sortStringCharacters(s))
		}
	}
	return result
}

func sortStringCharacters(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

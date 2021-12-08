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

type display string

func (d display) commonSections(d2 display) int {
	runesSeen := map[rune]int{}
	for _, r := range d {
		if _, exists := runesSeen[r]; !exists {
			runesSeen[r] = 0
		}
		runesSeen[r]++
	}
	for _, r := range d2 {
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

func (d display) overlapsWith(d2 display) bool {
	return d.commonSections(d2) == len(d2)
}

func (d display) equals(d2 display) bool {
	return len(d) == len(d2) && d.overlapsWith(d2)
}

func (d display) sorted() display {
	r := []rune(d)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return display(r)
}

type displayConfig struct {
	codesByNumber map[int]display
	numbersByCode map[display]int
}

func newDisplayConfig() *displayConfig {
	return &displayConfig{
		codesByNumber: map[int]display{},
		numbersByCode: map[display]int{},
	}
}

func (d *displayConfig) addConfig(number int, code display) {
	if _, exists := d.codesByNumber[number]; exists {
		panic(fmt.Sprintf("Number %d with code %s already exists in config", number, code))
	}
	d.codesByNumber[number] = code.sorted()
	d.numbersByCode[code.sorted()] = number
}

func (d *displayConfig) getByNumber(number int) display {
	code, exists := d.codesByNumber[number]
	if !exists {
		panic(fmt.Sprintf("No code exists for %d. %v", number, d.codesByNumber))
	}
	return code
}

func (d *displayConfig) getByCode(code display) int {
	number, exists := d.numbersByCode[code.sorted()]
	if !exists {
		panic(fmt.Sprintf("No number exists for %s. %v", code, d.numbersByCode))
	}
	return number
}

func (d *displayConfig) containsNumbers(numbers ...int) bool {
	for _, number := range numbers {
		if _, exists := d.codesByNumber[number]; !exists {
			return false
		}
	}
	return true
}

func Solve2(inputFilePath string) (int, error) {
	sum := 0
	for fileLine := range helper.FileLineReader(inputFilePath) {
		parts := strings.Split(fileLine, " | ")
		if len(parts) != 2 {
			return 0, errors.New(fmt.Sprintf("Expected %s to have 2 parts but has %d", fileLine, len(parts)))
		}

		config := newDisplayConfig()

		referenceParts := strings.Split(parts[0], " ")
		config.addConfig(1, findWithLength(referenceParts, 2))
		config.addConfig(7, findWithLength(referenceParts, 3))
		config.addConfig(4, findWithLength(referenceParts, 4))
		config.addConfig(8, findWithLength(referenceParts, 7))

		twoThreeFive := findAllWithLength(referenceParts, 5)
		if len(twoThreeFive) != 3 {
			return 0, errors.New(fmt.Sprintf("Expected 3 strings with length 5 but have %d (%v from %v)", len(twoThreeFive), twoThreeFive, referenceParts))
		}

		for _, s := range twoThreeFive {
			if s.overlapsWith(config.getByNumber(1)) {
				// 3 is the only one that fully overlaps with 1
				config.addConfig(3, s)
			} else if s.commonSections(config.getByNumber(4)) == 3 {
				// 5 has 3 characters in common with 4
				config.addConfig(5, s)
			} else {
				config.addConfig(2, s)
			}
		}
		if !config.containsNumbers(2, 3, 5) {
			return 0, errors.New(fmt.Sprintf("Expected 2, 3 and 5 to be in the config %v", config))
		}

		zeroSixNine := findAllWithLength(referenceParts, 6)
		if len(twoThreeFive) != 3 {
			return 0, errors.New(fmt.Sprintf("Expected 3 strings with length 6 but have %d (%v from %v)", len(zeroSixNine), zeroSixNine, referenceParts))
		}

		for _, s := range zeroSixNine {
			if !s.overlapsWith(config.getByNumber(1)) {
				// 6 does not overlap with one
				config.addConfig(6, s)
			} else if s.overlapsWith(config.getByNumber(4)) {
				// 9 overlaps with 4
				config.addConfig(9, s)
			} else {
				config.addConfig(0, s)
			}
		}

		if !config.containsNumbers(0, 6, 9) {
			return 0, errors.New(fmt.Sprintf("Expected 0, 6 and 9 to be in the map %v", config))
		}

		// calculate result
		var n int
		for _, resultCode := range strings.Split(parts[1], " ") {
			n = (n * 10) + config.getByCode(display(resultCode))
		}
		sum += n
	}
	return sum, nil
}

func findWithLength(ref []string, l int) display {
	for _, s := range ref {
		if len(s) == l {
			return display(s)
		}
	}
	return ""
}

func findAllWithLength(ref []string, l int) []display {
	var result []display
	for _, s := range ref {
		if len(s) == l {
			result = append(result, display(s))
		}
	}
	return result
}

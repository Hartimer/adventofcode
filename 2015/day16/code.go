package day16

import (
	"adventofcode/helper"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type aunt map[string]int

const idKey = "ID"

var reference = aunt{
	"children":    3,
	"cats":        7,
	"samoyeds":    2,
	"pomeranians": 3,
	"akitas":      0,
	"vizslas":     0,
	"goldfish":    5,
	"trees":       3,
	"cars":        2,
	"perfumes":    1,
}

func Solve1(inputFilePath string) (int, error) {
	return solve(inputFilePath, straightMatch)
}

func Solve2(inputFilePath string) (int, error) {
	return solve(inputFilePath, matches)
}

func solve(inputFilePath string, comparator func(aunt, string, int) bool) (int, error) {
	var possibleAunts []aunt
	for fileLine := range helper.FileLineReader(inputFilePath) {
		// Sue 1: goldfish: 6, trees: 9, akitas: 0
		idx := strings.Index(fileLine, ": ")
		if idx == -1 {
			return 0, errors.New(fmt.Sprintf("No sue number found in %s", fileLine))
		}

		a := aunt{}
		var err error
		a[idKey], err = strconv.Atoi(strings.Split(fileLine[:idx], " ")[1])
		if err != nil {
			return 0, errors.Wrap(err, "")
		}

		possibleMatch := true
		for _, p := range strings.Split(fileLine[idx+2:], ", ") {
			parts, err := helper.SplitAndCheck(p, ": ", 2)
			if err != nil {
				return 0, errors.Wrap(err, "")
			}
			val, err := strconv.Atoi(parts[1])
			if err != nil {
				return 0, errors.Wrap(err, "")
			}

			if !comparator(reference, parts[0], val) {
				possibleMatch = false
			}

			a[parts[0]] = val
		}
		if possibleMatch {
			possibleAunts = append(possibleAunts, a)
		}
	}
	if len(possibleAunts) != 1 {
		return 0, errors.New(fmt.Sprintf("Expected 1 element in %v", possibleAunts))
	}
	return possibleAunts[0][idKey], nil
}

func straightMatch(referenceAunt aunt, attr string, val int) bool {
	return val == referenceAunt[attr]
}

func matches(referenceAunt aunt, attr string, val int) bool {
	switch attr {
	case "cats", "trees":
		return val >= referenceAunt[attr]
	case "pomeranians", "goldfish":
		return val < referenceAunt[attr]
	default:
		return val == referenceAunt[attr]
	}
}

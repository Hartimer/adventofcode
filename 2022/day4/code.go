package day4

import (
	"adventofcode/helper"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type sectionRange struct {
	firstSection int
	lastSection  int
}

func (s sectionRange) contains(s2 sectionRange) bool {
	return s.firstSection <= s2.firstSection && s.lastSection >= s2.lastSection
}

func (s sectionRange) overlaps(s2 sectionRange) bool {
	return (s2.firstSection >= s.firstSection && s2.firstSection <= s.lastSection) ||
		(s2.lastSection >= s.firstSection && s2.lastSection <= s.lastSection)
}

func parseLine(line string) (sectionRange, sectionRange, error) {
	ranges := strings.Split(line, ",")

	leftSections := strings.Split(ranges[0], "-")
	rightSections := strings.Split(ranges[1], "-")

	var err error
	leftSectionRange := sectionRange{}
	leftSectionRange.firstSection, err = strconv.Atoi(leftSections[0])
	if err != nil {
		return sectionRange{}, sectionRange{}, errors.Wrap(err, "")
	}
	leftSectionRange.lastSection, err = strconv.Atoi(leftSections[1])
	if err != nil {
		return sectionRange{}, sectionRange{}, errors.Wrap(err, "")
	}
	rightSectionRange := sectionRange{}
	rightSectionRange.firstSection, err = strconv.Atoi(rightSections[0])
	if err != nil {
		return sectionRange{}, sectionRange{}, errors.Wrap(err, "")
	}
	rightSectionRange.lastSection, err = strconv.Atoi(rightSections[1])
	if err != nil {
		return sectionRange{}, sectionRange{}, errors.Wrap(err, "")
	}
	return leftSectionRange, rightSectionRange, nil
}

func Solve1(inputFilePath string) (int, error) {
	overlappingSections := 0
	for fileLine := range helper.FileLineReader(inputFilePath) {
		leftSection, rightSection, err := parseLine(fileLine)
		if err != nil {
			return 0, errors.Wrap(err, "")
		}
		if leftSection.contains(rightSection) || rightSection.contains(leftSection) {
			overlappingSections++
		}
	}
	return overlappingSections, nil
}

func Solve2(inputFilePath string) (int, error) {
	overlappingSections := 0
	for fileLine := range helper.FileLineReader(inputFilePath) {
		leftSection, rightSection, err := parseLine(fileLine)
		if err != nil {
			return 0, errors.Wrap(err, "")
		}
		if leftSection.overlaps(rightSection) || rightSection.overlaps(leftSection) {
			overlappingSections++
		}
	}
	return overlappingSections, nil
}

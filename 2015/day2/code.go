package day2

import (
	"adventofcode/helper"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func calculateWrappingPaper(rawDimensions string) (int, int, int, error) {
	dimensionParts := strings.Split(rawDimensions, "x")
	if len(dimensionParts) != 3 {
		return 0, 0, 0, errors.New(fmt.Sprintf("Expected %s to have 3 parts but has %d", rawDimensions, len(dimensionParts)))
	}
	rawL, rawH, rawW := dimensionParts[0], dimensionParts[1], dimensionParts[2]
	l, err := strconv.Atoi(rawL)
	if err != nil {
		return 0, 0, 0, errors.Wrap(err, "")
	}
	h, err := strconv.Atoi(rawH)
	if err != nil {
		return 0, 0, 0, errors.Wrap(err, "")
	}
	w, err := strconv.Atoi(rawW)
	if err != nil {
		return 0, 0, 0, errors.Wrap(err, "")
	}

	return l, h, w, nil
}

func minInt(ns ...int) int {
	result := math.MaxInt
	for _, n := range ns {
		if n < result {
			result = n
		}
	}
	return result
}

func Solve1(inputFilePath string) (int, error) {
	totalArea := 0
	for fileLine := range helper.FileLineReader(inputFilePath) {
		l, h, w, err := calculateWrappingPaper(fileLine)
		if err != nil {
			return 0, errors.Wrap(err, "")
		}
		totalArea += 2*l*w + 2*w*h + 2*h*l + minInt(l*w, w*h, h*l)
	}
	return totalArea, nil
}

func Solve2(inputFilePath string) (int, error) {
	totalRibbon := 0
	for fileLine := range helper.FileLineReader(inputFilePath) {
		l, h, w, err := calculateWrappingPaper(fileLine)
		if err != nil {
			return 0, errors.Wrap(err, "")
		}
		totalRibbon += l*h*w + minInt(2*l+2*w, 2*w+2*h, 2*h+2*l)
	}
	return totalRibbon, nil
}

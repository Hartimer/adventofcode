package day12

import (
	"adventofcode/helper"
	"encoding/json"
	"log"
	"regexp"
	"strconv"

	"github.com/pkg/errors"
)

func handleRawList(rawList string) (int, error) {
	var res []interface{}
	if err := json.Unmarshal([]byte(rawList), &res); err != nil {
		return 0, errors.Wrap(err, "")
	}
	log.Printf("%v", res)
	return countList(res)
}

func handleRawMap(rawMap string) (int, error) {
	res := map[string]interface{}{}
	if err := json.Unmarshal([]byte(rawMap), &res); err != nil {
		return 0, errors.Wrap(err, "")
	}
	return countMap(res)
}

func countList(l []interface{}) (int, error) {
	sum := 0
	for _, element := range l {
		partial, err := countElement(element)
		if err != nil {
			return 0, errors.Wrap(err, "")
		}
		sum += partial
	}
	return sum, nil
}

func countMap(m map[string]interface{}) (int, error) {
	hasRed := false
	sum := 0
	for _, v := range m {
		if v == "red" {
			hasRed = true
			break
		}
		partial, err := countElement(v)
		if err != nil {
			return 0, errors.Wrap(err, "")
		}
		sum += partial
	}
	if hasRed {
		return 0, nil
	}

	return sum, nil
}

func countElement(e interface{}) (int, error) {
	switch val := e.(type) {
	case float64:
		return int(val), nil
	case map[string]interface{}:
		partial, err := countMap(val)
		if err != nil {
			return 0, errors.Wrap(err, "")
		}
		return partial, nil
	case []interface{}:
		partial, err := countList(val)
		if err != nil {
			return 0, errors.Wrap(err, "")
		}
		return partial, nil
	}
	return 0, nil
}

var numberRegex = regexp.MustCompile("-?[0-9]+")

func Solve1(inputFilePath string) (int, error) {
	sum := 0
	for fileLine := range helper.FileLineReader(inputFilePath) {
		matches := numberRegex.FindAllString(fileLine, -1)
		for _, rawNumber := range matches {
			n, err := strconv.Atoi(rawNumber)
			if err != nil {
				return 0, errors.Wrap(err, "")
			}
			sum += n
		}
	}
	return sum, nil
}

func Solve2(inputFilePath string) (int, error) {
	sum := 0
	for fileLine := range helper.FileLineReader(inputFilePath) {
		var partial int
		var err error
		if rune(fileLine[0]) == '[' {
			partial, err = handleRawList(fileLine)
		} else {
			partial, err = handleRawMap(fileLine)
		}
		if err != nil {
			return 0, errors.Wrapf(err, "Failed to parse %s", fileLine)
		}
		sum += partial
	}
	return sum, nil
}

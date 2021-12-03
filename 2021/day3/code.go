package day3

import (
	"adventofcode/helper"
	"fmt"
	"log"
	"strconv"

	"github.com/pkg/errors"
)

const asciiOffset = 48

type position struct {
	zeros int
	ones  int
}

func (p *position) mostCommon() int {
	if p.zeros > p.ones {
		return 0
	}
	return 1
}

func (p *position) leastCommon() int {
	if p.zeros <= p.ones {
		return 0
	}
	return 1
}

func statsFromNumbers(numbers []number) []*position {
	var result []*position
	for idx := range numbers {
		n := numbers[idx]
		if len(result) == 0 {
			result = make([]*position, len(n))
		}
		for nIdx, binaryNumber := range n {
			if result[nIdx] == nil {
				result[nIdx] = &position{}
			}
			if binaryNumber == 1 {
				result[nIdx].ones++
			} else {
				result[nIdx].zeros++
			}
		}
	}
	return result
}

type number []int

func Solve1(inputFilePath string) (int, error) {
	var stats []*position

	for fileLine := range helper.FileLineReader(inputFilePath) {
		bs := []byte(fileLine)
		if len(stats) == 0 {
			stats = make([]*position, len(bs))
		}
		for idx, b := range bs {
			if stats[idx] == nil {
				stats[idx] = &position{}
			}
			if int(b)-asciiOffset == 0 {
				stats[idx].zeros++
			} else {
				stats[idx].ones++
			}
		}
	}

	gammaRateStr := ""
	epsilonRateStr := ""

	for _, p := range stats {
		gammaRateStr += strconv.Itoa(p.mostCommon())
		epsilonRateStr += strconv.Itoa(p.leastCommon())
	}

	gammaRate, err := strconv.ParseInt(gammaRateStr, 2, 64)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	epsilonRate, err := strconv.ParseInt(epsilonRateStr, 2, 64)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	return int(gammaRate * epsilonRate), nil
}

func Solve2(inputFilePath string) (int, error) {
	var stats []*position
	var numbers []number

	for fileLine := range helper.FileLineReader(inputFilePath) {
		bs := []byte(fileLine)
		if len(stats) == 0 {
			stats = make([]*position, len(bs))
		}
		var n number
		for idx, b := range bs {
			if stats[idx] == nil {
				stats[idx] = &position{}
			}
			if int(b)-asciiOffset == 0 {
				stats[idx].zeros++
				n = append(n, 0)
			} else {
				stats[idx].ones++
				n = append(n, 1)
			}
		}
		numbers = append(numbers, n)
	}

	oxygenRatings := numbers
	oxygenStats := statsFromNumbers(numbers)
	co2Ratings := numbers
	co2Stats := statsFromNumbers(numbers)

	for i := 0; len(oxygenRatings) > 1 || len(co2Ratings) > 1; i++ {
		if len(oxygenRatings) > 1 {
			oxygenRatings = filterPositions(stats, oxygenRatings, i, oxygenStats[i].mostCommon())
			oxygenStats = statsFromNumbers(oxygenRatings)
		}
		if len(co2Ratings) > 1 {
			co2Ratings = filterPositions(stats, co2Ratings, i, co2Stats[i].leastCommon())
			co2Stats = statsFromNumbers(co2Ratings)
		}
	}

	if len(oxygenRatings) != 1 {
		return 0, errors.New(fmt.Sprintf("Expected oxygen to have 1 value, it has %v", oxygenRatings))
	}
	if len(co2Ratings) != 1 {
		return 0, errors.New(fmt.Sprintf("Expected CO2 to have 1 value, it has %v", co2Ratings))
	}

	oRating, err := rawBinaryToDecimal(oxygenRatings[0])
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	co2Rating, err := rawBinaryToDecimal(co2Ratings[0])
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	return int(oRating * co2Rating), nil
}

func filterPositions(p []*position, numbers []number, idx int, criteria int) []number {
	var result []number
	for numberIdx := range numbers {
		n := numbers[numberIdx]
		if n[idx] == criteria {
			result = append(result, n)
		}
	}
	return result
}

func rawBinaryToDecimal(n []int) (int64, error) {
	var str string
	for _, i := range n {
		str += fmt.Sprint(i)
	}
	log.Printf("Converting %s", str)
	return strconv.ParseInt(str, 2, 64)
}

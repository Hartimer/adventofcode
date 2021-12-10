package day14

import (
	"adventofcode/helper"
	"math"
	"strconv"

	"github.com/pkg/errors"
)

type reindeer struct {
	kmPerSecond int
	flightTime  int
	restTime    int
}

func newReindeer(rawReindeer string) (reindeer, error) {
	parts, err := helper.SplitAndCheck(rawReindeer, " ", 15)
	if err != nil {
		return reindeer{}, errors.Wrap(err, "")
	}
	d := reindeer{}
	d.kmPerSecond, err = strconv.Atoi(parts[3])
	if err != nil {
		return reindeer{}, errors.Wrap(err, "")
	}
	d.flightTime, err = strconv.Atoi(parts[6])
	if err != nil {
		return reindeer{}, errors.Wrap(err, "")
	}
	d.restTime, err = strconv.Atoi(parts[13])
	if err != nil {
		return reindeer{}, errors.Wrap(err, "")
	}
	return d, nil
}

func (r reindeer) moveForSeconds(s int) int {
	timeWindow := r.flightTime + r.restTime
	fullWindows := s / timeWindow
	remainingSecs := s % timeWindow
	return r.kmPerSecond*r.flightTime*fullWindows + r.kmPerSecond*minInt(remainingSecs, r.flightTime)
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(ns ...int) int {
	max := math.MinInt
	for _, n := range ns {
		if n > max {
			max = n
		}
	}
	return max
}

func Solve1(inputFilePath string, duration int) (int, error) {
	maxDistance := 0
	for fileLine := range helper.FileLineReader(inputFilePath) {
		r, err := newReindeer(fileLine)
		if err != nil {
			return 0, errors.Wrap(err, "")
		}
		if dist := r.moveForSeconds(duration); dist > maxDistance {
			maxDistance = dist
		}

	}
	return maxDistance, nil
}

func Solve2(inputFilePath string, duration int) (int, error) {
	var reindeers []reindeer
	for fileLine := range helper.FileLineReader(inputFilePath) {
		r, err := newReindeer(fileLine)
		if err != nil {
			return 0, errors.Wrap(err, "")
		}
		reindeers = append(reindeers, r)

	}

	scores := make([]int, len(reindeers))

	for i := 1; i <= duration; i++ {
		currentDistance := make([]int, len(reindeers))
		maxDistance := -1
		for idx, r := range reindeers {
			distance := r.moveForSeconds(i)
			if distance > maxDistance {
				maxDistance = distance
			}
			currentDistance[idx] = distance
		}
		for idx := range currentDistance {
			if currentDistance[idx] == maxDistance {
				scores[idx]++
			}
		}
	}

	return maxInt(scores...), nil
}

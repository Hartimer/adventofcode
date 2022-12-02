package day2

import (
	"adventofcode/helper"
	"strings"
)

type Round struct {
	OpponentPlay string
	MyPlay       string
}

func (r Round) Score() int {
	var shapeScore int
	winningScore := 0
	switch r.MyPlay {
	case "X":
		shapeScore = 1
		if r.OpponentPlay == "A" {
			winningScore = 3
		} else if r.OpponentPlay == "C" {
			winningScore = 6
		}
	case "Y":
		shapeScore = 2
		if r.OpponentPlay == "A" {
			winningScore = 6
		} else if r.OpponentPlay == "B" {
			winningScore = 3
		}
	case "Z":
		shapeScore = 3
		if r.OpponentPlay == "B" {
			winningScore = 6
		} else if r.OpponentPlay == "C" {
			winningScore = 3
		}
	}
	return shapeScore + winningScore
}

func (r Round) Score2() int {
	switch r.MyPlay {
	case "X":
		if r.OpponentPlay == "A" {
			return 3
		} else if r.OpponentPlay == "B" {
			return 1
		} else if r.OpponentPlay == "C" {
			return 2
		}
	case "Y":
		if r.OpponentPlay == "A" {
			return 1 + 3
		} else if r.OpponentPlay == "B" {
			return 2 + 3
		} else if r.OpponentPlay == "C" {
			return 3 + 3
		}
	case "Z":
		if r.OpponentPlay == "A" {
			return 2 + 6
		} else if r.OpponentPlay == "B" {
			return 3 + 6
		} else if r.OpponentPlay == "C" {
			return 1 + 6
		}
	}
	return 0
}

func ParseInputs(filename string) []Round {
	result := []Round{}
	for fileLine := range helper.FileLineReader(filename) {
		lineParts := strings.Split(fileLine, " ")
		result = append(result, Round{
			OpponentPlay: lineParts[0],
			MyPlay:       lineParts[1],
		})
	}
	return result
}

func Solve1(rounds []Round) int {
	totalScore := 0

	for _, round := range rounds {
		totalScore += round.Score()
	}
	return totalScore
}

func Solve2(rounds []Round) int {
	totalScore := 0

	for _, round := range rounds {
		totalScore += round.Score2()
	}
	return totalScore
}

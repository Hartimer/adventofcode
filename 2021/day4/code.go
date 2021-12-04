package day4

import (
	"adventofcode/helper"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type cell struct {
	value  string
	marked bool
}

type card [][]*cell

func (c card) mark(rawNumber string) {
	for _, row := range c {
		for _, c := range row {
			if c.value == rawNumber {
				c.marked = true
				return
			}
		}
	}
}

func (c card) isWinner() bool {
	colHasUnmarkedCell := map[int]bool{}
	for _, row := range c {
		hasUnmarkedCell := false
		for colIdx, c := range row {
			if !c.marked {
				colHasUnmarkedCell[colIdx] = true
				hasUnmarkedCell = true
			}
		}
		if !hasUnmarkedCell {
			return true
		}
	}
	for colIdx := 0; colIdx < len(c[0]); colIdx++ {
		if hasUnmarked, exists := colHasUnmarkedCell[colIdx]; !exists || !hasUnmarked {
			return true
		}
	}
	return false
}

func (c card) sumUnmarked() int {
	sum := 0
	for _, row := range c {
		for _, c := range row {
			if !c.marked {
				n, err := strconv.Atoi(c.value)
				if err != nil {
					panic(err)
				}
				sum += n
			}
		}
	}
	return sum
}

func newCard(rawLines []string) card {
	var result card
	for _, rawLine := range rawLines {
		var row []*cell
		for _, rawNumber := range strings.Split(rawLine, " ") {
			if rawNumber == "" {
				continue
			}
			row = append(row, &cell{value: rawNumber})
		}
		result = append(result, row)
	}
	return result
}

func parseInputs(inputFilePath string) ([]string, []card) {
	var inputSections [][]string
	currentSectionIdx := 0
	var currentSection []string
	for fileLine := range helper.FileLineReader(inputFilePath) {
		if len(fileLine) == 0 {
			if len(currentSection) > 0 {
				inputSections = append(inputSections, currentSection)
			}
			currentSectionIdx++
			currentSection = []string{}
			continue
		}
		currentSection = append(currentSection, fileLine)
	}
	if len(currentSection) > 0 {
		inputSections = append(inputSections, currentSection)
	}

	var cards []card

	for _, section := range inputSections[1:] {
		cards = append(cards, newCard(section))
	}
	return strings.Split(inputSections[0][0], ","), cards
}

func Solve1(inputFilePath string) (int, error) {
	callOrder, cards := parseInputs(inputFilePath)
	for _, calledNumber := range callOrder {
		for _, c := range cards {
			c.mark(calledNumber)
			if c.isWinner() {
				calledNumberInt, err := strconv.Atoi(calledNumber)
				if err != nil {
					return 0, errors.Wrap(err, "")
				}
				return c.sumUnmarked() * calledNumberInt, nil
			}
		}
	}

	return 0, errors.New("No solution")
}

func Solve2(inputFilePath string) (int, error) {
	callOrder, cards := parseInputs(inputFilePath)

	obsoleteCards := map[int]struct{}{}

	for _, calledNumber := range callOrder {
		for cardIdx, c := range cards {
			if _, isObsolete := obsoleteCards[cardIdx]; isObsolete {
				continue
			}
			c.mark(calledNumber)
			if c.isWinner() {
				calledNumberInt, err := strconv.Atoi(calledNumber)
				if err != nil {
					return 0, errors.Wrap(err, "")
				}
				if len(cards)-len(obsoleteCards) == 1 {
					return c.sumUnmarked() * calledNumberInt, nil
				}
				obsoleteCards[cardIdx] = struct{}{}
			}
		}
	}

	return 0, errors.New("No solution")
}

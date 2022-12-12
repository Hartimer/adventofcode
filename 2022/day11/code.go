package day11

import (
	"adventofcode/helper"
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type monkeyGroup []*monkey

type monkey struct {
	id                int
	items             []int
	op                func(int) int
	divisibleByNumber int
	monkeyIfTrue      int
	monkeyIfFalse     int
	inspectedItems    int
	worryLevelFunc    func(int) int
}

func (m *monkey) inspectAllItems(group monkeyGroup) {
	for len(m.items) > 0 {
		item := m.items[0]
		m.items = m.items[1:]
		m.inspectItem(item, group)
	}
}

func (m *monkey) inspectItem(item int, group monkeyGroup) {
	m.inspectedItems++
	worryLevel := m.worryLevelFunc(m.op(item))

	if worryLevel%m.divisibleByNumber == 0 {
		group[m.monkeyIfTrue].items = append(group[m.monkeyIfTrue].items, worryLevel)
	} else {
		group[m.monkeyIfFalse].items = append(group[m.monkeyIfFalse].items, worryLevel)
	}
}

func parseOp(rawOp string) func(int) int {
	rawOp = strings.TrimPrefix(rawOp, "Operation: new = ")
	parts := strings.Split(rawOp, " ")

	firstRawNumber, secondRawNumber := parts[0], parts[2]

	numberParser := func(rawNumber string, old int) int {
		if rawNumber == "old" {
			return old
		}
		number, err := strconv.Atoi(rawNumber)
		if err != nil {
			panic(err)
		}
		return number
	}
	switch parts[1] {
	case "*":
		return func(n int) int {
			return numberParser(firstRawNumber, n) * numberParser(secondRawNumber, n)
		}
	case "+":
		return func(n int) int {
			return numberParser(firstRawNumber, n) + numberParser(secondRawNumber, n)
		}
	case "/":
		return func(n int) int {
			return numberParser(firstRawNumber, n) / numberParser(secondRawNumber, n)
		}
	case "-":
		return func(n int) int {
			return numberParser(firstRawNumber, n) - numberParser(secondRawNumber, n)
		}
	default:
		panic("unknown op " + parts[0])
	}
}

func parseMonkeyGroup(inputFilePath string, worryLevelFunc func(int) int) (monkeyGroup, error) {
	group := monkeyGroup{}
	monkeyRawContents := []string{}
	for fileLine := range helper.FileLineReader(inputFilePath) {
		if len(fileLine) == 0 {
			monkeyRawContents = []string{}
			continue
		}

		monkeyRawContents = append(monkeyRawContents, fileLine)
		if len(monkeyRawContents) == 6 {
			var err error
			m := monkey{worryLevelFunc: worryLevelFunc}
			m.id, err = strconv.Atoi(strings.TrimPrefix(strings.TrimSuffix(monkeyRawContents[0], ":"), "Monkey "))
			if err != nil {
				return nil, errors.Wrap(err, "")
			}

			rawItems := strings.TrimPrefix(strings.TrimSpace(monkeyRawContents[1]), "Starting items: ")
			for _, rawItem := range strings.Split(rawItems, ", ") {
				item, err := strconv.Atoi(rawItem)
				if err != nil {
					return nil, errors.Wrap(err, "")
				}
				m.items = append(m.items, item)
			}

			m.op = parseOp(strings.TrimSpace(monkeyRawContents[2]))

			rawDivisible := strings.TrimPrefix(strings.TrimSpace(monkeyRawContents[3]), "Test: divisible by ")
			m.divisibleByNumber, err = strconv.Atoi(rawDivisible)
			if err != nil {
				return nil, errors.Wrap(err, "")
			}

			rawIfTrue := strings.TrimPrefix(strings.TrimSpace(monkeyRawContents[4]), "If true: throw to monkey ")
			m.monkeyIfTrue, err = strconv.Atoi(rawIfTrue)
			if err != nil {
				return nil, errors.Wrap(err, "")
			}
			rawIfFalse := strings.TrimPrefix(strings.TrimSpace(monkeyRawContents[5]), "If false: throw to monkey ")
			m.monkeyIfFalse, err = strconv.Atoi(rawIfFalse)
			if err != nil {
				return nil, errors.Wrap(err, "")
			}
			group = append(group, &m)
		}
	}
	return group, nil
}

func Solve1(inputFilePath string) (int, error) {
	return solve(inputFilePath, func(w int) int {
		return w / 3
	}, 20)
}

func Solve2(inputFilePath string) (int, error) {
	return solve(inputFilePath, func(w int) int {
		return w % 9699690
	}, 10_000)
}

func solve(inputFilePath string, worryLevelFunc func(int) int, rounds int) (int, error) {
	group, err := parseMonkeyGroup(inputFilePath, worryLevelFunc)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	for round := 0; round < rounds; round++ {
		for _, monkey := range group {
			monkey.inspectAllItems(group)
		}
	}

	inspectedItems := make([]int, len(group))
	for monkeyIdx := range group {
		inspectedItems[monkeyIdx] = group[monkeyIdx].inspectedItems
	}
	sort.Ints(inspectedItems)

	return inspectedItems[len(inspectedItems)-2] * inspectedItems[len(inspectedItems)-1], nil
}

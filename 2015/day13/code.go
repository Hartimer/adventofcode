package day13

import (
	"adventofcode/helper"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type person struct {
	name      string
	happiness map[string]int
}

type people map[string]person

func (p people) clone() people {
	m2 := people{}
	for name := range p {
		m2[name] = p[name]
	}
	return m2
}

func parsePeople(inputFilePath string) (people, error) {
	p := people{}
	for fileLine := range helper.FileLineReader(inputFilePath) {
		parts, err := helper.SplitAndCheck(fileLine, " ", 11)
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		name := parts[0]
		target := strings.TrimSuffix(parts[10], ".")
		happiness, err := strconv.Atoi(parts[3])
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		if parts[2] == "lose" {
			happiness *= -1
		}
		if _, exists := p[name]; !exists {
			p[name] = person{name: name, happiness: map[string]int{}}
		}
		p[name].happiness[target] = happiness
	}
	return p, nil
}

func Solve1(inputFilePath string) (int, error) {
	p, err := parsePeople(inputFilePath)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	return solve(p)
}

func Solve2(inputFilePath string) (int, error) {
	p, err := parsePeople(inputFilePath)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	selfName := "zord"
	peopleToAdd := map[string]int{}

	for name := range p {
		p[name].happiness[selfName] = 0
		peopleToAdd[name] = 0
	}
	p[selfName] = person{name: selfName, happiness: peopleToAdd}

	return solve(p)
}

func solve(p people) (int, error) {
	for name := range p {
		first := p[name]
		chain, happy := calc([]person{first}, people{name: first}, p)
		if len(chain) == len(p) {
			return happy, nil
		}
	}
	return 0, errors.New("Couldn't find a seat arrangement")
}

func calc(currentChain []person, seen people, p people) ([]person, int) {
	current := currentChain[len(currentChain)-1]
	if len(currentChain) == len(p) {
		return currentChain, current.happiness[currentChain[0].name] + currentChain[0].happiness[current.name]
	}
	max := 0
	var bestChain []person
	for possiblePerson := range current.happiness {
		_, wasSeen := seen[possiblePerson]
		if wasSeen {
			continue
		}
		nextPerson := p[possiblePerson]

		newSeen := seen.clone()
		newSeen[possiblePerson] = nextPerson
		possibleChain, partial := calc(append(currentChain, nextPerson), newSeen, p)
		partial += current.happiness[possiblePerson] + nextPerson.happiness[current.name]
		if partial > max {
			max = partial
			bestChain = possibleChain
		}
	}
	return bestChain, max
}

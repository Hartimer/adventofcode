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

type table []string

type people map[string]person

func Solve1(inputFilePath string) (int, error) {
	p := people{}
	for fileLine := range helper.FileLineReader(inputFilePath) {
		parts, err := helper.SplitAndCheck(fileLine, " ", 11)
		if err != nil {
			return 0, errors.Wrap(err, "")
		}
		name := parts[0]
		target := strings.TrimSuffix(parts[10], ".")
		happiness, err := strconv.Atoi(parts[3])
		if err != nil {
			return 0, errors.Wrap(err, "")
		}
		if parts[2] == "lose" {
			happiness *= -1
		}
		if _, exists := p[name]; !exists {
			p[name] = person{name: name, happiness: map[string]int{}}
		}
		p[name].happiness[target] = happiness
	}


	panic("not implemented")
}


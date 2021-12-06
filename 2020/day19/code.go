package day19

import (
	"adventofcode/helper"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type rule interface {
	matches(map[int]rule, string) (string, bool)
}

type literalRule struct {
	letter string
}

func (l literalRule) matches(_ map[int]rule, str string) (string, bool) {
	if len(str) == 0 {
		return "", false
	}
	firstLetter := string(str[0])
	if firstLetter == l.letter {
		return strings.TrimPrefix(str, firstLetter), true
	}
	return str, false
}

var _ rule = literalRule{}

type compositeRule struct {
	subRules []rule
}

func (c compositeRule) matches(rules map[int]rule, str string) (string, bool) {
	workingStr := str
	for _, subRule := range c.subRules {
		var valid bool
		workingStr, valid = subRule.matches(rules, workingStr)
		if !valid {
			return str, false
		}
	}
	return workingStr, true
}

var _ rule = compositeRule{}

type orRule struct {
	leftSide  rule
	rightSide rule
}

func (o orRule) matches(rules map[int]rule, str string) (string, bool) {
	if result, valid := o.leftSide.matches(rules, str); valid {
		return result, true
	}
	if result, valid := o.rightSide.matches(rules, str); valid {
		return result, true
	}
	return str, false
}

var _ rule = orRule{}

type referenceRule struct {
	ruleIdx int
}

func (r referenceRule) matches(rules map[int]rule, str string) (string, bool) {
	return rules[r.ruleIdx].matches(rules, str)
}

var _ rule = referenceRule{}

func parseRule(rawRule string) rule {
	if orParts := strings.Split(rawRule, " | "); len(orParts) == 2 {
		return orRule{
			leftSide:  parseRule(orParts[0]),
			rightSide: parseRule(orParts[1]),
		}
	}

	if compositeParts := strings.Split(rawRule, " "); len(compositeParts) > 1 {
		c := compositeRule{}
		for _, compositePart := range compositeParts {
			c.subRules = append(c.subRules, parseRule(compositePart))
		}
		return c
	}

	if rawRule[0] == '"' {
		return literalRule{letter: string(rawRule[1])}
	}

	ruleIdx, err := strconv.Atoi(rawRule)
	if err != nil {
		panic(err)
	}
	return referenceRule{ruleIdx: ruleIdx}
}

func Solve1(inputFilePath string) (int, error) {
	currentRules := map[int]rule{}
	parsingRules := true
	matchCount := 0
	for fileLine := range helper.FileLineReader(inputFilePath) {
		if len(fileLine) == 0 {
			parsingRules = false
			continue
		}

		if parsingRules {
			ruleParts := strings.Split(fileLine, ": ")
			if len(ruleParts) != 2 {
				return 0, errors.New(fmt.Sprintf("Expected %s to have 2 parts but has %d", fileLine, len(ruleParts)))
			}
			ruleIdx, err := strconv.Atoi(ruleParts[0])
			if err != nil {
				return 0, errors.Wrap(err, "")
			}
			currentRules[ruleIdx] = parseRule(ruleParts[1])
		} else if remainingStr, valid := currentRules[0].matches(currentRules, fileLine); valid && len(remainingStr) == 0 {
			matchCount++
		}
	}

	return matchCount, nil
}

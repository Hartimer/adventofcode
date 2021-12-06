package day19

import (
	"adventofcode/helper"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

const verbose = false

func safePrintf(str string, args ...interface{}) {
	if verbose {
		log.Printf(str, args...)
	}
}

type rule interface {
	fmt.Stringer
	matches(map[int]rule, string) (string, bool)
}

type literalRule struct {
	letter string
}

func (l literalRule) matches(_ map[int]rule, str string) (string, bool) {
	if len(str) > 0 && string(str[0]) == l.letter {
		safePrintf("Literal rule %q matched %s. Returning remaining", l.letter, str)
		return str[1:], true
	}
	return "", false
}

func (l literalRule) String() string {
	return fmt.Sprintf("%q", l.letter)
}

var _ rule = literalRule{}

type compositeRule struct {
	subRules []rule
}

func (c compositeRule) matches(rules map[int]rule, str string) (string, bool) {
	workingStr := str
	for ruleIdx, subRule := range c.subRules {
		var valid bool
		workingStr, valid = subRule.matches(rules, workingStr)
		if !valid {
			return str, false
		}
		safePrintf("Composite index %d out of %s is valid", ruleIdx, c)
	}
	return workingStr, true
}

func (c compositeRule) String() string {
	var strs []string
	for _, subRule := range c.subRules {
		strs = append(strs, subRule.String())
	}
	return strings.Join(strs, " ")
}

var _ rule = compositeRule{}

type orRule struct {
	leftSide  rule
	rightSide rule
}

func (o orRule) matches(rules map[int]rule, str string) (string, bool) {
	if result, valid := o.leftSide.matches(rules, str); valid {
		safePrintf("Left side of %s is valid", o)
		return result, true
	}
	if result, valid := o.rightSide.matches(rules, str); valid {
		safePrintf("Right side of %s is valid", o)
		return result, true
	}
	return str, false
}

func (o orRule) String() string {
	return fmt.Sprintf("%s | %s", o.leftSide.String(), o.rightSide.String())
}

var _ rule = orRule{}

type referenceRule struct {
	ruleIdx int
}

func (r referenceRule) matches(rules map[int]rule, str string) (string, bool) {
	result, valid := rules[r.ruleIdx].matches(rules, str)
	if valid {
		safePrintf("Reference rule pointing at %d is valid", r.ruleIdx)
	}
	return result, valid
}

func (r referenceRule) String() string {
	return strconv.Itoa(r.ruleIdx)
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
			log.Printf("%s patches", fileLine)
		}
	}

	return matchCount, nil
}

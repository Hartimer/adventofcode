package day16

import (
	"adventofcode/helper"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type supportedRange struct {
	min int
	max int
}

func (r supportedRange) inRange(val int) bool {
	return val >= r.min && val <= r.max
}

type fieldRule struct {
	name  string
	rules []supportedRange
}

func (f *fieldRule) validFor(val int) bool {
	for _, rule := range f.rules {
		if rule.inRange(val) {
			return true
		}
	}
	return false
}

type ticket []int

func parseTicket(raw string) (ticket, error) {
	var result ticket
	for _, rawNumber := range strings.Split(raw, ",") {
		number, err := strconv.Atoi(rawNumber)
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		result = append(result, number)
	}
	return result, nil
}

func parseRule(line string) (*fieldRule, error) {
	ruleParts := strings.Split(line, ": ")
	if len(ruleParts) != 2 {
		return nil, errors.New(fmt.Sprintf("Expected rule %s to have 2 parts but has %d", line, len(ruleParts)))
	}
	r := &fieldRule{name: ruleParts[0]}

	for _, rawRange := range strings.Split(ruleParts[1], " or ") {
		parsedRange, err := parseRange(rawRange)
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		r.rules = append(r.rules, parsedRange)
	}

	return r, nil
}

func parseRange(raw string) (supportedRange, error) {
	rangeParts := strings.Split(raw, "-")
	if len(rangeParts) != 2 {
		return supportedRange{}, errors.New(fmt.Sprintf("Expected range %s to have 2 parts but has %d", raw, len(rangeParts)))
	}
	min, err := strconv.Atoi(rangeParts[0])
	if err != nil {
		return supportedRange{}, errors.Wrap(err, "")
	}
	max, err := strconv.Atoi(rangeParts[1])
	if err != nil {
		return supportedRange{}, errors.Wrap(err, "")
	}
	return supportedRange{min: int(min), max: int(max)}, nil
}

func Solve1(inputFilePath string) (int, error) {
	rawContents := [3][]string{}
	currentContentIdx := 0
	for fileLine := range helper.FileLineReader(inputFilePath) {
		if len(fileLine) == 0 {
			currentContentIdx++
			continue
		}
		rawContents[currentContentIdx] = append(rawContents[currentContentIdx], fileLine)
	}

	rules := make([]*fieldRule, len(rawContents[0]))
	var err error
	for ruleIdx, rawRule := range rawContents[0] {
		rules[ruleIdx], err = parseRule(rawRule)
		if err != nil {
			return 0, errors.Wrap(err, "")
		}
	}

	errorRate := 0

	for _, rawTicket := range rawContents[2][1:] {
		t, err := parseTicket(rawTicket)
		if err != nil {
			return 0, errors.Wrap(err, "")
		}
		for _, n := range t {
			satisfiesRule := false
			for _, rule := range rules {
				if rule.validFor(n) {
					satisfiesRule = true
					break
				}
			}
			if !satisfiesRule {
				errorRate += n
			}
		}
	}

	return errorRate, nil
}

func Solve2(inputFilePath string) (int, error) {
	rawContents := [3][]string{}
	currentContentIdx := 0
	for fileLine := range helper.FileLineReader(inputFilePath) {
		if len(fileLine) == 0 {
			currentContentIdx++
			continue
		}
		rawContents[currentContentIdx] = append(rawContents[currentContentIdx], fileLine)
	}

	rules := make([]*fieldRule, len(rawContents[0]))
	var err error
	for ruleIdx, rawRule := range rawContents[0] {
		rules[ruleIdx], err = parseRule(rawRule)
		if err != nil {
			return 0, errors.Wrap(err, "")
		}
	}

	possibleRulesByPosition := map[int][]*fieldRule{}

	for _, rawTicket := range rawContents[2][1:] {
		t, err := parseTicket(rawTicket)
		if err != nil {
			return 0, errors.Wrap(err, "")
		}
		matchTicketAndRules(t, possibleRulesByPosition, rules)
	}

	myTicket, err := parseTicket(rawContents[1][1])
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	matchTicketAndRules(myTicket, possibleRulesByPosition, nil)
	cleanupRules(possibleRulesByPosition)
	result := 1
	for fieldIdx, rules := range possibleRulesByPosition {
		if len(rules) != 1 {
			return 0, errors.New(fmt.Sprintf("Field index %d should have 1 rule but has %d", fieldIdx, len(rules)))
		}
		rule := rules[0]
		if strings.HasPrefix(rule.name, "departure") {
			result *= myTicket[fieldIdx]
		}
	}

	return result, nil
}

func matchTicketAndRules(t ticket, possibleRulesByPosition map[int][]*fieldRule, rules []*fieldRule) {
	validTicket := true
	invalidRulesByPosition := map[int][]*fieldRule{}
	for fieldIdx, n := range t {
		satisfiesRule := false
		if _, exists := possibleRulesByPosition[fieldIdx]; !exists {
			possibleRulesByPosition[fieldIdx] = rules
		}
		for ruleIdx := range possibleRulesByPosition[fieldIdx] {
			rule := possibleRulesByPosition[fieldIdx][ruleIdx]
			if rule.validFor(n) {
				satisfiesRule = true
			} else {
				invalidRulesByPosition[fieldIdx] = append(invalidRulesByPosition[fieldIdx], rule)
			}
		}
		if !satisfiesRule {
			validTicket = false
			break
		}
	}
	if validTicket {
		for fieldIdx, invalidRules := range invalidRulesByPosition {
			possibleRulesByPosition[fieldIdx] = removeInvalidRules(possibleRulesByPosition[fieldIdx], invalidRules)
		}
	}
}

func removeInvalidRules(original []*fieldRule, invalid []*fieldRule) []*fieldRule {
	var result []*fieldRule
	for ruleIdx := range original {
		rule := original[ruleIdx]
		valid := true
		for _, invalidRule := range invalid {
			if rule.name == invalidRule.name {
				valid = false
				break
			}
		}
		if valid {
			result = append(result, rule)
		}
	}
	return result
}

func cleanupRules(possibleRulesByPosition map[int][]*fieldRule) {
	for fieldIdx, rules := range possibleRulesByPosition {
		if len(rules) == 1 {
			for fieldIdx2, rules2 := range possibleRulesByPosition {
				if fieldIdx == fieldIdx2 {
					continue
				}
				possibleRulesByPosition[fieldIdx2] = removeInvalidRules(rules2, rules)
			}
			// cleanupRules(possibleRulesByPosition)
		}

	}
}

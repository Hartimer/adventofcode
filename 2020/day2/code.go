package day2

import (
	"adventofcode/helper"
	"fmt"
	"regexp"
	"strconv"

	"github.com/pkg/errors"
)

var policyRegex = regexp.MustCompile(`^(?P<Min>[0-9]+)-(?P<Max>[0-9]+) (?P<Letter>[a-z]): (?P<Password>.*)$`)

type Policy struct {
	min      int
	max      int
	letter   string
	password string
}

func newPolicy(policy string) (*Policy, error) {
	// We assume the following order of sub expression names: "", "Min", "Max", "Letter", "Password"
	policyParts := policyRegex.FindStringSubmatch(policy)
	if len(policyParts) != 5 {
		return nil, errors.New(fmt.Sprintf("Expected 5 policy parts, have %d. Policy was %s", len(policyParts), policy))
	}
	min, err := strconv.ParseInt(policyParts[1], 10, 32)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	max, err := strconv.ParseInt(policyParts[2], 10, 32)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return &Policy{
		min:      int(min),
		max:      int(max),
		letter:   policyParts[3],
		password: policyParts[4],
	}, nil
}

func (p *Policy) isValid() bool {
	letterCount := 0
	for _, char := range p.password {
		if string(char) == p.letter {
			letterCount++
		}
	}
	return letterCount >= p.min && letterCount <= p.max
}

func (p *Policy) isNewValid() bool {
	letterCount := 0
	if len(p.password)+1 < p.min || len(p.password)+1 < p.max {
		return false
	}
	if string(p.password[p.min-1]) == p.letter {
		letterCount++
	}
	if string(p.password[p.max-1]) == p.letter {
		letterCount++
	}

	return letterCount == 1
}

func Solve1(inputFilePath string) (int, error) {
	validPolicies := 0
	for fileLine := range helper.FileLineReader(inputFilePath) {
		policy, err := newPolicy(fileLine)
		if err != nil {
			return 0, errors.Wrap(err, "")
		}
		if policy.isValid() {
			validPolicies++
		}
	}

	return validPolicies, nil
}

func Solve2(inputFilePath string) (int, error) {
	validPolicies := 0
	for fileLine := range helper.FileLineReader(inputFilePath) {
		policy, err := newPolicy(fileLine)
		if err != nil {
			return 0, errors.Wrap(err, "")
		}
		if policy.isNewValid() {
			validPolicies++
		}
	}

	return validPolicies, nil
}

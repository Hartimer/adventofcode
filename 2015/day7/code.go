package day7

import (
	"adventofcode/helper"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type wires map[string]int

type bitwiseGateFunc func(wires) bool

type bitwiseGate struct {
	expr   string
	inputs []string
	f      bitwiseGateFunc
}

func (b bitwiseGate) hasOutput(o string) bool {
	parts := strings.Split(b.expr, " -> ")
	return len(parts) == 2 && parts[1] == o
}

func newGate(expr string, gate bitwiseGateFunc, inputs ...string) bitwiseGate {
	return bitwiseGate{expr: expr, f: gate, inputs: inputs}
}

var digitCheck = regexp.MustCompile(`^[0-9]+$`)

var and = func(r, l, output string) bitwiseGate {
	return newGate(fmt.Sprintf("%s AND %s -> %s", r, l, output), func(w wires) bool {
		var rBits, lBits int
		var exists bool
		if digitCheck.MatchString(r) {
			rBits, _ = strconv.Atoi(r)
		} else if rBits, exists = w[r]; !exists {
			return false

		}
		if digitCheck.MatchString(l) {
			lBits, _ = strconv.Atoi(l)
		} else if lBits, exists = w[l]; !exists {
			return false
		}
		w[output] = rBits & lBits
		return true
	}, r, l)
}

var or = func(r, l, output string) bitwiseGate {
	return newGate(fmt.Sprintf("%s OR %s -> %s", r, l, output), func(w wires) bool {
		var rBits, lBits int
		var exists bool
		if digitCheck.MatchString(r) {
			rBits, _ = strconv.Atoi(r)
		} else if rBits, exists = w[r]; !exists {
			return false

		}
		if digitCheck.MatchString(l) {
			lBits, _ = strconv.Atoi(l)
		} else if lBits, exists = w[l]; !exists {
			return false
		}
		w[output] = rBits | lBits
		return true
	}, r, l)
}

var not = func(input, output string) bitwiseGate {
	return newGate(fmt.Sprintf("NOT %s -> %s", input, output), func(w wires) bool {
		iBits, exists := w[input]
		if !exists {
			return false
		}
		w[output] = ^iBits
		return true
	}, input)
}

var leftShift = func(input string, arg int, output string) bitwiseGate {
	return newGate(fmt.Sprintf("%s << %d -> %s", input, arg, output), func(w wires) bool {
		iBits, exists := w[input]
		if !exists {
			return false
		}
		w[output] = iBits << arg
		return true
	}, input)
}

var rightShift = func(input string, arg int, output string) bitwiseGate {
	return newGate(fmt.Sprintf("%s >> %d -> %s", input, arg, output), func(w wires) bool {
		iBits, exists := w[input]
		if !exists {
			return false
		}
		w[output] = iBits >> arg
		return true
	}, input)
}

var literalConstant = func(arg int, output string) bitwiseGate {
	return newGate(fmt.Sprintf("%d -> %s", arg, output), func(w wires) bool {
		w[output] = arg
		return true
	})
}

var literal = func(arg string, output string) bitwiseGate {
	return newGate(fmt.Sprintf("%s -> %s", arg, output), func(w wires) bool {
		iBits, exists := w[arg]
		if !exists {
			return false
		}
		w[output] = iBits
		return true
	}, arg)
}

func parseExpression(raw string) (bitwiseGate, error) {
	eParts, err := helper.SplitAndCheck(raw, " -> ", 2)
	if err != nil {
		return bitwiseGate{}, errors.Wrap(err, "")
	}

	gateExpr := eParts[0]
	output := eParts[1]
	switch {
	case strings.Contains(gateExpr, "NOT "):
		return not(strings.TrimPrefix(gateExpr, "NOT "), output), nil
	case strings.Contains(gateExpr, " AND "):
		andParts, err := helper.SplitAndCheck(gateExpr, " AND ", 2)
		if err != nil {
			return bitwiseGate{}, errors.Wrap(err, "")
		}
		return and(andParts[0], andParts[1], output), nil
	case strings.Contains(gateExpr, " OR "):
		orParts, err := helper.SplitAndCheck(gateExpr, " OR ", 2)
		if err != nil {
			return bitwiseGate{}, errors.Wrap(err, "")
		}
		return or(orParts[0], orParts[1], output), nil
	case strings.Contains(gateExpr, " LSHIFT "):
		shiftParts, err := helper.SplitAndCheck(gateExpr, " LSHIFT ", 2)
		if err != nil {
			return bitwiseGate{}, errors.Wrap(err, "")
		}
		input := shiftParts[0]
		arg, err := strconv.Atoi(shiftParts[1])
		if err != nil {
			return bitwiseGate{}, errors.Wrap(err, "")
		}
		return leftShift(input, arg, output), nil
	case strings.Contains(gateExpr, " RSHIFT "):
		shiftParts, err := helper.SplitAndCheck(gateExpr, " RSHIFT ", 2)
		if err != nil {
			return bitwiseGate{}, errors.Wrap(err, "")
		}
		input := shiftParts[0]
		arg, err := strconv.Atoi(shiftParts[1])
		if err != nil {
			return bitwiseGate{}, errors.Wrap(err, "")
		}
		return rightShift(input, arg, output), nil
	default:
		input, err := strconv.Atoi(gateExpr)
		if err != nil {
			return literal(gateExpr, output), nil
		}
		return literalConstant(input, output), nil
	}
}

func Solve1(inputFilePath string) (int, error) {
	var gates []bitwiseGate
	for fileLine := range helper.FileLineReader(inputFilePath) {
		g, err := parseExpression(fileLine)
		if err != nil {
			return 0, errors.Wrap(err, "")
		}
		gates = append(gates, g)
	}

	return solve(gates)
}

func Solve2(inputFilePath string) (int, error) {
	var gates []bitwiseGate
	for fileLine := range helper.FileLineReader(inputFilePath) {
		g, err := parseExpression(fileLine)
		if err != nil {
			return 0, errors.Wrap(err, "")
		}
		gates = append(gates, g)
	}

	firstResult, err := solve(gates)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	for gateIdx, g := range gates {
		if g.hasOutput("b") {
			gates = append(gates[:gateIdx], gates[gateIdx+1:]...)
			break
		}
	}

	gates = append(gates, literalConstant(firstResult, "b"))
	return solve(gates)
}

func solve(gates []bitwiseGate) (int, error) {
	w := wires{}
	for len(gates) > 0 {
		var remainingGates []bitwiseGate
		for gIdx := range gates {
			g := gates[gIdx]
			if !g.f(w) {
				remainingGates = append(remainingGates, g)
			} else {
				log.Printf("Applied: %s", g.expr)
			}
		}
		if len(remainingGates) == len(gates) {
			return 0, errors.New(fmt.Sprintf("Gates got stuck with %d left.", len(remainingGates)))
		}
		gates = remainingGates
	}

	return w["a"], nil
}

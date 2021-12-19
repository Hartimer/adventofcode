package day18

import (
	"adventofcode/helper"
	"fmt"
	"math"
	"strconv"

	"github.com/pkg/errors"
)

type Number interface {
	fmt.Stringer
	Magnitude() int
	Add(Number) Number
	AddRight(int) Number
	AddLeft(int) Number
	Explode(int) (Number, int, int, bool)
	Split() (Number, bool)
}

type pair struct {
	left  Number
	right Number
}

func (p pair) Magnitude() int {
	return p.left.Magnitude()*3 + p.right.Magnitude()*2
}

func (p pair) Add(n2 Number) Number {
	return pair{left: p, right: n2}
}

func (p pair) AddRight(n int) Number {
	p.right = p.right.AddLeft(n)
	return p
}

func (p pair) AddLeft(n int) Number {
	p.left = p.left.AddRight(n)
	return p
}

func (p pair) Split() (Number, bool) {
	newNumber, valid := p.left.Split()
	if valid {
		p.left = newNumber
		return p, true
	}
	newNumber, valid = p.right.Split()
	if valid {
		p.right = newNumber
		return p, true
	}
	return p, false
}

func (p pair) Explode(currentDepth int) (Number, int, int, bool) {
	if currentDepth >= 4 {
		l1, ok1 := p.left.(literal)
		l2, ok2 := p.right.(literal)
		if ok1 && ok2 {
			return literal(0), int(l1), int(l2), true
		}
	}
	newNumber, leftVal, rightVal, exploded := p.left.Explode(currentDepth + 1)
	if exploded {
		return pair{left: newNumber, right: p.right.AddLeft(rightVal)}, 0, leftVal, false
	} else if leftVal > 0 {
		return pair{left: p.left.AddLeft(leftVal), right: p.right}, 0, 0, false
	} else if rightVal > 0 {
		return pair{left: p.left, right: p.right.AddRight(rightVal)}, 0, 0, false
	}
	newNumber, leftVal, rightVal, exploded = p.right.Explode(currentDepth + 1)
	if exploded {
		return pair{left: p.left.AddRight(leftVal), right: newNumber}, rightVal, 0, false
	} else if leftVal > 0 {
		return pair{left: p.left.AddLeft(leftVal), right: p.right}, 0, 0, false
	} else if rightVal > 0 {
		return pair{left: p.left, right: p.right.AddRight(rightVal)}, 0, 0, false
	}
	return p, 0, 0, false
}

func (p pair) String() string {
	return fmt.Sprintf("[%s,%s]", p.left, p.right)
}

var _ Number = pair{}

type literal int

func (l literal) Magnitude() int {
	return int(l)
}

func (l literal) Add(n2 Number) Number {
	return pair{left: l, right: n2}
}

func (l literal) AddRight(n2 int) Number {
	return l + literal(n2)
}

func (l literal) AddLeft(n2 int) Number {
	return l + literal(n2)
}

func (l literal) Explode(_ int) (Number, int, int, bool) {
	return l, 0, 0, false
}

func (l literal) Split() (Number, bool) {
	if l >= 10 {
		return pair{
			left:  literal(math.Floor(float64(l) / 2)),
			right: literal(math.Ceil(float64(l) / 2)),
		}, true
	}
	return l, false
}

func (l literal) String() string {
	return fmt.Sprintf("%d", l)
}

func ParseRemaining(str string) (Number, string, error) {
	switch str[0] {
	case '[':
		leftSide, remaining, err := ParseRemaining(str[1:])
		if err != nil {
			return nil, "", errors.Wrap(err, "")
		}
		rightSide, remaining, err := ParseRemaining(remaining)
		if err != nil {
			return nil, "", errors.Wrap(err, "")
		}
		return pair{left: leftSide, right: rightSide}, remaining, nil
	case ']', ',':
		return ParseRemaining(str[1:])
	default:
		n, err := strconv.Atoi(string(str[0]))
		if err != nil {
			return nil, "", errors.Wrap(err, "")
		}
		return literal(n), str[2:], nil
	}

}

func Solve1(inputFilePath string) (int, error) {
	var problemNumber Number
	for fileLine := range helper.FileLineReader(inputFilePath) {
		n, _, err := ParseRemaining(fileLine)
		if err != nil {
			return 0, err
		}
		if problemNumber == nil {
			problemNumber = n
		} else {
			problemNumber = problemNumber.Add(n)
			for {
				newN, _, _, Exploded := problemNumber.Explode(0)
				if Exploded {
					problemNumber = newN
					continue
				}
				newN, Split := problemNumber.Split()
				if Split {
					problemNumber = newN
					continue
				}
				break
			}
		}
	}

	return problemNumber.Magnitude(), nil
}

func Solve2(inputFilePath string) (int, error) {
	// t, err := parseInputs(inputFilePath)
	panic("not implemented")
}

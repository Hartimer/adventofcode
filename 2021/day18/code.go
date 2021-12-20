package day18

import (
	"adventofcode/helper"
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Node struct {
	ID       string
	Val      *int
	Previous *Node
	Children [2]*Node
}

func (n *Node) Add(n2 *Node) *Node {
	newNode := &Node{
		ID:       uuid.NewString(),
		Children: [2]*Node{n, n2},
	}
	n.Previous = newNode
	n2.Previous = newNode
	return newNode
}

func (n *Node) String() string {
	if n.Val != nil {
		return fmt.Sprintf("%d", *n.Val)
	}
	left := n.Children[0].String()
	right := n.Children[1].String()
	return fmt.Sprintf("[%s,%s]", left, right)
}

func (n *Node) Depth() int {
	counter := 0
	for current := n; current != nil; current = current.Previous {
		counter++
	}
	return counter
}

func (n *Node) IsLeaf() bool {
	return n.Val != nil
}

func (n *Node) IsLeftNode(n2 *Node) bool {
	return len(n.Children) > 0 && n.Children[0].ID == n2.ID
}

func (n *Node) IsRightNode(n2 *Node) bool {
	return len(n.Children) > 0 && n.Children[1].ID == n2.ID
}

func (n *Node) addInternal(idx, v int) {
	if n.Val != nil {
		nVal := *n.Val + v
		n.Val = &nVal
		return
	}

	workingNode := n.Children[idx]
	for workingNode != nil {
		if workingNode.IsLeaf() {
			newVal := *workingNode.Val + v
			workingNode.Val = &newVal
			break
		}
		workingNode = workingNode.Children[idx]
	}
}

func (n *Node) AddLeft(v int) {
	n.addInternal(0, v)
}

func (n *Node) AddRight(v int) {
	n.addInternal(1, v)
}

func (n *Node) Explode() bool {
	if n.Val != nil && *n.Val == 7 {
		log.Printf("Lets")
	}
	if n.Depth() > 4 && !n.IsLeaf() && n.Children[0].IsLeaf() && n.Children[1].IsLeaf() {
		zero := 0

		leftVal, rightVal := n.Children[0].Val, n.Children[1].Val
		if n.Previous.IsLeftNode(n) {
			n.Previous.Children[0] = &Node{ID: uuid.NewString(), Val: &zero}
			n.Previous.Children[1].AddLeft(*rightVal)
			current := n.Previous.Previous
			if current.IsLeftNode(n.Previous) {
				for current != nil {
					if current.Previous == nil {
						current = nil
						break
					} else {
						previousCurrent := current
						current = current.Previous
						if !current.IsLeftNode(previousCurrent) {
							break
						}
					}
				}
			}
			if current != nil {
				current.Children[0].AddRight(*leftVal)
			}
		} else {
			n.Previous.Children[1] = &Node{ID: uuid.NewString(), Val: &zero}
			n.Previous.Children[0].AddRight(*leftVal)
			current := n.Previous.Previous
			if current.IsRightNode(n.Previous) {
				for current != nil {
					if current.Previous == nil {
						current = nil
						break
					} else {
						previousCurrent := current
						current = current.Previous
						if !current.IsRightNode(previousCurrent) {
							break
						}
					}
				}
			}
			if current != nil {
				current.Children[1].AddLeft(*rightVal)
			}
		}
		return true
	}
	for _, c := range n.Children {
		if !c.IsLeaf() && c.Explode() {
			return true
		}
	}
	return false
}

func (n *Node) Split() bool {
	if n.Val != nil {
		if *n.Val >= 10 {
			left := int(math.Floor(float64(*n.Val) / 2))
			lNode := &Node{ID: uuid.NewString(), Val: &left}
			right := int(math.Ceil(float64(*n.Val) / 2))
			rNode := &Node{ID: uuid.NewString(), Val: &right}
			n.Children = [2]*Node{lNode, rNode}
			n.Val = nil
			return true
		}
		return false
	}
	if n.Children[0].Split() {
		return true
	}
	if n.Children[1].Split() {
		return true
	}
	return false
}

func ParseRemaining2(str string, previous *Node) (*Node, string, error) {
	switch str[0] {
	case '[':
		current := &Node{ID: uuid.NewString(), Previous: previous}
		leftSide, remaining, err := ParseRemaining2(str[1:], current)
		if err != nil {
			return nil, "", errors.Wrap(err, "")
		}
		rightSide, remaining, err := ParseRemaining2(remaining, current)
		if err != nil {
			return nil, "", errors.Wrap(err, "")
		}
		current.Children = [2]*Node{leftSide, rightSide}
		return current, remaining, nil
	case ']', ',':
		return ParseRemaining2(str[1:], previous)
	default:
		n, err := strconv.Atoi(string(str[0]))
		if err != nil {
			return nil, "", errors.Wrap(err, "")
		}
		return &Node{ID: uuid.NewString(), Val: &n, Previous: previous}, str[2:], nil
	}
}

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

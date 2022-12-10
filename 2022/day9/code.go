package day9

import (
	"adventofcode/helper"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type movement struct {
	direction rune
	// count     int
}

type position struct {
	x, y int
}

type rope []position

func (p position) move(m movement) (position, error) {
	// if m.count == 0 {
	// 	return p, m, fmt.Errorf("trying to move zero steps")
	// }
	// m.count--
	switch m.direction {
	case 'R':
		p.x++
	case 'L':
		p.x--
	case 'U':
		p.y++
	case 'D':
		p.y--
	default:
		return p, fmt.Errorf("unknown direction %v", m.direction)
	}
	return p, nil
}

func (p position) isTouching(p2 position) (bool, []movement) {
	if p.x == p2.x && p.y == p2.y {
		return true, nil
	}

	xDelta := math.Abs(float64(p.x - p2.x))
	yDelta := math.Abs(float64(p.y - p2.y))

	if xDelta+yDelta == 1 || (xDelta == 1 && yDelta == 1) {
		return true, nil
	}

	if p.x == p2.x {
		if p.y < p2.y {
			return false, []movement{{direction: 'D'}}
		}
		return false, []movement{{direction: 'U'}}
	}
	if p.y == p2.y {
		if p.x < p2.x {
			return false, []movement{{direction: 'L'}}
		}
		return false, []movement{{direction: 'R'}}
	}

	movements := []movement{}

	if p.y < p2.y {
		movements = append(movements, movement{direction: 'D'})
	} else {
		movements = append(movements, movement{direction: 'U'})
	}
	if p.x < p2.x {
		movements = append(movements, movement{direction: 'L'})
	} else {
		movements = append(movements, movement{direction: 'R'})
	}
	return false, movements
}

func parseMovements(inputFilePath string) ([]movement, error) {
	movements := []movement{}
	for fileLine := range helper.FileLineReader(inputFilePath) {
		mParts := strings.Split(fileLine, " ")
		m := movement{
			direction: rune(mParts[0][0]),
		}
		count, err := strconv.Atoi(mParts[1])
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		for i := 0; i < count; i++ {
			movements = append(movements, m)
		}
	}
	return movements, nil
}

func Solve1(inputFilePath string) (int, error) {
	return solve(inputFilePath, 2)
}

func Solve2(inputFilePath string) (int, error) {
	return solve(inputFilePath, 10)
}

func solve(inputFilePath string, ropeSize int) (int, error) {
	movements, err := parseMovements(inputFilePath)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	r := make([]position, ropeSize)
	for i := 0; i < ropeSize; i++ {
		r[i] = position{x: 0, y: 0}
	}

	visitedByTail := map[position]struct{}{
		position{x: 0, y: 0}: {},
	}

	for _, headMovement := range movements {
		var err error
		r[0], err = r[0].move(headMovement)
		if err != nil {
			return 0, errors.Wrap(err, "")
		}

		forwardKnot := r[0]
		for knotIdx, knot := range r[1:] {
			newPosition, moved, err := relativeMovement(forwardKnot, knot)
			if err != nil {
				return 0, errors.Wrap(err, "")
			}
			if !moved {
				break
			}
			r[knotIdx+1] = newPosition
			forwardKnot = newPosition
		}
		visitedByTail[r[len(r)-1]] = struct{}{}
	}

	return len(visitedByTail), nil
}

func relativeMovement(p1, p2 position) (position, bool, error) {
	var err error
	isTouching, neededMovements := p1.isTouching(p2)
	if isTouching {
		return p2, false, nil
	}

	for _, tailMovement := range neededMovements {
		p2, err = p2.move(tailMovement)
		if err != nil {
			return p2, false, errors.Wrap(err, "")
		}
	}
	return p2, true, nil
}

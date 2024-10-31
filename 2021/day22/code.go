package day22

import (
	"adventofcode/helper"
	"math"
	"strconv"

	"github.com/pkg/errors"
)

type cellState int

const (
	on cellState = iota
	off
	unknown
)

type coordinate struct {
	x int
	y int
	z int
}

type axis string

const (
	axisX axis = "x"
	axisY axis = "y"
	axisZ axis = "z"
)

type coordinateRange struct {
	from int
	to   int
}

var noRange = coordinateRange{}
var infiniteRange = coordinateRange{from: math.MinInt, to: math.MaxInt}

func (c coordinateRange) trim(c2 coordinateRange) coordinateRange {
	if c2.from > c.to && c2.to < c.from {
		return noRange
	}
	c.from = maxInt(c.from, c2.from)
	c.to = minInt(c.to, c2.to)
	return c
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type cuboid map[axis]coordinateRange

func (c cuboid) contains(c2 coordinate) bool {
	return c2.x >= c[axisX].from &&
		c2.x <= c[axisX].to &&
		c2.y >= c[axisY].from &&
		c2.y <= c[axisY].to &&
		c2.z >= c[axisZ].from &&
		c2.z <= c[axisZ].to
}

type toggleCuboid struct {
	cuboid
	state cellState
}

func (t toggleCuboid) getState(c coordinate) (cellState, bool) {
	if !t.cuboid.contains(c) {
		return 0, false
	}
	return t.state, true
}

type reactor struct {
	cuboids         []toggleCuboid
	dimensions      cuboid
	fixedDimensions bool
}

func (r *reactor) toggle(newState cellState, c cuboid) {
	t := toggleCuboid{cuboid: c, state: newState}
	if !r.fixedDimensions {
		for _, a := range []axis{axisX, axisY, axisZ} {
			newDimension := r.dimensions[a]
			newDimension.from = minInt(r.dimensions[a].from, c[a].from)
			newDimension.to = maxInt(r.dimensions[a].to, c[a].to)
			r.dimensions[a] = newDimension
		}
	}
	r.cuboids = append(r.cuboids, t)
}

func (r *reactor) count(s cellState) int {
	result := 0
	trimmedX := r.dimensions[axisX]
	trimmedY := r.dimensions[axisY]
	trimmedZ := r.dimensions[axisZ]
	for x := trimmedX.from; x <= trimmedX.to; x++ {
		for y := trimmedY.from; y <= trimmedY.to; y++ {
			for z := trimmedZ.from; z <= trimmedZ.to; z++ {
				coord := coordinate{x, y, z}
				coordState := unknown
				for _, c := range r.cuboids {
					if cuboidState, valid := c.getState(coord); valid {
						coordState = cuboidState
					}
				}
				if coordState == s {
					result++
				}
			}
		}
	}
	return result
}

func newReactor(dimensions *cuboid) *reactor {
	r := &reactor{fixedDimensions: false, dimensions: make(cuboid)}
	if dimensions != nil {
		r.dimensions = *dimensions
		r.fixedDimensions = true
	}
	return r
}

func parseInputs(inputFilePath string, dimensions *cuboid) (*reactor, error) {
	r := newReactor(dimensions)
	for fileLine := range helper.FileLineReader(inputFilePath) {
		parts, err := helper.SplitAndCheck(fileLine, " ", 2)
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		var newCellState cellState
		if parts[0] == "on" {
			newCellState = on
		} else {
			newCellState = off
		}

		rangeParts, err := helper.SplitAndCheck(parts[1], ",", 3)
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		c := cuboid{}
		for _, axle := range rangeParts {
			a := axis(axle[0])
			rawRange, err := helper.SplitAndCheck(axle[2:], "..", 2)
			if err != nil {
				return nil, errors.Wrap(err, "")
			}
			start, err := strconv.Atoi(rawRange[0])
			if err != nil {
				return nil, errors.Wrap(err, "")
			}
			end, err := strconv.Atoi(rawRange[1])
			if err != nil {
				return nil, errors.Wrap(err, "")
			}
			c[a] = coordinateRange{from: start, to: end}
		}
		r.toggle(newCellState, c)
	}
	return r, nil
}

func Solve1(inputFilePath string) (int, error) {
	dimensions := cuboid{
		axisX: coordinateRange{from: -50, to: 50},
		axisY: coordinateRange{from: -50, to: 50},
		axisZ: coordinateRange{from: -50, to: 50},
	}
	r, err := parseInputs(inputFilePath, &dimensions)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	return r.count(on), nil
}

func Solve2(inputFilePath string) (int, error) {
	r, err := parseInputs(inputFilePath, nil)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	return r.count(on), nil
}

package day17

import (
	"adventofcode/helper"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type targetArea struct {
	minX int
	maxX int
	minY int
	maxY int
}

func (t targetArea) inArea(x, y int) bool {
	return x >= t.minX && x <= t.maxX && y >= t.minY && y <= t.maxY
}

type probe struct {
	x      int
	y      int
	xSpeed int
	ySpeed int
}

func (p probe) move() probe {
	p.x += p.xSpeed
	p.y += p.ySpeed
	if p.xSpeed > 0 {
		p.xSpeed--
	} else if p.xSpeed < 0 {
		p.xSpeed++
	}
	p.ySpeed--
	return p
}

func (p probe) isMovingTowards(area targetArea) bool {
	if p.x < area.minX && p.xSpeed <= 0 {
		return false
	}
	if p.x > area.maxX && p.xSpeed >= 0 {
		return false
	}

	if p.y < area.minY {
		return false
	}

	return true
}

func (p probe) moveTowards(area targetArea) bool {
	for {
		p = p.move()
		if area.inArea(p.x, p.y) {
			return true
		}
		if !p.isMovingTowards(area) {
			return false
		}
	}
}

func (p probe) highestY() int {
	for currentY := p.y; ; currentY = p.y {
		p = p.move()
		if p.y < currentY {
			return currentY
		}
	}
}

func parseInputs(inputFilePath string) (targetArea, error) {
	fileLine := <-helper.FileLineReader(inputFilePath)
	line := strings.TrimPrefix(fileLine, "target area: ")
	parts, err := helper.SplitAndCheck(line, ", ", 2)
	if err != nil {
		return targetArea{}, errors.Wrap(err, "")
	}
	xParts, err := helper.SplitAndCheck(strings.TrimPrefix(parts[0], "x="), "..", 2)
	if err != nil {
		return targetArea{}, errors.Wrap(err, "")
	}
	yParts, err := helper.SplitAndCheck(strings.TrimPrefix(parts[1], "y="), "..", 2)
	if err != nil {
		return targetArea{}, errors.Wrap(err, "")
	}

	t := targetArea{}
	x1, err := strconv.Atoi(xParts[0])
	if err != nil {
		return targetArea{}, errors.Wrap(err, "")
	}
	x2, err := strconv.Atoi(xParts[1])
	if err != nil {
		return targetArea{}, errors.Wrap(err, "")
	}

	if x1 < x2 {
		t.minX = x1
		t.maxX = x2
	} else {
		t.minX = x2
		t.maxX = x1
	}

	y1, err := strconv.Atoi(yParts[0])
	if err != nil {
		return targetArea{}, errors.Wrap(err, "")
	}
	y2, err := strconv.Atoi(yParts[1])
	if err != nil {
		return targetArea{}, errors.Wrap(err, "")
	}

	if y1 < y2 {
		t.minY = y1
		t.maxY = y2
	} else {
		t.minY = y2
		t.maxY = y1
	}

	return t, nil
}

func Solve1(inputFilePath string) (int, error) {
	t, err := parseInputs(inputFilePath)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	bestY := 0
	for xSpeed := 0; xSpeed < 1000; xSpeed++ {
		for ySpeed := 0; ySpeed < 1000; ySpeed++ {
			p := probe{xSpeed: xSpeed, ySpeed: ySpeed}
			if p.moveTowards(t) {
				highY := p.highestY()
				if highY > bestY {
					bestY = highY
				}
			}
		}
	}
	return bestY, nil
}

func Solve2(inputFilePath string) (int, error) {
	t, err := parseInputs(inputFilePath)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	counter := 0
	for xSpeed := 0; xSpeed < 1000; xSpeed++ {
		for ySpeed := -1000; ySpeed < 1000; ySpeed++ {
			p := probe{xSpeed: xSpeed, ySpeed: ySpeed}
			if p.moveTowards(t) {
				counter++
			}
		}
	}
	return counter, nil
}

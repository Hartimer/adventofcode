package _test

import (
	"adventofcode/helper"
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

type Map8 [][]rune

func (m Map8) get(c Coordinate) (rune, bool) {
	if c.X < 0 || c.X >= len(m) ||
		c.Y < 0 || c.Y >= len(m[c.X]) {
		return 0, false
	}
	return m[c.X][c.Y], true
}

func (m Map8) uniqueAntinodes() int {
	nodesByFreq := map[rune][]Coordinate{}
	for x, row := range m {
		for y, cell := range row {
			if cell == '.' {
				continue
			}
			nodesByFreq[cell] = append(nodesByFreq[cell], Coordinate{
				X: x,
				Y: y,
			})
		}
	}
	antinodes := map[Coordinate]struct{}{}
	for _, antennas := range nodesByFreq {
		for idx, antenna := range antennas {
			var otherAntennas []Coordinate
			otherAntennas = append(otherAntennas, antennas[:idx]...)
			otherAntennas = append(otherAntennas, antennas[idx+1:]...)
			for _, otherAntenna := range otherAntennas {
				xOffset := int(math.Abs(float64(antenna.X) - float64(otherAntenna.X)))
				yOffset := int(math.Abs(float64(antenna.Y) - float64(otherAntenna.Y)))
				var antinodeA, antinodeB Coordinate
				if antenna.X > otherAntenna.X {
					antinodeA.X = antenna.X + xOffset
					antinodeB.X = otherAntenna.X - xOffset
				} else {
					antinodeA.X = antenna.X - xOffset
					antinodeB.X = otherAntenna.X + xOffset
				}
				if antenna.Y > otherAntenna.Y {
					antinodeA.Y = antenna.Y + yOffset
					antinodeB.Y = otherAntenna.Y - yOffset
				} else {
					antinodeA.Y = antenna.Y - yOffset
					antinodeB.Y = otherAntenna.Y + yOffset
				}
				if _, isValid := m.get(antinodeA); isValid {
					antinodes[antinodeA] = struct{}{}
				}
				if _, isValid := m.get(antinodeB); isValid {
					antinodes[antinodeB] = struct{}{}
				}
			}
		}
	}
	return len(antinodes)
}
func (m Map8) uniqueAntinodesInline() int {
	nodesByFreq := map[rune][]Coordinate{}
	for x, row := range m {
		for y, cell := range row {
			if cell == '.' {
				continue
			}
			nodesByFreq[cell] = append(nodesByFreq[cell], Coordinate{
				X: x,
				Y: y,
			})
		}
	}
	antinodes := map[Coordinate]struct{}{}
	for _, antennas := range nodesByFreq {
		for idx, antenna := range antennas {
			var otherAntennas []Coordinate
			otherAntennas = append(otherAntennas, antennas[:idx]...)
			otherAntennas = append(otherAntennas, antennas[idx+1:]...)
			for _, otherAntenna := range otherAntennas {
				antinodes[antenna] = struct{}{}
				antinodes[otherAntenna] = struct{}{}
				xOffset := int(math.Abs(float64(antenna.X) - float64(otherAntenna.X)))
				yOffset := int(math.Abs(float64(antenna.Y) - float64(otherAntenna.Y)))
				isValidA, isValidB := true, true
				round := 1
				for isValidA || isValidB {
					var antinodeA, antinodeB Coordinate
					if antenna.X > otherAntenna.X {
						antinodeA.X = antenna.X + (xOffset * round)
						antinodeB.X = otherAntenna.X - (xOffset * round)
					} else {
						antinodeA.X = antenna.X - (xOffset * round)
						antinodeB.X = otherAntenna.X + (xOffset * round)
					}
					if antenna.Y > otherAntenna.Y {
						antinodeA.Y = antenna.Y + (yOffset * round)
						antinodeB.Y = otherAntenna.Y - (yOffset * round)
					} else {
						antinodeA.Y = antenna.Y - (yOffset * round)
						antinodeB.Y = otherAntenna.Y + (yOffset * round)
					}
					_, isValidA = m.get(antinodeA)
					if isValidA {
						antinodes[antinodeA] = struct{}{}
					}
					_, isValidB = m.get(antinodeB)
					if isValidB {
						antinodes[antinodeB] = struct{}{}
					}
					round++
				}
			}
		}
	}
	return len(antinodes)
}

func TestDay8_1(t *testing.T) {
	inputs := []struct {
		filename string
		expected int
	}{
		{
			filename: "day8.1.1.input",
			expected: 14,
		},
		{
			filename: "day8.1.input",
			expected: 256,
		},
	}

	for _, input := range inputs {
		t.Run(fmt.Sprintf("%s produces %d", input.filename, input.expected), func(t *testing.T) {
			m := Map8{}
			for fileLine := range helper.FileLineReader(input.filename) {
				var row []rune
				for _, cell := range fileLine {
					row = append(row, cell)
				}
				m = append(m, row)
			}
			require.Equal(t, input.expected, m.uniqueAntinodes())
		})
	}
}
func TestDay8_2(t *testing.T) {
	inputs := []struct {
		filename string
		expected int
	}{
		{
			filename: "day8.1.1.input",
			expected: 34,
		},
		{
			filename: "day8.1.input",
			expected: 1005,
		},
	}

	for _, input := range inputs {
		t.Run(fmt.Sprintf("%s produces %d", input.filename, input.expected), func(t *testing.T) {
			m := Map8{}
			for fileLine := range helper.FileLineReader(input.filename) {
				var row []rune
				for _, cell := range fileLine {
					row = append(row, cell)
				}
				m = append(m, row)
			}
			require.Equal(t, input.expected, m.uniqueAntinodesInline())
		})
	}
}

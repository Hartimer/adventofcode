package day3

import (
	"adventofcode/helper"
	"log"
)

type CellContent string

const (
	Empty CellContent = "."
	Tree  CellContent = "#"
)

type slopes [][]CellContent

func (s *slopes) addRow(row string) {
	var newRow []CellContent
	log.Printf("Adding %s", row)
	for _, cellContent := range row {
		newRow = append(newRow, CellContent(cellContent))
	}
	*s = append(*s, newRow)
}

func (s slopes) getCell(column, row int) CellContent {
	relevantRow := s[row%s.height()]
	return relevantRow[column%len(relevantRow)]
}

func (s slopes) height() int {
	return len(s)
}

func (s slopes) treeCounter(horizontalIncrementer, verticalIncrementer int) int {
	treeCount := 0

	for horizontalOffset, verticalOffset := horizontalIncrementer, verticalIncrementer; verticalOffset <= s.height(); horizontalOffset, verticalOffset = horizontalOffset+horizontalIncrementer, verticalOffset+verticalIncrementer {
		cell := s.getCell(horizontalOffset, verticalOffset)
		log.Printf("Cell at (%d, %d) contains %v", horizontalOffset, verticalOffset, cell)
		if cell == Tree {
			treeCount++
		}
	}

	return treeCount
}

func Solve1(inputFilePath string) (int, error) {
	mountain := &slopes{}
	for fileLine := range helper.FileLineReader(inputFilePath) {
		mountain.addRow(fileLine)
	}

	return mountain.treeCounter(3, 1), nil
}

func Solve2(inputFilePath string) (int, error) {
	mountain := &slopes{}
	for fileLine := range helper.FileLineReader(inputFilePath) {
		mountain.addRow(fileLine)
	}

	treeMultiplier := mountain.treeCounter(1, 1)
	treeMultiplier *= mountain.treeCounter(3, 1)
	treeMultiplier *= mountain.treeCounter(5, 1)
	treeMultiplier *= mountain.treeCounter(7, 1)
	treeMultiplier *= mountain.treeCounter(1, 2)
	return treeMultiplier, nil
}

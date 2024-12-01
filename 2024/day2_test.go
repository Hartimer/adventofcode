package _test

import (
	"adventofcode/helper"
	"fmt"
	"testing"
)

func TestDay2_1(t *testing.T) {
	inputs := []struct {
		filename string
		expected int
	}{
		{
			filename: "day2.1.1.input",
			expected: 31,
		},
		{
			filename: "day2.1.input",
			expected: 23177084,
		},
	}
	for _, input := range inputs {
		t.Run(fmt.Sprintf("%s produces %d", input.filename, input.expected), func(t *testing.T) {
			for fileLine := range helper.FileLineReader(input.filename) {
			}
		})
	}
}

func TestDay2_2(t *testing.T) {
	inputs := []struct {
		filename string
		expected int
	}{
		{
			filename: "day2.1.1.input",
			expected: 31,
		},
		{
			filename: "day2.2.input",
			expected: 23177084,
		},
	}
	for _, input := range inputs {
		t.Run(fmt.Sprintf("%s produces %d", input.filename, input.expected), func(t *testing.T) {
			for fileLine := range helper.FileLineReader(input.filename) {
			}
		})
	}
}

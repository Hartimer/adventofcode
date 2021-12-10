package day10

import (
	"adventofcode/helper"
	"sort"
)

func isClosing(s rune) bool {
	switch s {
	case ')', ']', '}', '>':
		return true
	default:
		return false
	}
}

func expectedPair(s rune) rune {
	switch s {
	case ')':
		return '('
	case ']':
		return '['
	case '}':
		return '{'
	case '>':
		return '<'
	case '(':
		return ')'
	case '[':
		return ']'
	case '{':
		return '}'
	case '<':
		return '>'
	default:
		panic("No matching pair for " + string(s))
	}
}

func Solve1(inputFilePath string) (int, error) {
	errorCount := map[rune]int{}
	for fileLine := range helper.FileLineReader(inputFilePath) {
		s := Stack{}
		for _, c := range fileLine {
			if isClosing(c) {
				lastEntered, hasValue := s.Pop()
				if !hasValue {
					break
				}
				if expectedPair(c) != lastEntered {
					if _, exists := errorCount[c]; !exists {
						errorCount[c] = 0
					}
					errorCount[c]++
					break
				}
			} else {
				s.Push(c)
			}
		}
	}

	score := 0
	for c, count := range errorCount {
		switch c {
		case ')':
			score += 3 * count
		case ']':
			score += 57 * count
		case '}':
			score += 1197 * count
		case '>':
			score += 25137 * count
		}
	}
	return score, nil
}

func Solve2(inputFilePath string) (int, error) {
	incompleteSequences := [][]rune{}
	for fileLine := range helper.FileLineReader(inputFilePath) {
		s := Stack{}
		isValid := true
		for _, c := range fileLine {
			if isClosing(c) {
				lastEntered, hasValue := s.Pop()
				if !hasValue {
					break
				}
				if expectedPair(c) != lastEntered {
					isValid = false
					break
				}
			} else {
				s.Push(c)
			}
		}
		if isValid {
			var partial []rune
			for lastEntered, hasValue := s.Pop(); hasValue; lastEntered, hasValue = s.Pop() {
				partial = append(partial, lastEntered)
			}
			incompleteSequences = append(incompleteSequences, partial)
		}
	}
	var scores []int
	for _, seq := range incompleteSequences {
		score := 0
		for _, o := range seq {
			c := expectedPair(o)
			score *= 5
			switch c {
			case ')':
				score += 1
			case ']':
				score += 2
			case '}':
				score += 3
			case '>':
				score += 4
			}
		}
		scores = append(scores, score)
	}
	sort.Ints(scores)
	return scores[len(scores)/2], nil
}

// Stack code was copied from https://www.educative.io/edpresso/how-to-implement-a-stack-in-golang
type Stack []rune

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str rune) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (rune, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

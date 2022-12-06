package day6

import "errors"

type startSequence []byte

func (s startSequence) isValid() bool {
	chars := map[byte]struct{}{}
	for _, char := range s {
		if _, exists := chars[char]; exists {
			return false
		}
		chars[char] = struct{}{}
	}
	return true
}

func (s startSequence) add(b byte) startSequence {
	return append(s[1:], b)
}

func solve(input string, sequenceLength int) (int, error) {
	var ss startSequence = []byte(input[:sequenceLength])
	input = input[sequenceLength:]
	if ss.isValid() {
		return sequenceLength, nil
	}
	for charIndex, letter := range input {
		ss = ss.add(byte(letter))
		if ss.isValid() {
			return sequenceLength + charIndex + 1, nil
		}
	}

	return 0, errors.New("no solution found")
}

func Solve1(input string) (int, error) {
	return solve(input, 4)
}

func Solve2(input string) (int, error) {
	return solve(input, 14)
}

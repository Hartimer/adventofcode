package day11

import (
	"strings"
)

type numberRing struct {
	c int
}

func (n numberRing) increment() (numberRing, bool) {
	return numberRing{c: ((n.c - 'a' + 1) % 26) + 'a'}, rune(n.c) == 'z'
}

func (n numberRing) String() string {
	return string(rune(n.c))
}

type password string

func (p password) increment() password {
	rings := make([]numberRing, len(p))
	for idx, c := range p {
		rings[idx] = numberRing{c: int(c)}
	}
	carries := true
	for i := len(p) - 1; carries && i >= 0; i-- {
		if carries {
			rings[i], carries = rings[i].increment()
		}
	}
	if carries {
		rings = append([]numberRing{{c: int('a')}}, rings...)
	}
	var result password
	for _, r := range rings {
		result += password(r.String())
	}
	return result
}

func (p password) isValid(rules ...passwordRule) bool {
	for _, rule := range rules {
		if !rule(p) {
			return false
		}
	}
	return true
}

type passwordRule func(password) bool

var increasingString passwordRule = func(p password) bool {
	for i := 0; i < len(p)-3; i++ {
		subStr := p[i : i+3]
		if subStr[0]+1 == subStr[1] && subStr[1]+1 == subStr[2] {
			return true
		}
	}
	return false
}

var illegalLetters = []rune{'i', 'o', 'l'}

var noIllegalLetters passwordRule = func(p password) bool {
	for _, c := range p {
		for _, illegalC := range illegalLetters {
			if c == illegalC {
				return false
			}
		}
	}
	return true
}

var differentPairs passwordRule = func(p password) bool {
	charsWithPairs := map[rune]struct{}{}
	for c := 'a'; len(charsWithPairs) < 2 && c <= 'z'; c++ {
		pair := string(c) + string(c)
		idx := strings.Index(string(p), pair)
		if idx != -1 {
			charsWithPairs[c] = struct{}{}
		}
	}
	return len(charsWithPairs) >= 2
}

func skipIllegalLetters(p password, illegalLetters []rune) password {
	rawP := []rune(p)
	for idx, c := range rawP {
		for _, illegalLetter := range illegalLetters {
			if c == illegalLetter {
				rawP[idx] = c + 1
				for i := idx + 1; i < len(rawP); i++ {
					rawP[i] = 'a'
				}
				return password(rawP)
			}
		}
	}
	return p
}

func Solve1(input string) string {
	p := password(input)
	p = skipIllegalLetters(p, illegalLetters)
	for {
		p = skipIllegalLetters(p.increment(), illegalLetters)
		if p.isValid(increasingString, noIllegalLetters, differentPairs) {
			return string(p)
		}
	}
}

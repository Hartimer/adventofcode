package day21

import (
	"adventofcode/helper"
	"fmt"
	"math"
	"strconv"

	"github.com/pkg/errors"
)

type dice interface {
	roll() int
	count() int
}

type deterministicDice struct {
	current   int
	rollCount int
}

func (d *deterministicDice) roll() int {
	d.rollCount++
	if d.current == 100 {
		d.current = 1
	} else {
		d.current++
	}
	return d.current
}

func (d *deterministicDice) count() int {
	return d.rollCount
}

var _ dice = &deterministicDice{}

type board struct {
	positions map[int]int
	points    map[int]int
}

func (b *board) place(player int, startingPosition int) {
	b.positions[player] = startingPosition
	b.points[player] = 0
}

func (b *board) move(player int, spaces int) {
	currentPosition := b.positions[player] + spaces
	for currentPosition > 10 {
		currentPosition -= 10
	}
	b.points[player] += currentPosition
	b.positions[player] = currentPosition
}

func (b *board) playerWithPoints(target int) (int, bool) {
	for playerID, points := range b.points {
		if points >= target {
			return playerID, true
		}
	}
	return 0, false
}

func (b *board) hasPoints(target int) bool {
	for _, points := range b.points {
		if points >= target {
			return true
		}
	}
	return false
}

func (b *board) loserPoints() int {
	min := math.MaxInt
	for _, points := range b.points {
		if points < min {
			min = points
		}
	}
	return min
}

func (b *board) clone() *board {
	b2 := &board{positions: make(map[int]int), points: make(map[int]int)}
	for idx := range b.positions {
		b2.positions[idx] = b.positions[idx]
	}
	for idx := range b.points {
		b2.points[idx] = b.points[idx]
	}
	return b2
}

type board2 struct {
	moves     []int
	positions map[int]int
	points    map[int]int
}

func (b *board2) place(player int, startingPosition int) {
	b.positions[player] = startingPosition
	b.points[player] = 0
}

func (b *board2) lessOrEqual(spaces int) bool {
	for _, i := range b.moves {
		if i > spaces {
			return false
		}
	}
	return true
}

func (b *board2) move(player int, spaces int) {
	b.moves = append(b.moves, spaces)
	currentPosition := b.positions[player] + spaces
	wrapped := false
	for currentPosition > 10 {
		currentPosition -= 10
		wrapped = true
	}
	if wrapped {
		b.points[player] += currentPosition
	}
	b.positions[player] = currentPosition
}

func (b *board2) playerWithPoints(target int) (int, bool) {
	for playerID, points := range b.points {
		if points >= target {
			return playerID, true
		}
	}
	return 0, false
}

func (b *board2) multiplier() int64 {
	permuts := map[string][]int{}
	var rc func([]int, int)
	rc = func(a []int, k int) {
		if k == len(a) {
			permuts[fmt.Sprint(a)] = a
		} else {
			for i := k; i < len(b.moves); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(b.moves, 0)

	return int64(len(permuts))
}

func (b *board2) loserPoints() int {
	min := math.MaxInt
	for _, points := range b.points {
		if points < min {
			min = points
		}
	}
	return min
}

func (b *board2) clone() *board2 {
	m2 := make([]int, len(b.moves))
	copy(m2, b.moves)
	b2 := &board2{positions: make(map[int]int), points: make(map[int]int), moves: m2}
	for idx := range b.positions {
		b2.positions[idx] = b.positions[idx]
	}
	for idx := range b.points {
		b2.points[idx] = b.points[idx]
	}
	return b2
}

func parseInputs(inputFilePath string) (*board, error) {
	b := &board{positions: map[int]int{}, points: map[int]int{}}
	playerID := 0
	for fileLine := range helper.FileLineReader(inputFilePath) {
		parts, err := helper.SplitAndCheck(fileLine, "starting position: ", 2)
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		startingPosition, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		b.place(playerID, startingPosition)
		playerID++
	}
	return b, nil
}

func parseInputs2(inputFilePath string) (*board2, error) {
	b := &board2{positions: map[int]int{}, points: map[int]int{}}
	playerID := 0
	for fileLine := range helper.FileLineReader(inputFilePath) {
		parts, err := helper.SplitAndCheck(fileLine, "starting position: ", 2)
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		startingPosition, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		b.place(playerID, startingPosition)
		playerID++
	}
	return b, nil
}

func Solve1(inputFilePath string) (int, error) {
	b, err := parseInputs(inputFilePath)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	var die dice = &deterministicDice{}

	for playerTurn := 0; !b.hasPoints(1000); playerTurn = (playerTurn + 1) % 2 {
		sum := 0
		for roll := 0; roll < 3; roll++ {
			sum += die.roll()
		}
		b.move(playerTurn, sum)
	}
	return b.loserPoints() * die.count(), nil
}

func Solve2(inputFilePath string) (int64, error) {
	b, err := parseInputs2(inputFilePath)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	boards := []*board2{b}
	winnerUniverses := map[int]int64{0: 0, 1: 0}
	for playerTurn := 0; len(boards) > 0; playerTurn = (playerTurn + 1) % 2 {
		var newBoards []*board2
		for _, b := range boards {
			for dieRoll := 1; dieRoll <= 3; dieRoll++ {
				if !b.lessOrEqual(dieRoll) {
					continue
				}

				extraBoard := b.clone()
				extraBoard.move(playerTurn, dieRoll)
				if playerID, winner := extraBoard.playerWithPoints(21); winner {
					if playerID != playerTurn {
						panic("impossible amirite")
					}
					winnerUniverses[playerTurn] += extraBoard.multiplier()
				} else {
					newBoards = append(newBoards, extraBoard)
				}
			}
		}
		boards = newBoards
	}

	var max int64 = math.MinInt64
	for _, univ := range winnerUniverses {
		if univ > max {
			max = univ
		}
	}
	return max, nil
}

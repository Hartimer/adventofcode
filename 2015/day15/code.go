package day15

import (
	"adventofcode/helper"
	"strconv"

	"github.com/pkg/errors"
)

type ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func newIngredient(raw string) (ingredient, error) {
	parts, err := helper.SplitAndCheck(raw, ": ", 2)
	if err != nil {
		return ingredient{}, errors.Wrap(err, "")
	}
	rawAttrs, err := helper.SplitAndCheck(parts[1], ", ", 5)
	if err != nil {
		return ingredient{}, errors.Wrap(err, "")
	}
	i := ingredient{name: parts[0]}
	i.capacity, err = extractNumber(rawAttrs[0])
	if err != nil {
		return ingredient{}, errors.Wrap(err, "")
	}
	i.durability, err = extractNumber(rawAttrs[1])
	if err != nil {
		return ingredient{}, errors.Wrap(err, "")
	}
	i.flavor, err = extractNumber(rawAttrs[2])
	if err != nil {
		return ingredient{}, errors.Wrap(err, "")
	}
	i.texture, err = extractNumber(rawAttrs[3])
	if err != nil {
		return ingredient{}, errors.Wrap(err, "")
	}
	i.calories, err = extractNumber(rawAttrs[4])
	if err != nil {
		return ingredient{}, errors.Wrap(err, "")
	}
	return i, nil
}

func extractNumber(s string) (int, error) {
	parts, err := helper.SplitAndCheck(s, " ", 2)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	return strconv.Atoi(parts[1])
}

func (i ingredient) dosage(teaspoons int) dose {
	r := i
	r.capacity = r.capacity * teaspoons
	r.durability = r.durability * teaspoons
	r.flavor = r.flavor * teaspoons
	r.texture = r.texture * teaspoons
	r.calories = r.calories * teaspoons
	return dose(r)
}

type dose ingredient

func Solve1(inputFilePath string) (int, error) {
	return solve(inputFilePath, false)
}

func Solve2(inputFilePath string) (int, error) {
	return solve(inputFilePath, true)
}

func solve(inputFilePath string, caresAboutCalories bool) (int, error) {
	var ingredients []ingredient
	for fileLine := range helper.FileLineReader(inputFilePath) {
		i, err := newIngredient(fileLine)
		if err != nil {
			return 0, errors.Wrap(err, "")
		}
		ingredients = append(ingredients, i)
	}
	availableTeaspoons := 100

	return bestScore(availableTeaspoons, []dose{}, ingredients, caresAboutCalories), nil
}

func bestScore(availableTeaspoons int, usedIngredients []dose, remainingIngredients []ingredient, caresAboutCalories bool) int {
	if len(remainingIngredients) == 0 || availableTeaspoons == 0 {
		capacity := 0
		durability := 0
		flavor := 0
		texture := 0
		calories := 0
		for _, d := range usedIngredients {
			capacity += d.capacity
			durability += d.durability
			flavor += d.flavor
			texture += d.texture
			calories += d.calories
		}
		if caresAboutCalories && calories != 500 {
			return 0
		}
		return maxInt(0, capacity) *
			maxInt(0, durability) *
			maxInt(0, flavor) *
			maxInt(0, texture)
	}

	currentIngredient := remainingIngredients[0]

	best := -1
	i := 0
	if len(remainingIngredients) == 1 {
		i = availableTeaspoons
	}
	for ; i <= availableTeaspoons; i++ {
		d := currentIngredient.dosage(i)
		partial := bestScore(availableTeaspoons-i, append(usedIngredients, d), remainingIngredients[1:], caresAboutCalories)
		if partial > best {
			best = partial
		}
	}
	return best
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

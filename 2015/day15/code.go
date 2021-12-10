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

func (i ingredient) score(teaspoons int) int {
	return i.capacity*teaspoons +
		i.durability*teaspoons +
		i.flavor*teaspoons +
		i.texture*teaspoons +
		i.calories*teaspoons
}

func Solve1(inputFilePath string) (int, error) {
	var ingredients []ingredient
	for fileLine := range helper.FileLineReader(inputFilePath) {
		i, err := newIngredient(fileLine)
		if err != nil {
			return 0, errors.Wrap(err, "")
		}
		ingredients = append(ingredients, i)
	}
	availableTeaspoons := 100

	return bestScore(availableTeaspoons, ingredients), nil
}

func bestScore(availableTeaspoons int, ingredients []ingredient) int {
	if len(ingredients) == 0 || availableTeaspoons == 0 {
		return 1
	}

	currentIngredient := ingredients[0]

	best := -1
	for i := 0; i < availableTeaspoons; i++ {
		partial := currentIngredient.score(i)
		partial *= bestScore(availableTeaspoons-i, ingredients[1:])
		if partial > best {
			best = partial
		}
	}
	return best
}

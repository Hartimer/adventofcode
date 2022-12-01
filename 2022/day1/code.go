package day1

import (
	"math"
	"sort"
)

func Solve1(meals [][]int) int {

	maxCalories := math.MinInt

	for _, elfMeals := range meals {
		elfCalories := 0
		for _, meal := range elfMeals {
			elfCalories += meal
		}
		if elfCalories > maxCalories {
			maxCalories = elfCalories
		}
	}

	return maxCalories
}

func Solve2(meals [][]int) int {
	elfTotalCalories := make([]int, len(meals))

	for idx, elfMeals := range meals {
		elfCalories := 0
		for _, meal := range elfMeals {
			elfCalories += meal
		}
		elfTotalCalories[idx] = elfCalories
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elfTotalCalories)))
	topThreeTotal := 0
	for _, elfTotal := range elfTotalCalories[:3] {
		topThreeTotal += elfTotal
	}

	return topThreeTotal
}

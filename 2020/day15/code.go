package day15

func Solve1(input []int) int {
	spokenNumbers := map[int]int{}
	for idx, i := range input {
		spokenNumbers[i] = idx
	}
	nextNumber := 0
	for i := len(input); i < 2019; i++ {
		previousIdx, wasSpoken := spokenNumbers[nextNumber]
		spokenNumbers[nextNumber] = i
		if wasSpoken {
			nextNumber = i - previousIdx
		} else {
			nextNumber = 0
		}
	}
	return nextNumber
}

func Solve2(input []int) int {
	spokenNumbers := map[int]int{}
	for idx, i := range input {
		spokenNumbers[i] = idx
	}
	nextNumber := 0
	for i := len(input); i < 30000000-1; i++ {
		previousIdx, wasSpoken := spokenNumbers[nextNumber]
		spokenNumbers[nextNumber] = i
		if wasSpoken {
			nextNumber = i - previousIdx
		} else {
			nextNumber = 0
		}
	}
	return nextNumber
}

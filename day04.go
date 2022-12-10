package main

import (
	"fmt"
	"strings"
)

func dayFourTaskOneAndTwo(input []string) (int, int) {
	var (
		resOne int
		resTwo int
	)

	for _, line := range input {
		sectors := strings.Split(line, ",")

		firstSector := strings.Split(sectors[0], "-")
		secondSector := strings.Split(sectors[1], "-")

		firstRangeValueMin := mustInt(firstSector[0])
		firstRangeValueMax := mustInt(firstSector[1])
		secondRangeValueMin := mustInt(secondSector[0])
		secondRangeValueMax := mustInt(secondSector[1])

		if (firstRangeValueMin <= secondRangeValueMin && firstRangeValueMax >= secondRangeValueMax) ||
			(secondRangeValueMin <= firstRangeValueMin && secondRangeValueMax >= firstRangeValueMax) {
			resOne++
		}

		if (firstRangeValueMax >= secondRangeValueMin && firstRangeValueMax <= secondRangeValueMax) ||
			(secondRangeValueMax >= firstRangeValueMin && secondRangeValueMax <= firstRangeValueMax) {
			resTwo++
		}
	}

	return resOne, resTwo
}

func DayFour() {
	input := inputLines("input/4.txt")
	resultOne, resultTwo := dayFourTaskOneAndTwo(input)
	fmt.Printf("### DAY 4 ###\n 1: %d\n 2: %d\n", resultOne, resultTwo)
}

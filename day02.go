package main

import (
	"fmt"
)

var (
	scoreResult = map[string]int{
		"A X": 3,
		"A Y": 4,
		"A Z": 8,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 2,
		"C Y": 6,
		"C Z": 7,
	}

	scores = map[string]int{
		"A X": 4,
		"A Y": 8,
		"A Z": 3,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 7,
		"C Y": 2,
		"C Z": 6,
	}
)

func dayTwoTaskOne(input []string) int {
	var score int

	for _, line := range input {
		score += scores[line]
	}

	return score
}

func dayTwoTaskTwo(input []string) int {
	var score int

	for _, line := range input {
		score += scoreResult[line]
	}

	return score
}

func DayTwo() {
	input := inputLines("input/2.txt")
	fmt.Printf("### DAY 2 ###\n 1: %d\n 2: %d\n", dayTwoTaskOne(input), dayTwoTaskTwo(input))
}

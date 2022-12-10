package main

import (
	"fmt"
	"sort"
)

func dayOneTaskOne(input []string) int {
	var (
		maxNum     int
		currentNum int
	)

	for _, line := range input {
		if line == "" {
			if currentNum > maxNum {
				maxNum = currentNum
			}

			currentNum = 0
			continue
		}

		currentNum += mustInt(line)
	}

	return maxNum
}

func dayOneTaskTwo(input []string) int {
	var (
		currentNum int
		sum        int
		nums       []int
	)

	for _, line := range input {
		if line == "" {
			nums = append(nums, currentNum)

			currentNum = 0
			continue
		}

		currentNum += mustInt(line)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	for _, i := range nums[:3] {
		sum += i
	}

	return sum
}

func DayOne() {
	input := inputLines("input/1.txt")
	fmt.Printf("### DAY 1 ###\n 1: %d\n 2: %d\n", dayOneTaskOne(input), dayOneTaskTwo(input))
}

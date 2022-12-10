package main

import (
	"fmt"
	"strings"
)

func dayFiveTaskOne(input []string, stacks [][]string) string {
	var res string

	for _, line := range input {

		parts := strings.Split(line, " ")

		qty := mustInt(parts[1])
		fromIdx := mustInt(parts[3]) - 1
		toIdx := mustInt(parts[5]) - 1

		stack := stacks[fromIdx]
		for i := 0; i < qty; i++ {
			item := stack[len(stack)-1]

			stacks[fromIdx] = stack[:len(stack)-1]
			stacks[toIdx] = append(stacks[toIdx], item)

			stack = stacks[fromIdx]
		}
	}

	for _, i := range stacks {
		res += i[len(i)-1]
	}

	return res
}

func dayFiveTaskTwo(input []string, stacks [][]string) string {
	var res string

	for _, line := range input {

		parts := strings.Split(line, " ")

		qty := mustInt(parts[1])
		fromIdx := mustInt(parts[3]) - 1
		toIdx := mustInt(parts[5]) - 1

		stack := stacks[fromIdx]
		items := stack[len(stack)-qty:]

		stacks[fromIdx] = stack[:len(stack)-qty]
		stacks[toIdx] = append(stacks[toIdx], items...)
	}

	for _, i := range stacks {
		res += i[len(i)-1]
	}

	return res
}

func DayFive() {
	stacksOne := [][]string{
		{"R", "P", "C", "D", "B", "G"},
		{"H", "V", "G"},
		{"N", "S", "Q", "D", "J", "P", "M"},
		{"P", "S", "L", "G", "D", "C", "N", "M"},
		{"J", "B", "N", "C", "P", "F", "L", "S"},
		{"Q", "B", "D", "Z", "V", "G", "T", "S"},
		{"B", "Z", "M", "H", "F", "T", "Q"},
		{"C", "M", "D", "B", "F"},
		{"F", "C", "Q", "G"},
	}

	stacksTwo := make([][]string, len(stacksOne))
	copy(stacksTwo, stacksOne)

	input := inputLines("input/5.txt")
	fmt.Printf("### DAY 5 ###\n 1: %s\n 2: %s\n", dayFiveTaskOne(input, stacksOne), dayFiveTaskTwo(input, stacksTwo))
}

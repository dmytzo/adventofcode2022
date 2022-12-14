package main

import (
	"fmt"
	"sort"
	"strings"
)

type monkey struct {
	items          []int
	operation      func(int) int
	test           func(int)
	inspectedCount int
}

const lcm = 9699690 // TODO: calculate

func dayElevenTaskOne(input []string) int {
	var monkeys []*monkey
	instructionLen := 6

	for i := 0; i < len(input); i += instructionLen + 1 {
		instruction := input[i : i+instructionLen]

		startItemsLineRaw := strings.Split(instruction[1], ": ")[1]
		startItemsRaw := strings.Split(startItemsLineRaw, ", ")

		m := &monkey{}
		for _, rawItem := range startItemsRaw {
			m.items = append(m.items, mustInt(rawItem))
		}

		operationRawLine := strings.Split(instruction[2], "= ")[1]
		operationArgs := strings.Split(operationRawLine, " ")

		m.operation = func(v int) int {
			var (
				firstNum  int
				secondNum int
			)

			if operationArgs[0] == "old" {
				firstNum = v
			} else {
				firstNum = mustInt(operationArgs[0])
			}

			if operationArgs[2] == "old" {
				secondNum = v
			} else {
				secondNum = mustInt(operationArgs[2])
			}

			switch operationArgs[1] {
			case "*":
				return int((firstNum * secondNum) / 3)
			case "+":
				return int((firstNum + secondNum) / 3)
			}
			return 0
		}

		testRawDivisible := strings.Split(instruction[3], "by ")[1]
		testRawResultTrue := strings.Split(strings.TrimSpace(instruction[4]), "monkey ")[1]
		testRawResultFalse := strings.Split(strings.TrimSpace(instruction[5]), "monkey ")[1]

		m.test = func(v int) {
			if v%mustInt(testRawDivisible) == 0 {
				idx := mustInt(testRawResultTrue)
				monkeys[idx].items = append(monkeys[idx].items, v)
				return
			}

			idx := mustInt(testRawResultFalse)
			monkeys[idx].items = append(monkeys[idx].items, v)
		}

		monkeys = append(monkeys, m)
	}

	for i := 0; i < 20; i++ {
		for _, m := range monkeys {
			for _, item := range m.items {
				m.inspectedCount++
				m.test(m.operation(item))
			}

			m.items = []int{}
		}
	}

	var inspectedCounts []int
	for _, m := range monkeys {
		inspectedCounts = append(inspectedCounts, m.inspectedCount)
	}

	sort.Ints(inspectedCounts)

	return inspectedCounts[len(inspectedCounts)-1] * inspectedCounts[len(inspectedCounts)-2]
}

func dayElevenTaskTwo(input []string) int {
	var monkeys []*monkey
	instructionLen := 6

	for i := 0; i < len(input); i += instructionLen + 1 {
		instruction := input[i : i+instructionLen]

		startItemsLineRaw := strings.Split(instruction[1], ": ")[1]
		startItemsRaw := strings.Split(startItemsLineRaw, ", ")

		m := &monkey{}
		for _, rawItem := range startItemsRaw {
			m.items = append(m.items, mustInt(rawItem))
		}

		operationRawLine := strings.Split(instruction[2], "= ")[1]
		operationArgs := strings.Split(operationRawLine, " ")

		m.operation = func(v int) int {
			var (
				firstNum  int
				secondNum int
			)

			if operationArgs[0] == "old" {
				firstNum = v
			} else {
				firstNum = mustInt(operationArgs[0])
			}

			if operationArgs[2] == "old" {
				secondNum = v
			} else {
				secondNum = mustInt(operationArgs[2])
			}

			switch operationArgs[1] {
			case "*":
				res := firstNum * secondNum
				for res > lcm {
					res -= lcm
				}
				return res
			case "+":
				res := firstNum + secondNum
				for res > lcm {
					res -= lcm
				}
				return res
			}
			return 0
		}

		testRawDivisible := strings.Split(instruction[3], "by ")[1]
		testRawResultTrue := strings.Split(strings.TrimSpace(instruction[4]), "monkey ")[1]
		testRawResultFalse := strings.Split(strings.TrimSpace(instruction[5]), "monkey ")[1]

		m.test = func(v int) {
			if v%mustInt(testRawDivisible) == 0 {
				idx := mustInt(testRawResultTrue)
				monkeys[idx].items = append(monkeys[idx].items, v)
				return
			}

			idx := mustInt(testRawResultFalse)
			monkeys[idx].items = append(monkeys[idx].items, v)
		}

		monkeys = append(monkeys, m)
	}

	for i := 0; i < 10000; i++ {
		for _, m := range monkeys {
			for _, item := range m.items {
				m.inspectedCount++
				m.test(m.operation(item))
			}

			m.items = []int{}
		}
	}

	var inspectedCounts []int
	for _, m := range monkeys {
		inspectedCounts = append(inspectedCounts, m.inspectedCount)
	}

	sort.Ints(inspectedCounts)

	return inspectedCounts[len(inspectedCounts)-1] * inspectedCounts[len(inspectedCounts)-2]
}

func DayEleven() {
	input := inputLines("input/11.txt")
	fmt.Printf("### DAY 11 ###\n 1: %d\n 2: %d\n", dayElevenTaskOne(input), dayElevenTaskTwo(input))
}

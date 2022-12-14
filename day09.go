package main

import (
	"fmt"
	"strings"
)

func move(step, length int, direction string, units []coord, steps map[string]bool) {
	for s := 0; s < step; s++ {
		switch direction {
		case "U":
			units[0].y++
		case "D":
			units[0].y--
		case "R":
			units[0].x++
		case "L":
			units[0].x--
		}

		head := units[0]
		for i := 1; i < length; i++ {
			moveTail(&head, &units[i])
			head = units[i]
		}
		steps[fmt.Sprintf("%d:%d", head.x, head.y)] = true
	}
}

func moveTail(head, tail *coord) {
	yGap := head.y - tail.y
	xGap := head.x - tail.x

	if yGap > 1 {
		tail.y += yGap - 1
		if xGap == 1 || xGap == -1 {
			tail.x += xGap
		}
	}
	if yGap < -1 {
		tail.y += yGap + 1
		if xGap == 1 || xGap == -1 {
			tail.x += xGap
		}
	}

	if xGap > 1 {
		tail.x += xGap - 1
		if yGap == 1 || yGap == -1 {
			tail.y += yGap
		}
	}

	if xGap < -1 {
		tail.x += xGap + 1
		if yGap == 1 || yGap == -1 {
			tail.y += yGap
		}
	}
}

func dayNineTaskOneAndTwo(input []string, num int) int {
	steps := make(map[string]bool)
	units := make([]coord, num)

	for _, line := range input {
		parts := strings.Split(line, " ")

		step := mustInt(parts[1])

		move(step, num, parts[0], units, steps)

	}

	return len(steps)
}

func DayNine() {
	input := inputLines("input/9.txt")
	resultOne, resultTwo := dayNineTaskOneAndTwo(input, 2), dayNineTaskOneAndTwo(input, 10)
	fmt.Printf("### DAY 9 ###\n 1: %d\n 2: %d\n", resultOne, resultTwo)
}

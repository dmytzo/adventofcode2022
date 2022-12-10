package main

import (
	"fmt"
	"strings"
)

func tick(cmd []string, processing bool, x int) (int, bool) {
	switch cmd[0] {
	case "addx":
		if !processing {
			return x, true
		}

		return x + mustInt(cmd[1]), false
	case "noop":
		return x, false
	}

	return x, false
}

func dayTenTaskOne(input []string) int {
	var (
		res           int
		cmd           int
		cmdProcessing bool

		x = 1
	)

	for cycle := 0; cmd < len(input); cycle++ {
		if (cycle+21)%40 == 0 {
			res += x * (cycle + 1)
		}

		x, cmdProcessing = tick(strings.Split(input[cmd], " "), cmdProcessing, x)
		if !cmdProcessing {
			cmd++
		}
	}

	return res
}

func dayTenTaskTwo(input []string) [][]string {
	var (
		pos           int
		y             int
		cmd           int
		cmdProcessing bool

		x = 1
	)

	var mx [][]string
	for i := 0; i < 6; i++ {
		mx = append(mx, make([]string, 40))
	}

	for cycle := 0; cmd < len(input); cycle++ {
		if cycle != 0 && cycle%40 == 0 {
			y++
		}

		pos = cycle - y*40

		if pos == x || pos == x-1 || pos == x+1 {
			mx[y][pos] = "#"
		} else {
			mx[y][pos] = "."
		}

		x, cmdProcessing = tick(strings.Split(input[cmd], " "), cmdProcessing, x)
		if !cmdProcessing {
			cmd++
		}
	}

	return mx
}

func DayTen() {
	input := inputLines("input/10.txt")
	fmt.Printf("### DAY 10 ###\n 1: %d\n 2:\n", dayTenTaskOne(input))
	for _, m := range dayTenTaskTwo(input) {
		fmt.Println(m)
	}
}

package main

import (
	"fmt"
)

func daySixTaskOneAndTwo(input string, num int) int {
	var j int

	for i := num; i < len(input); i++ {
		d := map[string]bool{}
		for _, v := range input[j:i] {
			d[string(v)] = true
		}

		j++

		if len(d) == num {
			return i
		}
	}

	return 0
}

func DaySix() {
	input := string(readFile("input/6.txt"))
	resultOne, resultTwo := daySixTaskOneAndTwo(input, 4), daySixTaskOneAndTwo(input, 14)
	fmt.Printf("### DAY 6 ###\n 1: %d\n 2: %d\n", resultOne, resultTwo)
}

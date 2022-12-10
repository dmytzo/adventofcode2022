package main

import "fmt"

func dayThreeTaskOne(input []string) int {
	var res int

	alphabet := generateAlphabet()

	for _, line := range input {
		res += calculateSharedSum([]string{line[:(len(line))/2], line[(len(line))/2:]}, alphabet)
	}

	return res
}

func dayThreeTaskTwo(input []string) int {
	var res int

	alphabet := generateAlphabet()

	for i := 0; i <= len(input)-3; i += 3 {
		res += calculateSharedSum(input[i:i+3], alphabet)
	}

	return res
}

func calculateSharedSum(lines []string, alphabet map[rune]int) int {
	var res int

	shared := map[rune]int{}
	linesLen := len(lines)

	for idx, line := range lines {
		for _, symbol := range line {
			if shared[symbol] != idx {
				continue
			}

			shared[symbol]++
			if shared[symbol] == linesLen {
				res += alphabet[symbol]
			}
		}
	}

	return res
}

func generateAlphabet() map[rune]int {
	var alphabet = map[rune]int{}
	for ch := 'a'; ch <= 'z'; ch++ {
		alphabet[ch] = len(alphabet) + 1
	}

	for ch := 'A'; ch <= 'Z'; ch++ {
		alphabet[ch] = len(alphabet) + 1
	}

	return alphabet
}

func DayThree() {
	input := inputLines("input/3.txt")
	fmt.Printf("### DAY 3 ###\n 1: %d\n 2: %d\n", dayThreeTaskOne(input), dayThreeTaskTwo(input))
}

package main

import (
	"fmt"
)

func dayEightTaskOne(input []string) int {
	rowLen := len(input[0])
	columnLen := len(input)

	res := rowLen*2 + (columnLen-2)*2

	for i := 1; i < columnLen-1; i++ {
		for j := 1; j < rowLen-1; j++ {

			t := input[i][j]
			hiddenL := false
			hiddenR := false
			for jj := 0; jj < rowLen; jj++ {
				if jj < j {
					if t <= input[i][jj] {
						hiddenL = true
					}
				}
				if jj > j {
					if t <= input[i][jj] {
						hiddenR = true
					}
				}
			}

			hiddenT := false
			hiddenB := false
			for ii := 0; ii < columnLen; ii++ {
				if ii < i {
					if t <= input[ii][j] {
						hiddenT = true
					}
				}
				if ii > i {
					if t <= input[ii][j] {
						hiddenB = true
					}
				}
			}

			if !(hiddenL && hiddenR && hiddenT && hiddenB) {
				res += 1
			}
		}
	}

	return res
}

func dayEightTaskTwo(input []string) int {
	rowLen := len(input[0])
	columnLen := len(input)

	res := 0

	for i := 0; i < columnLen; i++ {
		for j := 0; j < rowLen; j++ {
			score := 1

			t := input[i][j]

			for jj := j - 1; jj >= 0; jj-- {
				if input[i][jj] >= t || jj == 0 {
					score *= j - jj
					break
				}
			}

			for jj := j + 1; jj < rowLen; jj++ {
				if input[i][jj] >= t || jj == rowLen-1 {

					score *= jj - j
					break
				}
			}

			for ii := i - 1; ii >= 0; ii-- {
				if input[ii][j] >= t || ii == 0 {
					score *= i - ii
					break
				}
			}

			for ii := i + 1; ii < columnLen; ii++ {
				if input[ii][j] >= t || ii == columnLen-1 {
					score *= ii - i
					break
				}
			}

			if score > res {
				res = score
			}
		}
	}

	return res
}

func DayEight() {
	input := inputLines("input/8.txt")
	fmt.Printf("### DAY 8 ###\n 1: %d\n 2: %d\n", dayEightTaskOne(input), dayEightTaskTwo(input))
}

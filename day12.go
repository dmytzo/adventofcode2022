package main

import (
	"fmt"
)

func bfs(lines []string, start, end coord) int {
	visited := make(map[coord]int)
	var (
		q []coord
		c coord

		yMax = len(lines) - 1
		xMax = len(lines[0]) - 1
	)

	q = append(q, start)
	visited[start] = 0

	for len(q) > 0 {
		c = q[0]
		q = q[1:]
		v := visited[c]

		for _, n := range getNextCoords(c, xMax, yMax) {
			if int(lines[n.y][n.x])-int(lines[c.y][c.x]) > 1 {
				continue
			}

			if n == end {
				return v + 1
			}

			if _, ok := visited[n]; !ok {
				visited[n] = v + 1
				q = append(q, n)
			}
		}
	}

	return 0
}

func dayTwelveTaskOne(input []string) int {
	var (
		lines []string
		start coord
		end   coord
	)

	for idxY, line := range input {
		lines = append(lines, "")
		for idxX, s := range line {
			if s == 'S' {
				start = coord{y: idxY, x: idxX}
				s = 'a'
			}

			if s == 'E' {
				end = coord{y: idxY, x: idxX}
				s = 'z'
			}

			lines[idxY] += string(s)
		}
	}

	return bfs(lines, start, end)
}

func dayTwelveTaskTwo(input []string) int {
	var (
		lines  []string
		starts []coord
		end    coord
		minRes int
	)

	for idxY, line := range input {
		lines = append(lines, "")
		for idxX, s := range line {
			if s == 'S' {
				s = 'a'
			}

			if s == 'a' {
				starts = append(starts, coord{y: idxY, x: idxX})
			}

			if s == 'E' {
				end = coord{y: idxY, x: idxX}
				s = 'z'
			}

			lines[idxY] += string(s)
		}
	}

	for _, start := range starts {
		if res := bfs(lines, start, end); (res < minRes && res != 0) || minRes == 0 {
			minRes = res
		}
	}

	return minRes
}

func getNextCoords(c coord, maxX, maxY int) []coord {
	var coords []coord
	if c.x < maxX {
		coords = append(coords, coord{x: c.x + 1, y: c.y})
	}
	if c.x >= 1 {
		coords = append(coords, coord{x: c.x - 1, y: c.y})
	}

	if c.y < maxY {
		coords = append(coords, coord{x: c.x, y: c.y + 1})
	}
	if c.y >= 1 {
		coords = append(coords, coord{x: c.x, y: c.y - 1})
	}

	return coords
}

func DayTwelve() {
	input := inputLines("input/12.txt")
	fmt.Printf("### DAY 12 ###\n 1: %d\n 2: %d\n", dayTwelveTaskOne(input), dayTwelveTaskTwo(input))
}

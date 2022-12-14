package main

import (
	"fmt"
	"log"

	"go4.org/sort"
)

const (
	unknownT = iota
	floatT
	sliceT
)

const (
	notOkResult = iota
	okResult
	continueResult
)

var (
	floatAndFloat = fmt.Sprintf("%d:%d", floatT, floatT)
	floatAndSlice = fmt.Sprintf("%d:%d", floatT, sliceT)
	sliceAndSlice = fmt.Sprintf("%d:%d", sliceT, sliceT)
	sliceAndFloat = fmt.Sprintf("%d:%d", sliceT, floatT)
)

func getT(v any) int {
	switch v.(type) {
	case float64:
		return floatT
	case []any:
		return sliceT
	}

	log.Fatalf("unknown type")
	return unknownT
}

func compare(x, y any) bool {
	return compareValues(x, y) == okResult
}

func compareValues(x, y any) int {
	switch fmt.Sprintf("%d:%d", getT(x), getT(y)) {
	case floatAndFloat:
		xV, yV := x.(float64), y.(float64)
		if xV < yV {
			return okResult
		}

		if xV == yV {
			return continueResult
		}

		return notOkResult
	case sliceAndSlice:
		xV, yV := x.([]any), y.([]any)

		for idx, v := range xV {
			if idx >= len(yV) {
				return notOkResult
			}

			res := compareValues(v, yV[idx])
			if res == continueResult {
				continue
			}

			return res
		}

		return okResult

	case floatAndSlice:
		return compareValues([]any{x}, y)
	case sliceAndFloat:
		return compareValues(x, []any{y})
	}

	return okResult
}

func dayThirteenTaskOne(input []string) int {
	var packetPairs [][]any
	for i := 0; i < len(input); i += 3 {
		var (
			p1 []any
			p2 []any
		)
		mustUnmarshalJSON([]byte(input[i]), &p1)
		mustUnmarshalJSON([]byte(input[i+1]), &p2)

		packetPairs = append(packetPairs, []any{p1, p2})
	}

	var res int
	for idx, pair := range packetPairs {
		if compare(pair[0], pair[1]) {
			res += idx + 1
		}
	}

	return res
}

func dayThirteenTaskTwo(input []string) int {
	packetPairs := []any{
		[]any{[]any{2.0}},
		[]any{[]any{6.0}},
	}

	for _, line := range input {
		if line == "" {
			continue
		}

		var p []any
		mustUnmarshalJSON([]byte(line), &p)
		packetPairs = append(packetPairs, p)
	}

	sort.Slice(packetPairs, func(ix, iy int) bool {
		return compare(packetPairs[ix], packetPairs[iy])
	})

	point1 := mustMarshalJSON([]any{[]any{2}})
	point2 := mustMarshalJSON([]any{[]any{6}})

	var (
		idxPoint1 int
		idxPoint2 int
	)

	for idx, line := range packetPairs {
		v := mustMarshalJSON(line)

		switch v {
		case point1:
			idxPoint1 = idx + 1
		case point2:
			idxPoint2 = idx + 1
		}
	}

	return idxPoint1 * idxPoint2
}

func DayThirteen() {
	input := inputLines("input/13.txt")
	fmt.Printf("### DAY 12 ###\n 1: %d\n 2: %d\n", dayThirteenTaskOne(input), dayThirteenTaskTwo(input))
}

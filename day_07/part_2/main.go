package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var rows []string
	for scanner.Scan() {
		row := scanner.Text()
		rows = append(rows, row)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	coldIx := strings.Index(rows[0], "S")
	timelineCount := timelineCount(0, coldIx, rows, make(map[int]int))

	fmt.Println(timelineCount)
}

func timelineCount(
	rowIdx, colIdx int,
	rows []string,
	memo map[int]int,
) int {
	if rowIdx >= len(rows) || colIdx < 0 || colIdx >= len(rows[0]) {
		return 0
	}

	memoIdx := rowIdx*len(rows[0]) + colIdx
	if val, found := memo[memoIdx]; found {
		return val
	}

	switch rows[rowIdx][colIdx] {
	case 'S':
		memo[memoIdx] = 1 +
			timelineCount(rowIdx+1, colIdx, rows, memo)
	case '.':
		memo[memoIdx] = timelineCount(rowIdx+1, colIdx, rows, memo)
	case '^':
		memo[memoIdx] = 1 +
			timelineCount(rowIdx+1, colIdx-1, rows, memo) +
			timelineCount(rowIdx+1, colIdx+1, rows, memo)
	default:
		return 0
	}

	return memo[memoIdx]
}

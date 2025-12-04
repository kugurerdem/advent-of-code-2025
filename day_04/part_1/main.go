package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fileStr := strings.TrimSpace(string(bytes))

	rows := strings.Split(fileStr, "\n")
	rowNumber := len(rows)
	colNumber := len(rows[0])

	count := 0
	for rowIdx, row := range rows {
		for colIdx, col := range row {
			if col == '@' {
				neighborCount := countNeighbours(
					rows,
					colIdx, rowIdx,
					colNumber, rowNumber,
				)

				if neighborCount < 4 {
					count++
				}

			}
		}
	}

	fmt.Println(count)
}

func countNeighbours(rows []string, colIdx, rowIdx, colNumber, rowNumber int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			neighbourColIdx := colIdx + i
			neighbourRowIdx := rowIdx + j

			if neighbourColIdx >= colNumber || neighbourColIdx < 0 {
				continue
			}

			if neighbourRowIdx >= rowNumber || neighbourRowIdx < 0 {
				continue
			}

			neighbour := rows[neighbourRowIdx][neighbourColIdx]
			if neighbour == '@' {
				count++
			}
		}
	}

	return count - 1
}

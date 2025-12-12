package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// First, let's create the cells, both numbers and operations
	cells := [][]string{}
	for scanner.Scan() {
		line := collapseSpaces(strings.TrimSpace(scanner.Text()))
		cells = append(cells, strings.Split(line, " "))
	}

	// Now, let's evaluate each column
	rowNum := len(cells)
	colNum := len(cells[0])
	totalVal := 0
	for i := 0; i < colNum; i++ {
		operation := cells[rowNum-1][i]    // Get the operation from the last row
		reduceVal := mustAtoi(cells[0][i]) // Start with the first row value
		// Iterate through the remaining rows
		for j := 1; j < rowNum-1; j++ {
			currentVal := mustAtoi(cells[j][i])

			switch operation {
			case "*":
				reduceVal *= currentVal
			case "+":
				reduceVal += currentVal
			}
		}

		totalVal += reduceVal
	}

	// Now scan the numbers and check if they fall within any range
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(totalVal)
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func collapseSpaces(s string) string {
	re := regexp.MustCompile(`\s+`)
	return re.ReplaceAllString(s, " ")
}

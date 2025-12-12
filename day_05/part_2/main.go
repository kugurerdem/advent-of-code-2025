package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	// Scan the ranges first
	ranges := make([][2]int, 0)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		var start, end int
		if line == "" {
			break
		}

		entries := strings.Split(line, "-")
		start = mustAtoi(entries[0])
		end = mustAtoi(entries[1])
		ranges = append(ranges, [2]int{start, end})
	}

	// 1. Sort ranges by their start values
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	// 2. Keep track of the highest end value seen so far
	// or keep that in mind while iterating through the ranges
	elementCount := 0
	highestEnd := -1
	for _, r := range ranges {
		start, end := r[0], r[1]

		if end <= highestEnd {
			// This range is completely covered by previous ranges
			continue
		}

		if start <= highestEnd {
			start = highestEnd + 1
		}

		elementCount += end - start + 1
		highestEnd = end
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(elementCount)
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

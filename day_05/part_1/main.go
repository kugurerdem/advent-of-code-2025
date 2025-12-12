package main

import (
	"bufio"
	"fmt"
	"os"
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

	// Now scan the numbers and check if they fall within any range
	count := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		number := mustAtoi(line)

		for _, r := range ranges {
			if number >= r[0] && number <= r[1] {
				count++
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(count)
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

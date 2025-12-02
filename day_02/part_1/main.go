package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(delimiterSplit)

	sum := 0

	for scanner.Scan() {
		// Read and parse each entry
		entry := strings.TrimSpace(scanner.Text())
		ranges := strings.Split(entry, "-")

		start, err := strconv.Atoi(ranges[0])
		if err != nil {
			panic(err)
		}

		end, err := strconv.Atoi(ranges[1])
		if err != nil {
			panic(err)
		}

		// Determine the half pattern to start with
		var patternHalf string
		if len(ranges[0])%2 == 0 {
			patternHalf = ranges[0][:(len(ranges[0]) / 2)]
		} else {
			patternHalf = "1" + strings.Repeat("0", len(ranges[0])/2)
		}

		// Convert patternHalf to an integer index,
		// we will use this to generate full patterns
		idx, err := strconv.Atoi(patternHalf)
		if err != nil {
			panic(err)
		}

		// Find the minimum idx such that the duplicated pattern is >= start
		num, ok := satisfiesLowerBound(idx, start)
		for !ok {
			idx++
			num, ok = satisfiesLowerBound(idx, start)
		}

		// Now keep finding patterns until we exceed end
		num, ok = satisfiesUpperBound(idx, end)
		for ok {
			sum += num
			idx++
			num, ok = satisfiesUpperBound(idx, end)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(sum)
}

// satisfiesLowerBound checks if the pattern formed by duplicating the string
// representation of idx is greater than or equal to start.
func satisfiesLowerBound(idx int, start int) (int, bool) {
	num := createPattern(idx)
	return num, num >= start
}

// satisfiesUpperBound checks if the pattern formed by duplicating the string
// representation of idx is less than or equal to end.
func satisfiesUpperBound(idx int, end int) (int, bool) {
	num := createPattern(idx)
	return num, num <= end
}

// createPattern creates a number by duplicating the string representation of idx.
func createPattern(idx int) int {
	halfPatternStr := strconv.Itoa(idx)
	patternStr := halfPatternStr + halfPatternStr
	num, err := strconv.Atoi(patternStr)
	if err != nil {
		panic(err)
	}
	return num
}

// delimiterSplit is a custom split function for a Scanner that splits on commas.
func delimiterSplit(data []byte, atEOF bool) (advance int, token []byte, err error) {
	for i := range data {
		if data[i] == ',' {
			if i > 0 {
				return i + 1, data[:i], nil
			}
			return 1, nil, nil
		}
	}

	if atEOF && len(data) > 0 {
		return len(data), data, nil
	}
	return 0, nil, nil
}

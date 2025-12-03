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
	sum := 0

	for scanner.Scan() {
		entry := strings.TrimSpace(scanner.Text())
		var digits [12]rune
		for i := range digits {
			digits[i] = '0'
		}
		length := len(entry)
		for i, ch := range entry {
			for j, digit := range digits {
				if ch > digit && i <= length-(len(digits)-j) {
					digits[j] = ch
					for k := j + 1; k < len(digits); k++ {
						digits[k] = '0'
					}
					break
				}
			}
		}

		numberStr := string(digits[:])
		number, _ := strconv.Atoi(numberStr)
		sum += number
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(sum)
}

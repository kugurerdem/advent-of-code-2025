package main

import (
	"bufio"
	"fmt"
	"os"
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
		var firstDigit, secondDigit int
		length := len(entry)
		for i, ch := range entry {
			digit := int(ch - '0')
			if digit > firstDigit && i < length-1 {
				firstDigit = digit
				secondDigit = 0
			} else if digit > secondDigit {
				secondDigit = digit
			}
		}

		result := firstDigit*10 + secondDigit
		sum += result
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(sum)
}

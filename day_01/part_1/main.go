package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// read the file input
	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	dial := 50
	zeroCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		dialTime, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		switch line[0] {
		case 'L':
			dial = (dial - dialTime + 100) % 100
		case 'R':
			dial = (dial + dialTime) % 100
		}

		if dial == 0 {
			zeroCount++
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(zeroCount)
}

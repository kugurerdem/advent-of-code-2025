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
	var dial int = 50
	var zeroCount uint = 0
	for scanner.Scan() {
		line := scanner.Text()
		move, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		prevDial := dial
		switch line[0] {
		case 'L':
			dial = dial - move
		case 'R':
			dial = dial + move
		}

		if dial <= 0 {
			zeroCount = zeroCount + uint(-dial/100)
			if prevDial > 0 {
				zeroCount++
			}
		} else if dial >= 100 {
			zeroCount = zeroCount + uint(dial/100)
		}

		dial = ((dial % 100) + 100) % 100
		fmt.Printf("Dial: %d\n", dial)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(zeroCount)
}

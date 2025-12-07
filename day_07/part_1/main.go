package main

import (
	"bufio"
	"fmt"
	"os"
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

	beams := make([]bool, len(rows[0])) // represents the positions of the laser beams
	splitCount := 0                     // counts the number of splits
	for _, row := range rows {
		// create a copy of the current beams for the next state
		nextBeams := make([]bool, len(beams))
		for i := range nextBeams {
			nextBeams[i] = beams[i]
		}

		// process the current row
		for i, ch := range row {
			if ch == 'S' {
				nextBeams[i] = true
			} else if ch == '^' && beams[i] {
				splitCount++
				nextBeams[i] = false
				nextBeams[i-1] = true
				nextBeams[i+1] = true
			}

		}
		beams = nextBeams
		// fmt.Println(row, beams)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(splitCount)
}

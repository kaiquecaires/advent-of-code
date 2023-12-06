package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func extractNumbers(row string) []int {
	output := make([]int, 0)
	numbers := strings.Split(
		strings.Trim(
			strings.Split(row, ":")[1],
			" ",
		),
		" ",
	)

	for _, number := range numbers {
		if number == "" {
			continue
		}

		currentNumber, _ := strconv.Atoi(number)
		output = append(output, currentNumber)
	}

	return output
}

func main() {
	pwd, _ := os.Getwd()
	f, _ := os.ReadFile(pwd + "/go/2023/day_06/input.txt")
	rows := strings.Split(string(f), "\n")

	times := extractNumbers(rows[0])
	distances := extractNumbers(rows[1])
	wins := 1

	for i, distance := range distances {
		time := times[i]
		winCount := 0

		for j := 0; j <= time; j++ {
			currentDistance := j * (time - j)
			if currentDistance > distance {
				winCount++
			}
		}

		wins *= winCount
	}

	fmt.Println(wins)
}

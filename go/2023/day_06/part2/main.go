package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func extractNumbers(row string) int {
	completeNumber := ""
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
		completeNumber += number
	}

	output, _ := strconv.Atoi(completeNumber)

	return output
}

func main() {
	pwd, _ := os.Getwd()
	f, _ := os.ReadFile(pwd + "/go/2023/day_06/input.txt")
	rows := strings.Split(string(f), "\n")

	time := extractNumbers(rows[0])
	distance := extractNumbers(rows[1])
	minRange := 0
	maxRange := 0

	for i := 0; i <= time; i++ {
		currentDistance := i * (time - i)
		if currentDistance > distance {
			maxRange = time - i
			break
		}
	}

	for i := time; i > 0; i-- {
		currentDistance := i * (time - i)
		if currentDistance > distance {
			minRange = time - i
			break
		}
	}

	fmt.Println((maxRange - minRange) + 1)
}

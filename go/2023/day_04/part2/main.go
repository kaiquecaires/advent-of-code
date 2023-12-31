package main

import (
	"fmt"
	"os"
	"strings"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func getWinCopies(rows []string, row int) int {
	cardColumns := strings.Split(rows[row], ": ")
	cardValues := strings.Split(cardColumns[1], " | ")
	winningNumbers := strings.Split(cardValues[0], " ")
	ownedNumbers := strings.Split(cardValues[1], " ")
	count := 0

	for _, ownedNumber := range ownedNumbers {
		if ownedNumber != "" && contains(winningNumbers, ownedNumber) {
			count++
		}
	}

	return count
}

func main() {
	pwd, _ := os.Getwd()
	f, _ := os.ReadFile(pwd + "/go/2023/day_04/input.txt")
	rows := strings.Split(string(f), "\n")
	total := 0

	additionalCards := make(map[int]int)

	for row := range rows {
		winCopies := getWinCopies(rows, row)
		total += additionalCards[row] + 1

		sum := additionalCards[row] + 1
		for i := 0; i < winCopies; i++ {
			additionalCards[row+i+1] += sum
		}
	}

	fmt.Print(total)
}

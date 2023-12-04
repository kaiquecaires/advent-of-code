package main

import (
	"fmt"
	"math"
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

func main() {
	pwd, _ := os.Getwd()
	f, _ := os.ReadFile(pwd + "/go/2023/day_04/input.txt")
	rows := strings.Split(string(f), "\n")
	total := 0

	for _, rowContent := range rows {
		cardColumns := strings.Split(rowContent, ": ")
		cardValues := strings.Split(cardColumns[1], " | ")
		winningNumbers := strings.Split(cardValues[0], " ")
		ownedNumbers := strings.Split(cardValues[1], " ")
		count := 0

		for _, ownedNumber := range ownedNumbers {
			if ownedNumber != "" && contains(winningNumbers, ownedNumber) {
				count++
			}
		}

		cardTotal := math.Pow(2, float64(count-1))

		if count == 0 {
			cardTotal = 0
		}

		total += int(cardTotal)
	}

	fmt.Println(total)
}

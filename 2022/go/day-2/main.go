package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func unpack(line string) (string, string) {
	X := strings.Split(line, " ")
	return X[0], X[1]
}

func hasPairInPairs(pair [2]string, pairs [][2]string) bool {
	for _, currentPair := range pairs {
		if currentPair == pair {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.Open("./2022/go/day-2/input.txt")

	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		panic("Error")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	shapes := map[string]string{
		"A": "Rock",
		"B": "Paper",
		"C": "Scissors",
		"X": "Rock",
		"Y": "Paper",
		"Z": "Scissors",
	}
	pointsPerSchape := map[string]int{
		"Rock":     1,
		"Paper":    2,
		"Scissors": 3,
	}

	pointsPerOutcome := map[string]int{
		"Win":  6,
		"Draw": 3,
		"Lose": 0,
	}

	winningPairs := [][2]string{
		{"Paper", "Rock"},
		{"Rock", "Scissors"},
		{"Scissors", "Paper"},
	}

	score := 0

	for scanner.Scan() {
		line := scanner.Text()
		opponentPlay, ourPlay := unpack(line)
		opponentShape := shapes[opponentPlay]
		ourShape := shapes[ourPlay]

		if opponentShape == ourShape {
			score += pointsPerSchape[ourShape] + pointsPerOutcome["Draw"]
		} else if pair := [2]string{opponentShape, ourShape}; hasPairInPairs(pair, winningPairs) {
			score += pointsPerSchape[ourShape] + pointsPerOutcome["Lose"]
		} else {
			score += pointsPerSchape[ourShape] + pointsPerOutcome["Win"]
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file", err)
		panic("Error")
	}

	fmt.Println(score)
}

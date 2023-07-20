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

	winShapes := map[string]string{
		"Rock":     "Paper",
		"Paper":    "Scissors",
		"Scissors": "Rock",
	}

	loseShapes := map[string]string{
		"Rock":     "Scissors",
		"Paper":    "Rock",
		"Scissors": "Paper",
	}

	results := map[string]string{
		"X": "Lose",
		"Y": "Draw",
		"Z": "Win",
	}

	score := 0

	for scanner.Scan() {
		line := scanner.Text()
		opponentPlay, ourPlay := unpack(line)
		expectedResult := results[ourPlay]
		oponentShape := shapes[opponentPlay]

		if expectedResult == "Draw" {
			score += pointsPerSchape[oponentShape] + pointsPerOutcome["Draw"]
		} else if expectedResult == "Lose" {
			loseShape := loseShapes[oponentShape]
			score += pointsPerSchape[loseShape] + pointsPerOutcome["Lose"]
		} else {
			winShape := winShapes[oponentShape]
			score += pointsPerSchape[winShape] + pointsPerOutcome["Win"]
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file", err)
		panic("Error")
	}

	fmt.Println(score)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var maxCubeValues = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func getGameIdAndSets(characters string) (int, []string) {
	splittedCharacters := strings.Split(characters, ": ")

	game := splittedCharacters[0]
	gameId, _ := strconv.Atoi(strings.Split(game, " ")[1])

	cubeResults := splittedCharacters[1]
	sets := strings.Split(cubeResults, "; ")

	return gameId, sets
}

func isValidSet(set string) bool {
	cubes := strings.Split(set, ", ")

	for _, cube := range cubes {
		values := strings.Split(cube, " ")

		cubeQtd, _ := strconv.Atoi(values[0])

		if maxCubeValues[values[1]] < cubeQtd {
			return false
		}
	}

	return true
}

func main() {
	pwd, _ := os.Getwd()
	f, err := os.Open(pwd + "/go/2023/day_02/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	total := 0

	for scanner.Scan() {
		characters := scanner.Text()
		gameId, sets := getGameIdAndSets(characters)

		for _, set := range sets {
			if !isValidSet(set) {
				gameId = 0
			}
		}

		total += gameId
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(total)
}

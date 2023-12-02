package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getGameIdAndSets(characters string) (int, []string) {
	splittedCharacters := strings.Split(characters, ": ")

	game := splittedCharacters[0]
	gameId, _ := strconv.Atoi(strings.Split(game, " ")[1])

	cubeResults := splittedCharacters[1]
	sets := strings.Split(cubeResults, "; ")

	return gameId, sets
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
		_, sets := getGameIdAndSets(characters)

		var maxCubeValues = map[string]int{
			"red":   1,
			"green": 1,
			"blue":  1,
		}

		for _, set := range sets {
			cubes := strings.Split(set, ", ")

			for _, cube := range cubes {
				arrCube := strings.Split(cube, " ")

				cubeValue, _ := strconv.Atoi(arrCube[0])
				cubeColor := arrCube[1]

				if maxCubeValues[cubeColor] < cubeValue {
					maxCubeValues[cubeColor] = cubeValue
				}
			}
		}

		total += maxCubeValues["red"] * maxCubeValues["green"] * maxCubeValues["blue"]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(total)
}

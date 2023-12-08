package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	pwd, _ := os.Getwd()
	f, _ := os.ReadFile(pwd + "/go/2023/day_08/input.txt")
	rows := strings.Split(string(f), "\n")

	threeMap := map[string]int{}
	three := []string{}

	for row := 2; row < len(rows); row++ {
		if rows[row] == "" {
			continue
		}

		columns := strings.Split(rows[row], " = ")
		rootNode := columns[0]

		values := strings.ReplaceAll(columns[1], "(", "")
		values = strings.ReplaceAll(values, ")", "")
		nodes := strings.Split(values, ", ")

		three = append(three, rootNode)
		threeMap[rootNode] = len(three) - 1

		three = append(three, nodes[0])
		three = append(three, nodes[1])
	}

	coordinates := rows[0]
	currentNode := three[threeMap["AAA"]]

	i := 0

	for currentNode != "ZZZ" {
		coordinate := i % len(coordinates)

		if coordinates[coordinate] == 'L' {
			currentIndex := threeMap[currentNode]
			leftNodeIndex := threeMap[three[currentIndex+1]]
			currentNode = three[leftNodeIndex]
		} else {
			currentIndex := threeMap[currentNode]
			rightNodeIndex := threeMap[three[currentIndex+2]]
			currentNode = three[rightNodeIndex]
		}

		i++
	}

	fmt.Println(i)
}

package main

import (
	"fmt"
	"os"
	"strings"
)

type Node struct {
	left  string
	right string
}

func getNumOfSteps(coordinates string, nodes map[string]Node, nodeKey string) int {
	steps := 0
	curr := nodeKey

	for curr[len(curr)-1] != 'Z' {
		coordinate := coordinates[steps%len(coordinates)]

		if coordinate == 'L' {
			curr = nodes[curr].left
		} else {
			curr = nodes[curr].right
		}

		steps++
	}

	return steps
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func main() {
	pwd, _ := os.Getwd()
	f, _ := os.ReadFile(pwd + "/go/2023/day_08/input.txt")
	rows := strings.Split(string(f), "\n")

	nodes := map[string]Node{}
	startingNodes := []string{}

	for row := 2; row < len(rows); row++ {
		if rows[row] == "" {
			continue
		}

		columns := strings.Split(rows[row], " = ")
		rootNode := columns[0]
		values := strings.ReplaceAll(columns[1], "(", "")
		values = strings.ReplaceAll(values, ")", "")
		nodeValues := strings.Split(values, ", ")

		nodes[rootNode] = Node{left: nodeValues[0], right: nodeValues[1]}

		if rootNode[len(rootNode)-1] == 'A' {
			startingNodes = append(startingNodes, rootNode)
		}
	}

	coordinates := rows[0]
	steps := []int{}

	for _, node := range startingNodes {
		steps = append(steps, getNumOfSteps(coordinates, nodes, node))
	}

	fmt.Println(LCM(steps[0], steps[1], steps[2:]...))
}

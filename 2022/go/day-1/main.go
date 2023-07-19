package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var topMostCalories [3]int
	var currentCalories int = 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			pos := -1
			for i, num := range topMostCalories {
				if currentCalories > num {
					pos = i
					break
				}
			}

			if pos != -1 {
				prev := topMostCalories[pos]
				topMostCalories[pos] = currentCalories
				for i := pos + 1; i < len(topMostCalories); i++ {
					temp := topMostCalories[i]
					topMostCalories[i] = prev
					prev = temp
				}
			}

			currentCalories = 0
		} else {
			num, err := strconv.Atoi(line)

			if err == nil {
				currentCalories += num
			} else {
				fmt.Printf("Error converting number: %v", err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file", err)
	}

	fmt.Print(topMostCalories)
	total := 0

	for _, num := range topMostCalories {
		total += num
	}

	fmt.Printf("Total of: %d", total)
}

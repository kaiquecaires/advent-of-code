package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	pwd, _ := os.Getwd()
	f, _ := os.ReadFile(pwd + "/go/2023/day_14/input.txt")
	grid := makeGrid(strings.Split(string(f), "\n"))

	for row := range grid {
		for col := range grid[row] {
			if string(grid[row][col]) != "." {
				continue
			}
			closestRow := getClosesRock(grid, row, col)
			if closestRow == -1 {
				continue
			}
			grid[row][col] = "O"
			grid[closestRow][col] = "."
		}
	}

	total := 0
	for row := range grid {
		total += countRocks(grid[row]) * (len(grid) - row)
	}

	fmt.Println(total)
}

func getClosesRock(grid [][]string, row int, col int) int {
	for row < len(grid) {
		if string(grid[row][col]) == "O" {
			return row
		}
		if string(grid[row][col]) == "#" {
			break
		}
		row++
	}
	return -1
}

func makeGrid(lines []string) [][]string {
	grid := [][]string{}
	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}
	return grid
}

func countRocks(s []string) int {
	count := 0
	for i := range s {
		if string(s[i]) == "O" {
			count++
		}
	}
	return count
}

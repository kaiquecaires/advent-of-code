package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	pwd, _ := os.Getwd()
	f, _ := os.ReadFile(pwd + "/go/2023/day_13/input.txt")
	grids := grids(string(f))
	total := 0

	for _, grid := range grids {
		total += findMirror(grid) * 100
		total += findMirror(transpose(grid))
	}

	fmt.Println(total)
}

func grids(f string) [][]string {
	lines := strings.Split(string(f), "\n\n")
	grids := make([][]string, len(lines))
	for i, line := range lines {
		grids[i] = strings.Split(line, "\n")
	}
	return grids
}

func findMirror(grid []string) int {
	for r := 1; r < len(grid); r++ {
		above := make([]string, len(grid[:r]))
		copy(above, grid[:r])

		below := grid[r:]

		slices.Reverse[[]string](above)

		if len(below) < len(above) {
			above = above[:len(below)]
		}

		if len(above) < len(below) {
			below = below[:len(above)]
		}

		total := 0

		for i := range above {
			total += sumDiff(above[i], below[i])
		}

		if total == 1 {
			return r
		}
	}

	return 0
}

func transpose(grid []string) []string {
	invs := make([]string, len(grid[0]))

	for r := range grid {
		for c := range grid[r] {
			invs[c] += string(grid[r][c])
		}
	}

	return invs
}

func sumDiff(s1 string, s2 string) int {
	sum := 0
	for i := range s1 {
		if s1[i] != s2[i] {
			sum++
		}
	}
	return sum
}

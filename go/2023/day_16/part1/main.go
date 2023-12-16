package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

var directions = map[string][2]int{
	"R": {0, 1},
	"D": {1, 0},
	"L": {0, -1},
	"U": {-1, 0},
}

var splitters = map[rune]map[string][]string{
	'-': {
		"R": {"R"},
		"D": {"R", "L"},
		"L": {"L"},
		"U": {"R", "L"},
	},
	'|': {
		"R": {"U", "D"},
		"D": {"D"},
		"L": {"U", "D"},
		"U": {"U"},
	},
}

var mirrors = map[rune]map[string]string{
	'/': {
		"R": "U",
		"D": "L",
		"L": "D",
		"U": "R",
	},
	'\\': {
		"R": "D",
		"D": "R",
		"L": "U",
		"U": "L",
	},
}

var visited = map[string][]string{
	"row_0_col_0": {"D"},
}

func Run(grid []string, pos [2]int, cdk string) {
	for {
		direction := directions[cdk]
		r, c := pos[0]+direction[0], pos[1]+direction[1]
		pos = [2]int{r, c}

		if r < 0 || r > len(grid)-1 || c < 0 || c > len(grid[0])-1 {
			break
		}

		if item, ok := visited[fmt.Sprintf("row_%d_col_%d", r, c)]; ok {
			if slices.Contains[[]string](item, cdk) {
				break
			}
		}

		visited[fmt.Sprintf("row_%d_col_%d", r, c)] = append(visited[fmt.Sprintf("row_%d_col_%d", r, c)], cdk)

		if _, ok := mirrors[rune(grid[r][c])]; ok {
			cdk = mirrors[rune(grid[r][c])][cdk]
		}

		if _, ok := splitters[rune(grid[r][c])]; ok {
			for _, dir := range splitters[rune(grid[r][c])][cdk] {
				Run(grid, [2]int{r, c}, dir)
			}
			break
		}
	}
}

func main() {
	pwd, _ := os.Getwd()
	f, _ := os.ReadFile(pwd + "/go/2023/day_16/input.txt")
	grid := strings.Split(string(f), "\n")
	Run(grid, [2]int{0, 0}, "D")
	fmt.Println(len(visited))
}

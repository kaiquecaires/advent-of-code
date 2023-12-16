package main

import (
	"fmt"
	"math"
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

func Run(grid []string, pos [2]int, cdk string, visited *map[string][]string) int {
	for {
		direction := directions[cdk]
		r, c := pos[0]+direction[0], pos[1]+direction[1]
		pos = [2]int{r, c}

		if r < 0 || r > len(grid)-1 || c < 0 || c > len(grid[0])-1 {
			break
		}

		if item, ok := (*visited)[fmt.Sprintf("row_%d_col_%d", r, c)]; ok {
			if slices.Contains[[]string](item, cdk) {
				break
			}
		}

		(*visited)[fmt.Sprintf("row_%d_col_%d", r, c)] = append((*visited)[fmt.Sprintf("row_%d_col_%d", r, c)], cdk)

		if _, ok := mirrors[rune(grid[r][c])]; ok {
			cdk = mirrors[rune(grid[r][c])][cdk]
		}

		if _, ok := splitters[rune(grid[r][c])]; ok {
			for _, dir := range splitters[rune(grid[r][c])][cdk] {
				Run(grid, [2]int{r, c}, dir, visited)
			}
			break
		}
	}

	return len(*visited)
}

func main() {
	pwd, _ := os.Getwd()
	f, _ := os.ReadFile(pwd + "/go/2023/day_16/input.txt")
	grid := strings.Split(string(f), "\n")
	max := 0

	for c := range grid[0] {
		for _, d := range []string{"D", "R", "L"} {
			total := Run(grid, [2]int{0, c}, d, &map[string][]string{
				fmt.Sprintf("row_%d_col_%d", 0, c): {d},
			})

			max = int(math.Max(float64(total), float64(max)))
		}
	}

	for c := range grid[len(grid)-1] {
		for _, d := range []string{"U", "R", "L"} {
			total := Run(grid, [2]int{0, c}, d, &map[string][]string{
				fmt.Sprintf("row_%d_col_%d", 0, c): {d},
			})

			max = int(math.Max(float64(total), float64(max)))
		}
	}

	for row := range grid {
		for _, d := range []string{"R", "D", "L", "U"} {
			total := Run(grid, [2]int{row, 0}, d, &map[string][]string{
				fmt.Sprintf("row_%d_col_%d", row, 0): {d},
			})

			max = int(math.Max(float64(total), float64(max)))

			total = Run(grid, [2]int{row, row - 1}, d, &map[string][]string{
				fmt.Sprintf("row_%d_col_%d", row, row-1): {d},
			})

			max = int(math.Max(float64(total), float64(max)))
		}
	}

	fmt.Println(max)
}

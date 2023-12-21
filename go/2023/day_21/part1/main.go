package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	pwd, _ := os.Getwd()
	f, _ := os.ReadFile(pwd + "/go/2023/day_21/input.txt")
	grid := strings.Split(string(f), "\n")
	startPos := [2]int{0, 0}

	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == 'S' {
				startPos = [2]int{row, col}
				break
			}
		}
	}

	queue := [][][2]int{{startPos}}
	steps := 64
	nGardenPlots := 0

	for steps >= 0 {
		positions := queue[0]
		queue = queue[1:]

		nGardenPlots = len(positions)
		newPositions := [][2]int{}

		visited := map[string]bool{}

		for _, p := range positions {
			r, c := p[0], p[1]

			for _, d := range [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
				nr, nc := r+d[0], c+d[1]

				if nr < 0 || nr >= len(grid) || nc < 0 || nc >= len(grid[0]) {
					continue
				}

				if (grid[nr][nc] == '.' || grid[nr][nc] == 'S') && !visited[fmt.Sprintf("%d_%d", nr, nc)] {
					visited[fmt.Sprintf("%d_%d", nr, nc)] = true
					newPositions = append(newPositions, [2]int{nr, nc})
				}
			}
		}

		queue = append(queue, newPositions)

		steps--
	}

	fmt.Println(nGardenPlots)
}

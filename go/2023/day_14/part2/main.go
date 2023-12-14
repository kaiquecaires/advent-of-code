package main

import (
	"fmt"
	"hash/fnv"
	"os"
	"reflect"
	"slices"
	"strings"
)

func main() {
	pwd, _ := os.Getwd()
	f, _ := os.ReadFile(pwd + "/go/2023/day_14/input.txt")
	grid := makeGrid(strings.Split(string(f), "\n"))
	seen := map[uint32]bool{}
	seen[hashGrid(grid)] = true

	count := 0
	grids := [][][]string{grid}

	for {
		count++
		grid = cycle(grid)

		if _, ok := seen[hashGrid(grid)]; ok {
			break
		}

		// Create a copy of the grid before appending it to grids
		gridCopy := make([][]string, len(grid))
		for i, row := range grid {
			gridCopy[i] = make([]string, len(row))
			copy(gridCopy[i], row)
		}

		grids = append(grids, gridCopy)
		seen[hashGrid(gridCopy)] = true
	}

	startCycle := 0

	for i, g := range grids {
		if reflect.DeepEqual(g, grid) {
			startCycle = i
			break
		}
	}

	grid = grids[(1000000000-startCycle)%(count-startCycle)+startCycle]
	total := 0

	for r, row := range grid {
		total += countRocks(row) * (len(grid) - r)
	}

	fmt.Println(total)
}

func cycle(grid [][]string) [][]string {
	// move to north
	grid = moveRocks(grid)

	// move to west
	grid = transpose(moveRocks(transpose(grid)))

	// move to south
	grid = transpose(flip(transpose(grid)))
	grid = moveRocks(grid)
	grid = transpose(flip(transpose(grid)))

	// move to east
	grid = flip(transpose(moveRocks(transpose(flip(grid)))))

	return grid
}

func moveRocks(grid [][]string) [][]string {
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
	return grid
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

func transpose(grid [][]string) [][]string {
	newGrid := make([][]string, len(grid[0]))

	for col := range grid[0] {
		for row := range grid {
			newGrid[col] = append(newGrid[col], grid[row][col])
		}
	}

	return newGrid
}

func flip(grid [][]string) [][]string {
	for row := range grid {
		slices.Reverse[[]string](grid[row])
	}
	return grid
}

func printGrid(grid [][]string) {
	fmt.Println("\n\nGrid:")
	for _, row := range grid {
		fmt.Println(strings.Join(row, ""))
	}
}

func hashGrid(grid [][]string) uint32 {
	s := ""

	for row := range grid {
		s += strings.Join(grid[row], "")
	}

	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

// code can be used for both, part1 and part2 solutions
func main() {
	pwd, _ := os.Getwd()
	f, _ := os.ReadFile(pwd + "/go/2023/day_11/input.txt")
	grid := strings.Split(string(f), "\n")
	emptyRows, emptyCols := getEmpty(&grid)
	points := getPoints(&grid)

	scale := 1000000
	total := 0

	for i, x := range points {
		for _, y := range points[:i] {
			r1, c1 := x[0], x[1]
			r2, c2 := y[0], y[1]

			sr, er := int(math.Min(float64(r1), float64(r2))), int(math.Max(float64(r1), float64(r2)))
			sc, ec := int(math.Min(float64(c1), float64(c2))), int(math.Max(float64(c1), float64(c2)))

			for r := sr; r < er; r++ {
				if slices.Contains[[]int, int](emptyRows, r) {
					total += scale
				} else {
					total += 1
				}
			}

			for c := sc; c < ec; c++ {
				if slices.Contains[[]int, int](emptyCols, c) {
					total += scale
				} else {
					total += 1
				}
			}
		}
	}

	fmt.Println(total)
}

func getEmpty(grid *[]string) ([]int, []int) {
	er := []int{}
	ec := []int{}

	for r := range *grid {
		if !strings.Contains((*grid)[r], "#") {
			er = append(er, r)
		}
	}

	for c := range (*grid)[0] {
		hasGalaxy := false
		for r := range *grid {
			if (*grid)[r][c] == '#' {
				hasGalaxy = true
			}
		}
		if !hasGalaxy {
			ec = append(ec, c)
		}
	}

	return er, ec
}

func getPoints(grid *[]string) [][2]int {
	points := [][2]int{}
	for r := range *grid {
		for c := range (*grid)[r] {
			if (*grid)[r][c] == '#' {
				points = append(points, [2]int{r, c})
			}
		}
	}
	return points
}

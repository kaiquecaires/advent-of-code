package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"regexp"
	"strings"
)

var (
	U = [2]int{-1, 0}
	D = [2]int{1, 0}
	L = [2]int{0, -1}
	R = [2]int{0, 1}
)

var dirs = map[string][2][2]int{
	"|": {U, D},
	"-": {L, R},
	"L": {U, R},
	"J": {U, L},
	"7": {D, L},
	"F": {D, R},
	".": {},
}

func main() {
	pwd, _ := os.Getwd()
	f, err := os.ReadFile(pwd + "/go/2023/day_10/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	grid := strings.Split(string(f), "\n")
	si, sj := findS(grid)
	grid[si] = strings.Replace(grid[si], "S", findSValue(si, sj, grid), -1)

	d := dirs[string(grid[si][sj])][0]
	dr, dc := d[0], d[1]

	pr, pc := si, sj
	nr, nc := pr+dr, pc+dc

	seen := map[string]bool{}

	seen[fmt.Sprintf("row_%d_column_%d", si, sj)] = true

	for !reflect.DeepEqual([2]int{nr, nc}, [2]int{si, sj}) {
		seen[fmt.Sprintf("row_%d_column_%d", nr, nc)] = true
		for _, d := range dirs[string(grid[nr][nc])] {
			if !reflect.DeepEqual([2]int{nr + d[0], nc + d[1]}, [2]int{pr, pc}) {
				pr = nr
				pc = nc
				nr = nr + d[0]
				nc = nc + d[1]
				break
			}
		}
	}

	for r := range grid {
		newRow := ""
		for c := range grid[r] {
			if _, ok := seen[fmt.Sprintf("row_%d_column_%d", r, c)]; !ok {
				newRow += "."
			} else {
				newRow += string(grid[r][c])
			}
		}
		grid[r] = newRow
	}

	pattern := regexp.MustCompile("L-*J|F-*7")
	replacer := func(match string) string {
		return strings.Repeat("", len(match))
	}

	inside := 0

	for _, row := range grid {
		row = pattern.ReplaceAllStringFunc(row, replacer)
		fmt.Println(row)
		within := false
		for _, ch := range row {
			if strings.Contains("|FL", string(ch)) {
				within = !within
			}
			if within && string(ch) == "." {
				inside++
			}
		}
	}

	fmt.Println(inside)
}

func findS(grid []string) (int, int) {
	for i, line := range grid {
		for j, c := range line {
			if c == 'S' {
				return i, j
			}
		}
	}
	return -1, -1
}

func findSValue(si int, sj int, grid []string) string {
	sValues := map[string]bool{}
	for k := range dirs {
		if k == "." {
			continue
		}
		sValues[k] = true
	}

	if si == 0 || !isIn(dirs[string(grid[si-1][sj])], D) {
		delete(sValues, "|")
		delete(sValues, "L")
		delete(sValues, "J")
	}

	if si == len(grid)-1 || !isIn(dirs[string(grid[si+1][sj])], U) {
		delete(sValues, "|")
		delete(sValues, "7")
		delete(sValues, "F")
	}

	if sj == 0 || !isIn(dirs[string(grid[si][sj-1])], R) {
		delete(sValues, "-")
		delete(sValues, "J")
		delete(sValues, "7")
	}

	if sj == len(grid[0])-1 || !isIn(dirs[string(grid[si][sj+1])], L) {
		delete(sValues, "-")
		delete(sValues, "L")
		delete(sValues, "F")
	}

	for k := range sValues {
		return k
	}

	return ""
}

func isIn(arrs [2][2]int, arr [2]int) bool {
	for _, cArr := range arrs {
		if reflect.DeepEqual(cArr, arr) {
			return true
		}
	}
	return false
}

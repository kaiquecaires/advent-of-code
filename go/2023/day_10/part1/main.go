package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	pwd, _ := os.Getwd()
	f, err := os.ReadFile(pwd + "/go/2023/day_10/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	rows := filter(strings.Split(string(f), "\n"), func(t string) bool {
		return t != ""
	})

	matrix := splitColumsFromRow(rows, func(row string) []string {
		return strings.Split(row, "")
	})

	i, j := findStartPoint(matrix)

	queue := [][3]int{{i, j, 0}}
	visited := map[string]bool{}
	maxDist := 0

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		if _, ok := visited[fmt.Sprintf("row_%d_column_%d", item[0], item[1])]; ok {
			continue
		}

		visited[fmt.Sprintf("row_%d_column_%d", item[0], item[1])] = true
		maxDist = int(math.Max(float64(maxDist), float64(item[2])))

		for _, nbr := range getNbrs(item[0], item[1], matrix) {
			if _, ok := visited[fmt.Sprintf("row_%d_column_%d", nbr[0], nbr[1])]; ok {
				continue
			}
			queue = append(queue, [3]int{nbr[0], nbr[1], item[2] + 1})
		}
	}

	fmt.Println(maxDist)
}

func filter[T any](t []T, f func(T) bool) []T {
	output := []T{}
	for _, i := range t {
		if f(i) {
			output = append(output, i)
		}
	}
	return output
}

func splitColumsFromRow(rows []string, f func(string) []string) [][]string {
	output := [][]string{}
	for _, row := range rows {
		output = append(output, f(row))
	}
	return output
}

func findStartPoint(m [][]string) (int, int) {
	for i := range m {
		for j := range m[i] {
			if m[i][j] == "S" {
				return i, j
			}
		}
	}
	return -1, -1
}

func getDnbrs(i int, j int, matrix [][]string) [][2]int {
	res := [][2]int{}

	if string(matrix[i][j]) == "S" {
		for _, values := range [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			ii, jj := values[0]+i, values[1]+j

			if !(ii >= 0 && ii < len(matrix) && jj >= 0 && jj < len(matrix[i])) {
				continue
			}

			if contains[[2]int](getNbrs(ii, jj, matrix), func(t [2]int) bool { return t[0] == i && t[1] == j }) {
				res = append(res, [2]int{values[0], values[1]})
			}
		}
	} else {
		res = map[string][][2]int{
			"|": {{1, 0}, {-1, 0}},
			"-": {{0, 1}, {0, -1}},
			"L": {{-1, 0}, {0, 1}},
			"J": {{-1, 0}, {0, -1}},
			"7": {{1, 0}, {0, -1}},
			"F": {{1, 0}, {0, 1}},
			".": {},
		}[string(matrix[i][j])]
	}

	return res
}

func getNbrs(i int, j int, matrix [][]string) [][2]int {
	res := [][2]int{}

	for _, values := range getDnbrs(i, j, matrix) {
		ii, jj := values[0]+i, values[1]+j

		if !(ii >= 0 && ii < len(matrix) && jj >= 0 && jj < len(matrix[i])) {
			continue
		}

		res = append(res, [2]int{ii, jj})
	}

	return res
}

func contains[T any](list []T, f func(T) bool) bool {
	for _, item := range list {
		if f(item) {
			return true
		}
	}

	return false
}

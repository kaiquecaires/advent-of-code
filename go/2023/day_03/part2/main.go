package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func findUpIndex(row int, column int, rows []string) (int, int) {
	if row-1 < 0 {
		return -1, -1
	}

	if unicode.IsDigit(rune(rows[row-1][column])) {
		return row - 1, column
	}

	return -1, -1
}

func findBottomIndex(row int, column int, rows []string) (int, int) {
	if row+1 > len(rows)-1 {
		return -1, -1
	}

	if unicode.IsDigit(rune(rows[row+1][column])) {
		return row + 1, column
	}

	return -1, -1
}

func findLeftIndex(row int, column int, rows []string) (int, int) {
	if column-1 < 0 {
		return -1, -1
	}

	if unicode.IsDigit(rune(rows[row][column-1])) {
		return row, column - 1
	}

	return -1, -1
}

func findRightIndex(row int, column int, rows []string) (int, int) {
	if column+1 > len(rows[row]) {
		return -1, -1
	}

	if unicode.IsDigit(rune(rows[row][column+1])) {
		return row, column + 1
	}

	return -1, -1
}

func findTopRightIndex(row int, column int, rows []string) (int, int) {
	if row-1 < 0 {
		return -1, -1
	}

	if column+1 > len(rows[row-1])-1 {
		return -1, -1
	}

	if unicode.IsDigit(rune(rows[row-1][column])) {
		return -1, -1
	}

	if unicode.IsDigit(rune(rows[row-1][column+1])) {
		return row - 1, column + 1
	}

	return -1, -1
}

func findTopLeftIndex(row int, column int, rows []string) (int, int) {
	if row-1 < 0 {
		return -1, -1
	}

	if column-1 < 0 {
		return -1, -1
	}

	if unicode.IsDigit(rune(rows[row-1][column])) {
		return -1, -1
	}

	if unicode.IsDigit(rune(rows[row-1][column-1])) {
		return row - 1, column - 1
	}

	return -1, -1
}

func findBottomLeftIndex(row int, column int, rows []string) (int, int) {
	if row+1 > len(rows)-1 {
		return -1, -1
	}

	if column-1 < 0 {
		return -1, -1
	}

	if unicode.IsDigit(rune(rows[row+1][column])) {
		return -1, -1
	}

	if unicode.IsDigit(rune(rows[row+1][column-1])) {
		return row + 1, column - 1
	}

	return -1, -1
}

func findBottomRightIndex(row int, column int, rows []string) (int, int) {
	if row+1 > len(rows)-1 {
		return -1, -1
	}

	if column+1 > len(rows[row+1])-1 {
		return -1, -1
	}

	if unicode.IsDigit(rune(rows[row+1][column])) {
		return -1, -1
	}

	if unicode.IsDigit(rune(rows[row+1][column+1])) {
		return row + 1, column + 1
	}

	return -1, -1
}

func getCompleteNumber(row int, column int, rows []string) int {
	numberBuilder := ""

	for i := column; i < len(rows[row]) && unicode.IsDigit(rune(rows[row][i])); i++ {
		numberBuilder += string(rows[row][i])
	}

	for i := column - 1; i >= 0 && unicode.IsDigit(rune(rows[row][i])); i-- {
		numberBuilder = string(rows[row][i]) + numberBuilder
	}

	number, _ := strconv.Atoi(numberBuilder)

	return number
}

func main() {
	pwd, _ := os.Getwd()
	f, _ := os.ReadFile(pwd + "/go/2023/day_03/input.txt")
	rows := strings.Split(string(f), "\n")
	total := 0

	for row, rowContent := range rows {
		for column, columnContent := range rowContent {
			if string(columnContent) == "*" {
				adjacentNumberIndexes := make([][]int, 0)

				if rowIndex, columnIndex := findUpIndex(row, column, rows); rowIndex != -1 {
					adjacentNumberIndexes = append(adjacentNumberIndexes, []int{rowIndex, columnIndex})
				}

				if rowIndex, columnIndex := findBottomIndex(row, column, rows); rowIndex != -1 {
					adjacentNumberIndexes = append(adjacentNumberIndexes, []int{rowIndex, columnIndex})
				}

				if rowIndex, columnIndex := findLeftIndex(row, column, rows); rowIndex != -1 {
					adjacentNumberIndexes = append(adjacentNumberIndexes, []int{rowIndex, columnIndex})
				}

				if rowIndex, columnIndex := findRightIndex(row, column, rows); rowIndex != -1 {
					adjacentNumberIndexes = append(adjacentNumberIndexes, []int{rowIndex, columnIndex})
				}

				if rowIndex, columnIndex := findBottomLeftIndex(row, column, rows); rowIndex != -1 {
					adjacentNumberIndexes = append(adjacentNumberIndexes, []int{rowIndex, columnIndex})
				}

				if rowIndex, columnIndex := findBottomRightIndex(row, column, rows); rowIndex != -1 {
					adjacentNumberIndexes = append(adjacentNumberIndexes, []int{rowIndex, columnIndex})
				}

				if rowIndex, columnIndex := findTopLeftIndex(row, column, rows); rowIndex != -1 {
					adjacentNumberIndexes = append(adjacentNumberIndexes, []int{rowIndex, columnIndex})
				}

				if rowIndex, columnIndex := findTopRightIndex(row, column, rows); rowIndex != -1 {
					adjacentNumberIndexes = append(adjacentNumberIndexes, []int{rowIndex, columnIndex})
				}

				if len(adjacentNumberIndexes) == 2 {
					completeNumbers := make([]int, 0)

					for _, indexes := range adjacentNumberIndexes {
						completeNumbers = append(completeNumbers, getCompleteNumber(indexes[0], indexes[1], rows))
					}

					total += completeNumbers[0] * completeNumbers[1]
				}
			}
		}
	}

	fmt.Print(total)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func isSymbol(char rune) bool {
	return string(char) != "." && !unicode.IsDigit(char)
}

func hasAdjacentSymbol(row int, column int, rows []string) bool {
	adjacentKeys := []map[string]int{
		{"row": row - 1, "column": column},
		{"row": row + 1, "column": column},
		{"row": row, "column": column - 1},
		{"row": row, "column": column + 1},
		{"row": row - 1, "column": column - 1},
		{"row": row - 1, "column": column + 1},
		{"row": row + 1, "column": column - 1},
		{"row": row + 1, "column": column + 1},
	}

	for _, adjacentKeys := range adjacentKeys {
		if adjacentKeys["row"] > len(rows)-1 || adjacentKeys["row"] < 0 {
			continue
		}

		if adjacentKeys["column"] > len(rows[adjacentKeys["row"]])-1 || adjacentKeys["column"] < 0 {
			continue
		}

		if isSymbol(rune(rows[adjacentKeys["row"]][adjacentKeys["column"]])) {
			return true
		}
	}

	return false
}

func main() {
	pwd, _ := os.Getwd()
	f, err := os.Open(pwd + "/go/2023/day_03/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	content := ""

	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	rows := strings.Split(content, "\n")
	total := 0

	for row, rowContent := range rows {
		numberBuilder := ""
		isAdjacent := false

		for column, char := range rowContent {
			if unicode.IsDigit(char) {
				numberBuilder += string(char)
				if !isAdjacent {
					isAdjacent = hasAdjacentSymbol(row, column, rows)
				}
				if column == len(rowContent)-1 {
					if numberBuilder != "" && isAdjacent {
						number, _ := strconv.Atoi(numberBuilder)
						total += number
						numberBuilder = ""
						isAdjacent = false
					}
				}
			} else {
				if numberBuilder != "" && isAdjacent {
					number, _ := strconv.Atoi(numberBuilder)
					total += number
				}
				numberBuilder = ""
				isAdjacent = false
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total => ", total)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

type Token struct {
	IsDigit bool
	Value   string
}

var numbers = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
}

func getNumberFromTokens(tokens []Token) int {
	var twoDigits [2]string

	for _, token := range tokens {
		if token.IsDigit {
			if twoDigits[0] == "" {
				twoDigits[0] = token.Value
			} else {
				twoDigits[1] = token.Value
			}
		}
	}

	if twoDigits[1] == "" {
		twoDigits[1] = twoDigits[0]
	}

	num, _ := strconv.Atoi(numbers[twoDigits[0]] + numbers[twoDigits[1]])

	return num
}

func processCharactersIntoTokens(characters string) []Token {
	tokens := make([]Token, 0)

	for _, char := range characters {
		t := Token{IsDigit: false, Value: string(char)}

		if unicode.IsDigit(char) {
			t.IsDigit = true
		} else {

			for index := range tokens {
				if tokens[index].IsDigit {
					continue
				}

				tokens[index].Value = tokens[index].Value + string(char)

				if _, ok := numbers[tokens[index].Value]; ok {
					tokens[index].IsDigit = true
				}
			}
		}

		tokens = append(tokens, t)
	}

	return tokens
}

func main() {
	pwd, _ := os.Getwd()
	f, err := os.Open(pwd + "/go/2023/day_01/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	total := 0

	for scanner.Scan() {
		characters := scanner.Text()
		tokens := processCharactersIntoTokens(characters)
		total += getNumberFromTokens(tokens)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(total)
}

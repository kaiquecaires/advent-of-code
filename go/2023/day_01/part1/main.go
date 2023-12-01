package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

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
		var twoChars [2]rune

		for _, char := range characters {
			if unicode.IsDigit(char) {
				if twoChars[0] == 0 {
					twoChars[0] = char
				} else {
					twoChars[1] = char
				}
			}
		}

		if twoChars[1] == 0 {
			twoChars[1] = twoChars[0]
		}

		num, _ := strconv.Atoi(string(twoChars[0]) + string(twoChars[1]))

		total += num
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(total)
}

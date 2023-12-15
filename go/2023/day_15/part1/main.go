package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	pwd, _ := os.Getwd()
	f, _ := os.ReadFile(pwd + "/go/2023/day_15/input.txt")
	initializationSequence := strings.Split(string(f), ",")

	total := 0

	for _, value := range initializationSequence {
		currentValue := 0
		for _, c := range value {
			currentValue += int(c)
			currentValue *= 17
			currentValue %= 256
		}
		total += currentValue
	}

	fmt.Println(total)
}

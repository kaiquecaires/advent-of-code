package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Len struct {
	label    string
	quantity int
}

func main() {
	pwd, _ := os.Getwd()
	f, _ := os.ReadFile(pwd + "/go/2023/day_15/input.txt")
	initializationSequence := strings.Split(string(f), ",")
	boxes := [256][]Len{}

	for _, len := range initializationSequence {
		if strings.Contains(len, "=") {
			values := strings.Split(len, "=")
			label := values[0]
			quantity, _ := strconv.Atoi(values[1])
			index := getBoxIndex(label)
			insert(&boxes, index, Len{label: label, quantity: quantity})
		} else {
			values := strings.Split(len, "-")
			label := values[0]
			index := getBoxIndex(label)
			remove(&boxes, index, label)
		}
	}

	total := 0

	for boxIdx, box := range boxes {
		for lenIdx, len := range box {
			total += (boxIdx + 1) * (lenIdx + 1) * len.quantity
		}
	}

	fmt.Println(total)
}

func getBoxIndex(chars string) int {
	currentValue := 0
	for _, c := range chars {
		currentValue += int(c)
		currentValue *= 17
		currentValue %= 256
	}
	return currentValue
}

func insert(boxes *[256][]Len, boxesIndex int, newLen Len) {
	isThereAnExistentLen := false
	for i, len := range boxes[boxesIndex] {
		if newLen.label == len.label {
			boxes[boxesIndex][i] = newLen
			isThereAnExistentLen = true
		}
	}
	if !isThereAnExistentLen {
		boxes[boxesIndex] = append(boxes[boxesIndex], newLen)
	}
}

func remove(boxes *[256][]Len, boxesIndex int, label string) {
	newBox := []Len{}

	for _, len := range boxes[boxesIndex] {
		if len.label == label {
			continue
		}
		newBox = append(newBox, len)
	}

	boxes[boxesIndex] = newBox
}

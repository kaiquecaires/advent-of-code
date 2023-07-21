package main

import (
	"bufio"
	"fmt"
	"os"
)

func getItemsPriority() map[string]int {
	itemsArray := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
		"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
		"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	}

	itemsPriority := make(map[string]int)
	itemPriority := 1

	for _, item := range itemsArray {
		itemsPriority[item] = itemPriority
		itemPriority++
	}

	return itemsPriority
}

func unpackCompartments(items string) (string, string) {
	size := len(items) / 2
	compartment1 := items[:size]
	compartment2 := items[size:]
	return compartment1, compartment2
}

func findCommonItem(c1 string, c2 string) string {
	itemsCoveredInC1 := make(map[string]string)
	itemsCoveredInC2 := make(map[string]string)

	for index, item := range c1 {
		c1String := string(item)
		c2String := string(c2[index])

		if c1String == c2String {
			return c1String
		}

		if _, ok := itemsCoveredInC2[c1String]; ok {
			return c1String
		}

		if _, ok := itemsCoveredInC1[c2String]; ok {
			return c2String
		}

		itemsCoveredInC1[c1String] = c1String
		itemsCoveredInC2[c2String] = c2String
	}

	return ""
}

func main() {
	file, err := os.Open("./2022/go/day-3/input.txt")

	if err != nil {
		fmt.Printf("Error opening file %v", err)
		panic("Error")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	itemsPriority := getItemsPriority()
	prioritiesSum := 0

	for scanner.Scan() {
		text := scanner.Text()
		c1, c2 := unpackCompartments(text)
		commonItem := findCommonItem(c1, c2)
		if commonItem != "" {
			fmt.Println(c1, c2, commonItem, itemsPriority[commonItem])
			prioritiesSum += itemsPriority[commonItem]
		}
	}

	fmt.Print(prioritiesSum)
}

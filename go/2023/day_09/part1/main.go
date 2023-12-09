package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	pwd, _ := os.Getwd()
	f, _ := os.ReadFile(pwd + "/go/2023/day_09/input.txt")
	rows := strings.Split(string(f), "\n")
	total := 0

	for _, row := range rows {
		if row == "" {
			continue
		}
		nums := convert[string, int](strings.Split(row, " "), func(x string) int {
			num, _ := strconv.Atoi(x)
			return num
		})

		parsedNums := parseNums(nums, [][]int{nums})
		total += predict(parsedNums)
	}

	fmt.Println(total)
}

func convert[T any, V any](slice []T, f func(T) V) []V {
	output := []V{}

	for _, s := range slice {
		output = append(output, f(s))
	}

	return output
}

func parseNums(nums []int, numsMatrix [][]int) [][]int {
	temp := []int{}

	for i := 0; i < len(nums)-1; i++ {
		temp = append(temp, nums[i+1]-nums[i])
	}

	numsMatrix = append(numsMatrix, temp)

	if count[int](temp, func(x int) bool { return x == 0 }) == len(temp) {
		return numsMatrix
	}

	return parseNums(temp, numsMatrix)
}

func count[T any](slice []T, f func(T) bool) int {
	count := 0
	for _, s := range slice {
		if f(s) {
			count++
		}
	}
	return count
}

func predict(nums [][]int) int {
	prev := 0

	for i := len(nums) - 2; i >= 0; i-- {
		newNum := nums[i][len(nums[i])-1] + prev
		nums[i] = append(nums[i], newNum)
		prev = newNum
	}

	return prev
}

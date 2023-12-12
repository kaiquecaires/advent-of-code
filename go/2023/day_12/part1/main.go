package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	pwd, _ := os.Getwd()
	f, _ := os.ReadFile(pwd + "/go/2023/day_12/input.txt")
	lines := strings.Split(string(f), "\n")

	total := 0
	for _, line := range lines {
		columns := strings.Split(line, " ")
		cfg := columns[0]
		nums := convertNums(strings.Split(columns[1], ","))
		total += count(cfg, nums)
	}

	fmt.Println(total)

}

func convertNums(numsStr []string) []int {
	nums := []int{}

	for _, num := range numsStr {
		newNum, _ := strconv.Atoi(num)
		nums = append(nums, newNum)
	}

	return nums
}

func count(cfg string, nums []int) int {
	if cfg == "" {
		if len(nums) == 0 {
			return 1
		} else {
			return 0
		}
	}

	if len(nums) == 0 {
		if strings.Contains(cfg, "#") {
			return 0
		} else {
			return 1
		}
	}

	result := 0

	if strings.Contains(".?", string(cfg[0])) {
		result += count(cfg[1:], nums)
	}

	if strings.Contains("#?", string(cfg[0])) {
		if nums[0] <= len(cfg) && !strings.Contains(cfg[:nums[0]], ".") && (nums[0] == len(cfg) || cfg[nums[0]] != '#') {
			if nums[0] == len(cfg) {
				result += count("", nums[1:])
			} else {
				result += count(cfg[nums[0]+1:], nums[1:])
			}
		}
	}

	return result
}

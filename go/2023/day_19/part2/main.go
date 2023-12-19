package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Workflow struct {
	rules    []Rule
	fallback string
}

type Rule struct {
	key    string
	cmp    string
	n      int
	target string
}

var workflows = map[string]Workflow{}

func main() {
	pwd, _ := os.Getwd()
	f, _ := os.ReadFile(pwd + "/go/2023/day_19/input.txt")
	input := strings.Split(string(f), "\n\n")

	for _, line := range strings.Split(input[0], "\n") {
		values := strings.Split(line[:len(line)-1], "{")
		name := values[0]
		rest := values[1]
		rules := strings.Split(rest, ",")

		fallback := rules[len(rules)-1]

		var workflowRules []Rule
		for _, rule := range rules[0 : len(rules)-1] {
			ruleValues := strings.Split(rule, ":")
			key := ruleValues[0][0]
			cmp := ruleValues[0][1]
			n, _ := strconv.Atoi(ruleValues[0][2:])
			workflowRules = append(workflowRules, Rule{key: string(key), cmp: string(cmp), n: n, target: ruleValues[1]})
		}

		workflows[name] = Workflow{
			fallback: fallback,
			rules:    workflowRules,
		}
	}

	ranges := map[string][2]int{}

	for _, k := range "xmas" {
		ranges[string(k)] = [2]int{1, 4000}
	}

	fmt.Println(count(ranges, "in"))
}

func count(ranges map[string][2]int, name string) int {
	if name == "R" {
		return 0
	}

	if name == "A" {
		product := 1
		for _, r := range ranges {
			product *= r[1] - r[0] + 1
		}
		return product
	}

	workflow := workflows[name]

	total := 0

	isBreaked := false

	for _, rule := range workflow.rules {
		low := ranges[rule.key][0]
		high := ranges[rule.key][1]

		var T, F [2]int

		if rule.cmp == "<" {
			T = [2]int{low, int(math.Min(float64(rule.n-1), float64(high)))}
			F = [2]int{int(math.Max(float64(rule.n), float64(low))), high}
		} else {
			T = [2]int{int(math.Max(float64(rule.n+1), float64(low))), high}
			F = [2]int{low, int(math.Min(float64(rule.n), float64(high)))}
		}

		if T[0] <= T[1] {
			mapCopy := deepCopy(ranges)
			mapCopy[rule.key] = T
			total += count(mapCopy, rule.target)
		}

		if F[0] <= F[1] {
			ranges[rule.key] = F
		} else {
			isBreaked = true
			break
		}
	}

	if !isBreaked {
		total += count(ranges, workflow.fallback)
	}

	return total
}

func deepCopy(original map[string][2]int) map[string][2]int {
	copy := make(map[string][2]int)
	for k, v := range original {
		copy[k] = v
	}
	return copy
}

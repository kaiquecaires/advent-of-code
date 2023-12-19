package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var workflows map[string]string

func main() {
	pwd, _ := os.Getwd()
	f, _ := os.ReadFile(pwd + "/go/2023/day_19/input.txt")
	input := strings.Split(string(f), "\n\n")
	rawWorkflows := strings.Split(input[0], "\n")
	workflows = make(map[string]string, len(rawWorkflows))

	for _, item := range rawWorkflows {
		values := strings.Split(item, "{")
		workflows[values[0]] = values[1][0 : len(values[1])-1]
	}

	partRatings := strings.Split(input[1], "\n")
	total := 0

	for _, rating := range partRatings {
		result := process(rating, "in")

		if result == "A" {
			for _, value := range ratingValuesMap(rating) {
				total += value
			}
		}
	}

	fmt.Println(total)
}

func process(rating string, workflowKey string) string {
	ratingValues := ratingValuesMap(rating)

	steps := strings.Split(workflows[workflowKey], ",")
	computations := []string{}

	for _, step := range steps {
		if !strings.Contains(step, "<") && !strings.Contains(step, ">") {
			computations[len(computations)-1] += "," + step
		} else {
			computations = append(computations, step)
		}
	}

	for _, computation := range computations {
		values := strings.Split(computation, ":")
		operation := values[0]

		key := string(operation[0])
		operator := string(operation[1])
		value, _ := strconv.Atoi(string(operation[2:]))

		consequences := strings.Split(values[1], ",")

		if operator == ">" {
			if ratingValues[key] > value {
				if strings.Contains("AR", string(consequences[0])) {
					return string(consequences[0])
				} else {
					return process(rating, string(consequences[0]))
				}
			} else {
				if len(consequences) == 2 {
					if strings.Contains("AR", string(consequences[1])) {
						return string(consequences[1])
					} else {
						return process(rating, string(consequences[1]))
					}
				}
			}
		} else if operator == "<" {
			if ratingValues[key] < value {
				if strings.Contains("AR", string(consequences[0])) {
					return string(consequences[0])
				} else {
					return process(rating, string(consequences[0]))
				}
			} else {
				if len(consequences) == 2 {
					if strings.Contains("AR", string(consequences[1])) {
						return string(consequences[1])
					} else {
						return process(rating, string(consequences[1]))
					}
				}
			}
		}
	}

	return ""
}

func ratingValuesMap(rating string) map[string]int {
	rawRatingValues := strings.Split(rating[1:len(rating)-1], ",")
	ratingValues := make(map[string]int, len(rawRatingValues))

	for _, currentRating := range rawRatingValues {
		values := strings.Split(currentRating, "=")
		ratingKey := values[0]
		ratingValue, _ := strconv.Atoi(values[1])
		ratingValues[ratingKey] = ratingValue
	}

	return ratingValues
}

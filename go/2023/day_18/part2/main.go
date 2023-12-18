package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	pwd, _ := os.Getwd()
	f, _ := os.ReadFile(pwd + "/go/2023/day_18/input.txt")
	lines := strings.Split(string(f), "\n")

	directions := map[string][2]int{
		"R": {0, 1},
		"D": {1, 0},
		"L": {0, -1},
		"U": {-1, 0},
	}

	points := [][2]int{{0, 0}}
	b := 0

	for _, line := range lines {

		moveTimes, dk := Parse(line)
		b += moveTimes
		point := points[len(points)-1]
		dir := directions[dk]
		r, c := point[0], point[1]
		points = append(points, [2]int{r + dir[0]*moveTimes, c + dir[1]*moveTimes})
	}

	area := 0

	for i := range points {
		x, y1, y2 := i, (i+1)%(len(points)-1), i-1

		if y2 == -1 {
			y2 = len(points) - 1
		}

		area = area + (points[x][0] * (points[y1][1] - points[y2][1]))
	}

	// Shoalesce formula
	area = int(math.Abs(float64(area))) / 2

	// Pick's theorem
	i := area - int(b/2) + 1

	// total area = inner area + boundary area
	fmt.Println(i + b)
}

func Parse(line string) (int, string) {
	dirs := map[int]string{
		0: "R",
		1: "D",
		2: "L",
		3: "U",
	}

	values := strings.Split(line, " ")
	dirValue, _ := strconv.ParseInt(values[2][2:len(values[2])-2], 16, 64)
	v, _ := strconv.Atoi(string(values[2][len(values[2])-2]))

	return int(dirValue), dirs[v]
}

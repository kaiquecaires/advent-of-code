package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

type State struct {
	hl, r, c, dr, dc, n int
}

type PriorityQueue []State

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].hl < pq[j].hl }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(State)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func main() {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/go/2023/day_17/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid [][]int

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for i, char := range line {
			row[i], _ = strconv.Atoi(string(char))
		}
		grid = append(grid, row)
	}

	seen := make(map[string]bool)
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	pq.Push(State{0, 0, 0, 0, 0, 0})

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(State)

		if current.r == len(grid)-1 && current.c == len(grid[0])-1 {
			fmt.Println(current.hl)
			break
		}

		if seen[fmt.Sprintf("%d_%d_%d_%d_%d", current.r, current.c, current.dr, current.dc, current.n)] {
			continue
		}

		seen[fmt.Sprintf("%d_%d_%d_%d_%d", current.r, current.c, current.dr, current.dc, current.n)] = true

		if current.n < 3 && (current.dr != 0 || current.dc != 0) {
			nr := current.r + current.dr
			nc := current.c + current.dc
			if 0 <= nr && nr < len(grid) && 0 <= nc && nc < len(grid[0]) {
				heap.Push(&pq, State{current.hl + grid[nr][nc], nr, nc, current.dr, current.dc, current.n + 1})
			}
		}

		for _, move := range [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			ndr, ndc := move[0], move[1]
			if (ndr != current.dr || ndc != current.dc) && (ndr != -current.dr || ndc != -current.dc) {
				nr := current.r + ndr
				nc := current.c + ndc
				if 0 <= nr && nr < len(grid) && 0 <= nc && nc < len(grid[0]) {
					heap.Push(&pq, State{current.hl + grid[nr][nc], nr, nc, ndr, ndc, 1})
				}
			}
		}
	}
}

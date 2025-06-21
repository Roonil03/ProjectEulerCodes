package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cell struct {
	i, j  int
	dist  int
	index int
}

type MinHeap []*Cell

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(a, b int) bool { return h[a].dist < h[b].dist }
func (h MinHeap) Swap(a, b int) {
	h[a], h[b] = h[b], h[a]
	h[a].index, h[b].index = a, b
}
func (h *MinHeap) Push(x any) {
	n := len(*h)
	cell := x.(*Cell)
	cell.index = n
	*h = append(*h, cell)
}
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	cell := old[n-1]
	old[n-1] = nil
	*h = old[0 : n-1]
	return cell
}

func main() {
	filepath := "Problem0081. PathSum_TwoWays/input.txt"
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var mat [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ",")
		row := make([]int, len(parts))
		for k, s := range parts {
			row[k], _ = strconv.Atoi(s)
		}
		mat = append(mat, row)
	}
	n := len(mat)
	const INF = 1<<31 - 1
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = INF
		}
	}
	dist[0][0] = mat[0][0]
	pq := &MinHeap{}
	heap.Init(pq)
	heap.Push(pq, &Cell{0, 0, dist[0][0], 0})
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for pq.Len() > 0 {
		u := heap.Pop(pq).(*Cell)
		if u.i == n-1 && u.j == n-1 {
			fmt.Println("The answer is:", u.dist)
			return
		}
		if u.dist > dist[u.i][u.j] {
			continue
		}
		for _, d := range dirs {
			ni, nj := u.i+d[0], u.j+d[1]
			if ni >= 0 && ni < n && nj >= 0 && nj < n {
				nd := u.dist + mat[ni][nj]
				if nd < dist[ni][nj] {
					dist[ni][nj] = nd
					heap.Push(pq, &Cell{ni, nj, nd, 0})
				}
			}
		}
	}
}

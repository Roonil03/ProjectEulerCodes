package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Item struct {
	i, j int
	dist int
}
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(a, b int) bool { return pq[a].dist < pq[b].dist }
func (pq PriorityQueue) Swap(a, b int)      { pq[a], pq[b] = pq[b], pq[a] }
func (pq *PriorityQueue) Push(x any)        { *pq = append(*pq, x.(*Item)) }
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func main() {
	fielpath := "Problem0082. PathSum_ThreeWays/input.txt"
	file, _ := os.Open(fielpath)
	defer file.Close()
	sc := bufio.NewScanner(file)
	var mat [][]int
	for sc.Scan() {
		parts := strings.Split(sc.Text(), ",")
		row := make([]int, len(parts))
		for k, s := range parts {
			row[k], _ = strconv.Atoi(s)
		}
		mat = append(mat, row)
	}
	n := len(mat)
	const INF = 1 << 60
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = INF
		}
	}
	pq := &PriorityQueue{}
	heap.Init(pq)
	for i := 0; i < n; i++ {
		dist[i][0] = mat[i][0]
		heap.Push(pq, &Item{i, 0, dist[i][0]})
	}
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, 1}}
	for pq.Len() > 0 {
		u := heap.Pop(pq).(*Item)
		if u.j == n-1 {
			fmt.Println("The answer is: ", u.dist)
			return
		}
		if u.dist > dist[u.i][u.j] {
			continue
		}
		for _, d := range dirs {
			ni, nj := u.i+d[0], u.j+d[1]
			if ni >= 0 && ni < n && nj < n {
				nd := u.dist + mat[ni][nj]
				if nd < dist[ni][nj] {
					dist[ni][nj] = nd
					heap.Push(pq, &Item{ni, nj, nd})
				}
			}
		}
	}
}

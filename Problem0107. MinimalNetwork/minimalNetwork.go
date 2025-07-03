package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const N = 40

type e struct{ u, v, w int }

func main() {
	filepath := "Problem0107. MinimalNetwork/inputs.txt"
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	var es []e
	sum := 0
	i := 0

	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if line == "" {
			continue // Skip empty lines
		}

		ss := strings.Split(line, ",")
		if len(ss) != 40 {
			fmt.Printf("Warning: Row %d has %d entries, expected 40\n", i, len(ss))
			continue
		}

		for j, x := range ss {
			if x != "-" && j > i {
				w, err := strconv.Atoi(x)
				if err != nil {
					fmt.Printf("Error parsing weight '%s' at (%d,%d): %v\n", x, i, j, err)
					continue
				}
				es = append(es, e{i, j, w})
				sum += w
			}
		}
		i++
	}

	if err := s.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	fmt.Printf("Read %d edges with total weight %d\n", len(es), sum)

	if len(es) == 0 {
		fmt.Println("No edges found! Check file format and path.")
		return
	}

	sort.Slice(es, func(a, b int) bool { return es[a].w < es[b].w })
	p := make([]int, N)
	for i := 0; i < N; i++ {
		p[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if p[x] != x {
			p[x] = find(p[x])
		}
		return p[x]
	}
	mst := 0
	for _, ed := range es {
		ru, rv := find(ed.u), find(ed.v)
		if ru != rv {
			p[ru] = rv
			mst += ed.w
		}
	}
	fmt.Println("The answer is:", sum-mst)
}

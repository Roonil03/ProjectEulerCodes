package main

import "fmt"

const Limit = 1000000

func main() {
	sd := make([]int, Limit+1)
	for i := 1; i <= Limit/2; i++ {
		for j := i * 2; j <= Limit; j += i {
			sd[j] += i
		}
	}
	pos := make([]int, Limit+1)
	seen := make([]bool, Limit+1)
	bestLen, res := 0, 0
	for i := 1; i <= Limit; i++ {
		if seen[i] {
			continue
		}
		var path []int
		c := i
		for {
			if c > Limit {
				break
			}
			if pos[c] > 0 {
				if c == i && len(path) > bestLen {
					bestLen = len(path)
					res = i
					for _, n := range path {
						seen[n] = true
					}
				}
				break
			}
			pos[c] = len(path) + 1
			path = append(path, c)
			c = sd[c]
		}
		for _, n := range path {
			pos[n] = 0
		}
	}
	fmt.Println("The answer is:", res)
}

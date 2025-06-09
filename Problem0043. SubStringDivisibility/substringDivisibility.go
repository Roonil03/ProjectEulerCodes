package main

import "fmt"

var primes = []int{2, 3, 5, 7, 11, 13, 17}

func main() {
	used := [10]bool{}
	var digs [10]int
	sum := 0
	var dfs func(pos int)
	dfs = func(pos int) {
		if pos == 10 {
			n := 0
			for i := 0; i < 10; i++ {
				n = n*10 + digs[i]
			}
			sum += n
			return
		}
		for d := 0; d <= 9; d++ {
			if used[d] {
				continue
			}
			digs[pos] = d
			if pos >= 3 {
				v := digs[pos-2]*100 + digs[pos-1]*10 + digs[pos]
				if v%primes[pos-3] != 0 {
					continue
				}
			}
			used[d] = true
			dfs(pos + 1)
			used[d] = false
		}
	}
	dfs(0)
	fmt.Println("The answer is: ", sum)
}

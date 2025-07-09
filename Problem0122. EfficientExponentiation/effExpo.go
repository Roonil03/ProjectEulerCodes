package main

import "fmt"

// var best = make([]int, tgt+1)

// func binaryMethodBound(n int) int {
// 	if n == 1 {
// 		return 0
// 	}
// 	bits := 0
// 	temp := n
// 	for temp > 0 {
// 		if temp&1 == 1 {
// 			bits++
// 		}
// 		temp >>= 1
// 	}
// 	return bits - 1 + popcount(n) - 1
// }

// func popcount(n int) int {
// 	count := 0
// 	for n > 0 {
// 		if n&1 == 1 {
// 			count++
// 		}
// 		n >>= 1
// 	}
// 	return count
// }

// func dfs(c []int, target, maxDepth int) bool {
// 	if len(c) > maxDepth {
// 		return false
// 	}
// 	last := c[len(c)-1]
// 	if last == target {
// 		return true
// 	}
// 	if last > target {
// 		return false
// 	}
// 	for i := len(c) - 1; i >= 0; i-- {
// 		sum := c[i] + last
// 		if sum > target {
// 			continue
// 		}
// 		exists := false
// 		for _, val := range c {
// 			if val == sum {
// 				exists = true
// 				break
// 			}
// 		}
// 		if exists {
// 			continue
// 		}
// 		newChain := append(c, sum)
// 		if dfs(newChain, target, maxDepth) {
// 			return true
// 		}
// 	}
// 	return false
// }

// func m(k int) int {
// 	if k == 1 {
// 		return 0
// 	}
// 	if best[k] != 0 {
// 		return best[k]
// 	}
// 	upperBound := binaryMethodBound(k)
// 	for depth := 1; depth <= upperBound; depth++ {
// 		if dfs([]int{1}, k, depth) {
// 			best[k] = depth
// 			return depth
// 		}
// 	}
// 	best[k] = upperBound
// 	return upperBound
// }

const lim = 200

var memo = [lim + 1]int{0}

func dfs(chain []int, left, target int) bool {
	if left == 0 {
		return chain[len(chain)-1] == target
	}
	max := chain[len(chain)-1]
	for i := len(chain) - 1; i >= 0; i-- {
		next := max + chain[i]
		if next > target || next <= max {
			continue
		}
		chain = append(chain, next)
		if dfs(chain, left-1, target) {
			return true
		}
		chain = chain[:len(chain)-1]
	}
	return false
}

func m(k int) int {
	if memo[k] != 0 || k == 1 {
		return memo[k]
	}
	for d := 1; ; d++ {
		if dfs([]int{1}, d, k) {
			memo[k] = d
			return d
		}
	}
}

func main() {
	sum := 0
	for k := 1; k <= lim; k++ {
		sum += m(k)
	}
	fmt.Println("The answer is:", sum)
}

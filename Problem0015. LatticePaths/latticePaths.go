package main

import (
	"fmt"
)

func main() {
	n := 20
	grid := make([][]int, n+1)
	for i := range grid {
		grid[i] = make([]int, n+1)
	}
	for i := 0; i <= n; i++ {
		grid[i][n] = 1
		grid[n][i] = 1
	}
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			grid[i][j] = grid[i+1][j] + grid[i][j+1]
		}
	}
	fmt.Println("The result is: ", grid[0][0])
	// res := fact(int64(n*2)) / int64(math.Pow(float64(fact(int64(n))), float64(2)))
	// fmt.Println("The result is: ", res)
	// fmt.Println(fact(40))
}

// func fact(a int64) int64 {
// 	if a == 2 {
// 		return int64(2)
// 	}
// 	return a * fact(a-1)
// }

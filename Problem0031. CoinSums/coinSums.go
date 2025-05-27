package main

import "fmt"

func main() {
	target := 200
	coins := []int{1, 2, 5, 10, 20, 50, 100, 200}
	dp := make([]int, target+1)
	dp[0] = 1
	for _, c := range coins {
		for j := c; j <= target; j++ {
			dp[j] += dp[j-c]
		}
	}
	fmt.Println("Number of ways is: ", dp[target])
}

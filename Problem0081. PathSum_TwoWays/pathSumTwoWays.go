package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filepath := "Problem0081. PathSum_TwoWays/input.txt"
	file, _ := os.Open(filepath)
	defer file.Close()
	var mat [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ",")
		row := make([]int, len(parts))
		for i, s := range parts {
			row[i], _ = strconv.Atoi(s)
		}
		mat = append(mat, row)
	}
	n := len(mat)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = mat[0][0]
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + mat[0][j]
	}
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + mat[i][0]
	}
	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			if dp[i-1][j] < dp[i][j-1] {
				dp[i][j] = dp[i-1][j] + mat[i][j]
			} else {
				dp[i][j] = dp[i][j-1] + mat[i][j]
			}
		}
	}
	fmt.Println("The answer is: ", dp[n-1][n-1])
}

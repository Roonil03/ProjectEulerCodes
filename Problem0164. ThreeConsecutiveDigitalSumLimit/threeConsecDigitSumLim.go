package main

import "fmt"

func main() {
	var dp [10][10]int64
	for a := 1; a <= 9; a++ {
		for b := 0; b <= 9; b++ {
			if a+b <= 9 {
				dp[a][b] = 1
			}
		}
	}
	for l := 3; l <= 20; l++ {
		var ndp [10][10]int64
		for b := 0; b <= 9; b++ {
			for c := 0; c <= 9; c++ {
				if b+c <= 9 {
					var sum int64
					for a := 0; a <= 9-b-c; a++ {
						sum += dp[a][b]
					}
					ndp[b][c] = sum
				}
			}
		}
		dp = ndp
	}
	var res int64
	for b := 0; b <= 9; b++ {
		for c := 0; c <= 9; c++ {
			res += dp[b][c]
		}
	}
	fmt.Println(res)
}

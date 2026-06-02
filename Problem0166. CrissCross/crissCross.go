package main

import "fmt"

func main() {
	res := 0
	for a := 0; a <= 9; a++ {
		for b := 0; b <= 9; b++ {
			for c := 0; c <= 9; c++ {
				for d := 0; d <= 9; d++ {
					sum := a + b + c + d
					for e := 0; e <= 9; e++ {
						for f := 0; f <= 9; f++ {
							for g := 0; g <= 9; g++ {
								h := sum - e - f - g
								if h < 0 || h > 9 {
									continue
								}
								for m := 0; m <= 9; m++ {
									j := sum - d - g - m
									if j < 0 || j > 9 {
										continue
									}
									i := sum - a - e - m
									if i < 0 || i > 9 {
										continue
									}
									n := sum - b - f - j
									if n < 0 || n > 9 {
										continue
									}
									val := m + n + sum - c - g - a - f
									if val&1 != 0 {
										continue
									}
									k := val / 2
									if k < 0 || k > 9 {
										continue
									}
									l := sum - i - j - k
									if l < 0 || l > 9 {
										continue
									}
									o := sum - c - g - k
									if o < 0 || o > 9 {
										continue
									}
									p := sum - m - n - o
									if p < 0 || p > 9 {
										continue
									}
									if d+h+l+p == sum {
										res++
									}
								}
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(res)
}

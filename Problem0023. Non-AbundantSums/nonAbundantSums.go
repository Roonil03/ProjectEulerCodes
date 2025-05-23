package main

import "fmt"

func main() {
	const d = 28123
	temp := make([]int, d+1)
	for i := 1; i <= d; i++ {
		for j := i * 2; j <= d; j += i {
			temp[j] += i
		}
	}
	var num []int
	for i := 12; i < d; i++ {
		if temp[i] > i {
			num = append(num, i)
		}
	}
	vis := make([]bool, d+1)
	for i := 0; i < len(num); i++ {
		for j := i; j < len(num); j++ {
			sum := num[i] + num[j]
			if sum <= d {
				vis[sum] = true
			} else {
				break
			}
		}
	}
	res := 0
	for i := 1; i <= d; i++ {
		if !vis[i] {
			res += i
		}
	}
	fmt.Println("The answer is:", res)
}

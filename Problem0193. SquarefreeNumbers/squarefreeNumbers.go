package main

import "fmt"

func main() {
	const lim = (1 << 25) - 1
	mu := make([]int8, lim+1)
	pr := make([]int, 0, lim/10)
	ip := make([]bool, lim+1)
	for i := 2; i <= lim; i++ {
		ip[i] = true
	}
	mu[1] = 1
	for i := 2; i <= lim; i++ {
		if ip[i] {
			pr = append(pr, i)
			mu[i] = -1
		}
		for _, p := range pr {
			if i*p > lim {
				break
			}
			ip[i*p] = false
			if i%p == 0 {
				mu[i*p] = 0
				break
			}
			mu[i*p] = -mu[i]
		}
	}
	var res int64
	x := (int64(1) << 50) - 1
	for i := int64(1); i <= lim; i++ {
		if mu[i] != 0 {
			res += int64(mu[i]) * (x / (i * i))
		}
	}
	fmt.Println(res)
}

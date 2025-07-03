package main

import (
	"fmt"
)

func comb(n, k int64) int64 {
	if k < 0 || k > n {
		return 0
	}
	if k > n-k {
		k = n - k
	}
	var res int64 = 1
	for i := int64(1); i <= k; i++ {
		res = res * (n - k + i) / i
	}
	return res
}

func u(n int64) int64 {
	var sum int64 = 0
	var power int64 = 1
	for m := int64(0); m <= 10; m++ {
		if m > 0 {
			power *= n
		}
		if m%2 == 0 {
			sum += power
		} else {
			sum -= power
		}
	}
	return sum
}

func main() {
	var res int64 = 0
	for k := int64(1); k <= 10; k++ {
		var fit int64 = 0
		for i := int64(1); i <= k; i++ {
			yi := u(i)
			sign := int64(1)
			if (k-i)%2 != 0 {
				sign = -1
			}
			numerator := yi * i * comb(k, i) * sign
			denominator := k + 1 - i
			fit += numerator / denominator
		}
		res += fit
	}
	fmt.Println("The answer is: ", res)
}

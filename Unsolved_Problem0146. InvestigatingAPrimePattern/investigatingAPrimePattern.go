package main

import (
	"fmt"
)

func pow(v, c, p int64) int64 {
	r := int64(1)
	v = v % p
	for c > 0 {
		if c&1 == 1 {
			r = (r * v) % p
		}
		c = c >> 1
		v = (v * v) % p
	}
	return r
}

func prime(n int64, k int) bool {
	if n <= 1 || n == 4 {
		return false
	}
	if n <= 3 {
		return true
	}
	d := n - 1
	for d%2 == 0 {
		d /= 2
	}
	for i := 0; i < k; i++ {
		if !test(d, n) {
			return false
		}
	}
	return true
}

func main() {
	s := int64(0)
	for j := int64(10); j < 150000000; j += 10 {
		if j%3 == 0 || j%7 == 0 || j%13 == 0 {
			continue
		}
		t := j * j
		if prime(t+1, 4) && prime(t+3, 4) && prime(t+7, 4) && prime(t+9, 4) && !prime(t+11, 4) && prime(t+13, 4) && !prime(t+17, 4) && !prime(t+19, 4) && !prime(t+21, 4) && !prime(t+23, 4) && prime(t+27, 4) {
			s += j
		}
	}
	fmt.Println(s)
}

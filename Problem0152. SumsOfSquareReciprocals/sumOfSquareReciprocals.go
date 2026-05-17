package main

import (
	"fmt"
	"sort"
)

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int64) int64 {
	return a / gcd(a, b) * b
}

func main() {
	var n []int64
	var m int64 = 1
	for i := int64(2); i <= 80; i++ {
		x := i
		for _, p := range []int64{2, 3, 5, 7, 13} {
			for x%p == 0 {
				x /= p
			}
		}
		if x == 1 {
			n = append(n, i)
			m = lcm(m, i)
		}
	}
	m2 := m * m
	var v []int64
	for _, x := range n {
		v = append(v, m2/(x*x))
	}
	sort.Slice(v, func(i, j int) bool {
		return v[i] > v[j]
	})
	s := make([]int64, len(v)+1)
	for i := len(v) - 1; i >= 0; i-- {
		s[i] = s[i+1] + v[i]
	}
	t := m2 / 2
	var h1 func(i int, c int64) int
	h1 = func(i int, c int64) int {
		if c == t {
			return 1
		}
		if i == len(v) || c > t || c+s[i] < t {
			return 0
		}
		return h1(i+1, c+v[i]) + h1(i+1, c)
	}
	fmt.Println(h1(0, 0))
}

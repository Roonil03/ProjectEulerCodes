package main

import "fmt"

func main() {
	const n = 15
	a := fib(30)
	b := fib(31)
	fmt.Println("The answer is:", a*b)
}

func fib(n int) int64 {
	if n <= 2 {
		return 1
	}
	m := [2][2]int64{{1, 1}, {1, 0}}
	r := pow(m, n-1)
	return r[0][0]
}

func pow(m [2][2]int64, p int) [2][2]int64 {
	r := [2][2]int64{{1, 0}, {0, 1}}
	for p > 0 {
		if p&1 == 1 {
			r = mul(r, m)
		}
		m = mul(m, m)
		p >>= 1
	}
	return r
}

func mul(a, b [2][2]int64) [2][2]int64 {
	return [2][2]int64{
		{a[0][0]*b[0][0] + a[0][1]*b[1][0], a[0][0]*b[0][1] + a[0][1]*b[1][1]},
		{a[1][0]*b[0][0] + a[1][1]*b[1][0], a[1][0]*b[0][1] + a[1][1]*b[1][1]},
	}
}

package main

import "fmt"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	const L = 100000000
	cnt := 0
	for m := 2; m*m < L; m++ {
		for n := 1; n < m; n++ {
			if (m-n)%2 == 1 && gcd(m, n) == 1 {
				a := m*m - n*n
				b := 2 * m * n
				c := m*m + n*n
				if a > b {
					a, b = b, a
				}
				if c%(b-a) == 0 {
					p := a + b + c
					if p < L {
						cnt += L / p
					}
				}
			}
		}
	}
	fmt.Println("The answer is:", cnt)
}

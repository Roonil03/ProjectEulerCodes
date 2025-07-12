package main

import "fmt"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func powmod(a, e, m int) int {
	r := 1 % m
	for e > 0 {
		if e&1 == 1 {
			r = (r * a) % m
		}
		a = (a * a) % m
		e >>= 1
	}
	return r
}

func ord(a, n int) int {
	if gcd(a, n) != 1 {
		return 0
	}
	k := 1
	x := a % n
	for x != 1 {
		x = (x * a) % n
		k++
	}
	return k
}

func hasOnlyFactors2And5(n int) bool {
	for n%2 == 0 {
		n /= 2
	}
	for n%5 == 0 {
		n /= 5
	}
	return n == 1
}

func main() {
	const L = 100000
	bs := make([]bool, L+1)
	ps := []int{}
	for i := 2; i <= L; i++ {
		if !bs[i] {
			ps = append(ps, i)
			for j := i * i; j <= L; j += i {
				bs[j] = true
			}
		}
	}
	s := 0
	for _, p := range ps {
		if p == 2 || p == 5 {
			s += p
			continue
		}
		o := ord(10, 9*p)
		if o > 0 && !hasOnlyFactors2And5(o) {
			s += p
		}
	}
	fmt.Println("The answer is:", s)
}

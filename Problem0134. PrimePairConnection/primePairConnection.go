package main

import "fmt"

func gcd(a, b int64) (int64, int64, int64) {
	if b == 0 {
		return a, 1, 0
	}
	g, x1, y1 := gcd(b, a%b)
	return g, y1, x1 - (a/b)*y1
}

func tens(x int64) int64 {
	t := int64(1)
	for t <= x {
		t *= 10
	}
	return t
}

func solve(p1, p2 int64) int64 {
	m1 := p2
	m2 := tens(p1)
	_, x, _ := gcd(m1, m2)
	res := p1 * x * m1
	prod := m1 * m2
	res %= prod
	if res < 0 {
		res += prod
	}
	return res
}

func main() {
	const L = 1000000
	bs := make([]bool, L+100)
	ps := []int64{}
	for i := int64(2); i <= L+50; i++ {
		if !bs[i] {
			ps = append(ps, i)
			for j := i * i; j <= L+50; j += i {
				bs[j] = true
			}
		}
	}
	s := int64(0)
	for i := 0; i < len(ps)-1; i++ {
		p1, p2 := ps[i], ps[i+1]
		if p1 < 5 || p1 > L {
			continue
		}
		s += solve(p1, p2)
	}
	fmt.Println("The answer is:", s)
}

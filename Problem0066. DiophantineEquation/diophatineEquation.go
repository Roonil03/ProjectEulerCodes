package main

import (
	"fmt"
	"math/big"
)

func isqrt(n int) int {
	if n == 0 {
		return 0
	}
	x := n
	y := (x + 1) / 2
	for y < x {
		x = y
		y = (x + n/x) / 2
	}
	return x
}

func isPerfectSquare(n int) bool {
	root := isqrt(n)
	return root*root == n
}

func continuedFractionPeriod(d int) (int, []int) {
	if isPerfectSquare(d) {
		return 0, []int{}
	}
	a0 := isqrt(d)
	m, n, a := 0, 1, a0
	period := []int{}
	for {
		m = n*a - m
		n = (d - m*m) / n
		a = (a0 + m) / n
		period = append(period, a)
		if a == 2*a0 {
			break
		}
	}
	return a0, period
}

func solvePellEquation(d int) (*big.Int, *big.Int) {
	if isPerfectSquare(d) {
		return nil, nil
	}
	a0, period := continuedFractionPeriod(d)
	if len(period) == 0 {
		return nil, nil
	}
	p_prev := big.NewInt(1)
	p_curr := big.NewInt(int64(a0))
	q_prev := big.NewInt(0)
	q_curr := big.NewInt(1)
	for i := 0; ; i++ {
		a_i := period[i%len(period)]
		p_next := new(big.Int).Mul(big.NewInt(int64(a_i)), p_curr)
		p_next.Add(p_next, p_prev)
		q_next := new(big.Int).Mul(big.NewInt(int64(a_i)), q_curr)
		q_next.Add(q_next, q_prev)
		left := new(big.Int).Mul(p_next, p_next)
		right := new(big.Int).Mul(big.NewInt(int64(d)), q_next)
		right.Mul(right, q_next)
		left.Sub(left, right)
		if left.Cmp(big.NewInt(1)) == 0 {
			return p_next, q_next
		}
		p_prev.Set(p_curr)
		p_curr.Set(p_next)
		q_prev.Set(q_curr)
		q_curr.Set(q_next)
	}
}

func main() {
	maxX := big.NewInt(0)
	res := 0
	for d := 2; d <= 1000; d++ {
		if isPerfectSquare(d) {
			continue
		}
		x, _ := solvePellEquation(d)
		if x != nil && x.Cmp(maxX) > 0 {
			maxX.Set(x)
			res = d
		}
	}
	fmt.Printf("The answer is:  %d\n", res)
}

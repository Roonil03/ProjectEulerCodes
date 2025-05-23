package main

import (
	"fmt"
	"slices"
)

var mem = make(map[int]int)

func fact(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	if val, ok := mem[n]; ok {
		return val
	}
	res := n * fact(n-1)
	mem[n] = res
	return res
}

func fike(dig []int, num int) []int {
	n := len(dig)
	res := make([]int, 0, n)
	a := make([]int, n)
	copy(a, dig)
	num--
	for i := n; i > 0; i-- {
		f := fact(i - 1)
		id := num / f
		res = append(res, a[id])
		a = slices.Delete(a, id, id+1)
		num = num % f
	}
	return res
}

func main() {
	dig := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	pos := 1000000
	res := fike(dig, pos)
	fmt.Print(pos, ". The answer is: ")
	for _, q := range res {
		fmt.Print(q)
	}
	fmt.Println()
}

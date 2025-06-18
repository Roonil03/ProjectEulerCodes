package main

import "fmt"

const lim = 1000000

func fact(n int) int {
	if n == 1 {
		return 1
	}
	return fact(n-1) * n
}

func main() {
	sum := 6 * fact(9)
	f := [...]int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}
	mem := make([]int, sum+1)
	for _, v := range []int{169, 363601, 1454} {
		mem[v] = 3
	}
	for _, v := range []int{871, 45361, 872, 45362} {
		mem[v] = 2
	}
	res := 0
	buf := make([]int, 0, 70)
	for n := 1; n < lim; n++ {
		if mem[n] != 0 {
			if mem[n] == 60 {
				res++
			}
			continue
		}
		buf = buf[:0]
		x := n
		for mem[x] == 0 {
			buf = append(buf, x)
			s := 0
			for y := x; y > 0; y /= 10 {
				s += f[y%10]
			}
			x = s
		}
		l := mem[x]
		for i := len(buf) - 1; i >= 0; i-- {
			l++
			mem[buf[i]] = l
		}
		if mem[n] == 60 {
			res++
		}
	}
	fmt.Println("The answer is: ", res)
}

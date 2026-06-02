package main

import (
	"fmt"
	"math/bits"
	"strconv"
)

var mem map[string]int

func h1(n int, digit int) int {
	if n < 0 {
		return 0
	}
	if n < 10 {
		if digit == 0 {
			return 1
		} else if n >= digit {
			return 1
		}
		return 0
	}
	key := strconv.Itoa(n) + "-" + strconv.Itoa(digit)
	if val, ok := mem[key]; ok {
		return val
	}
	s := strconv.Itoa(n)
	a := int(s[0] - '0')
	pow := 1
	for i := 1; i < len(s); i++ {
		pow *= 10
	}
	rem := n % pow
	count := 0
	if digit < a {
		count += pow
	} else if digit == a {
		count += rem + 1
	}
	count += a * (len(s) - 1) * (pow / 10)
	count += h1(rem, digit)
	if digit == 0 {
		p := pow
		for p > 1 {
			count -= p
			p /= 10
		}
	}
	mem[key] = count
	return count
}

func s(d int) int {
	sum := 0
	n := 0
	for n < 100000000000 {
		c := h1(n, d)
		if c == n {
			sum += n
			n++
		} else if c > n {
			step := (c - n) / bits.Len(uint(n))
			if step == 0 {
				step = 1
			}
			n += step
		} else {
			step := (n - c) / bits.Len(uint(n))
			if step == 0 {
				step = 1
			}
			n += step
		}
	}
	return sum
}

func main() {
	mem = make(map[string]int)
	res := 0
	for d := 1; d <= 9; d++ {
		res += s(d)
	}
	fmt.Println(res)
}

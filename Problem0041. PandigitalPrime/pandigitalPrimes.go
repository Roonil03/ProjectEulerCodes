package main

import (
	"fmt"
	"sort"
)

func main() {
	res := 0
	for i := 7; i >= 1; i-- {
		digits := make([]int, i)
		for a := 0; a < i; a++ {
			digits[a] = a + 1
		}
		var perms [][]int
		genPerms(digits, &perms, []int{})
		var n []int
		for _, perm := range perms {
			n = append(n, permToNum(perm))
		}
		sort.Sort(sort.Reverse(sort.IntSlice(n)))
		for _, num := range n {
			if isPrime(num) {
				res = num
				break
			}
		}
		if res > 0 {
			break
		}
	}
	fmt.Println(res)
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func genPerms(digits []int, res *[][]int, cur []int) {
	if len(cur) == len(digits) {
		perm := make([]int, len(cur))
		copy(perm, cur)
		*res = append(*res, perm)
		return
	}

	for _, digit := range digits {
		used := false
		for _, c := range cur {
			if c == digit {
				used = true
				break
			}
		}
		if !used {
			cur = append(cur, digit)
			genPerms(digits, res, cur)
			cur = cur[:len(cur)-1]
		}
	}
}

func permToNum(perm []int) int {
	num := 0
	for _, digit := range perm {
		num = num*10 + digit
	}
	return num
}

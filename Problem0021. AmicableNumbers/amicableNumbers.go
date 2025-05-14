package main

import "fmt"

const max = 1e5

var d []int

func main() {
	d = make([]int, max+1)
	for j := 2; j <= max; j++ {
		d[j] = 1
	}
	for i := 2; i <= max; i++ {
		for j := i * 2; j <= max; j += i {
			d[j] += i
		}
	}
	fmt.Println("The answer is:", amicable(10000))
}

func amicable(n int) int {
	sum := 0
	for a := 2; a < n; a++ {
		if b := d[a]; b > a && b <= n && d[b] == a {
			sum += a + b
		}
	}
	return sum
}

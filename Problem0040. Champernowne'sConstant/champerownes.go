package main

import "fmt"

func main() {
	pos := []int{1, 10, 100, 1000, 10000, 100000, 1000000}
	res := 1
	for _, i := range pos {
		dig := helper(i)
		res *= dig
	}
	fmt.Println("The result is: ", res)
}

func helper(n int) int {
	dig, count, s := 1, 9, 1
	for n > dig*count {
		n -= dig * count
		dig++
		count *= 10
		s *= 10
	}
	num := s + (n-1)/dig
	a := (n - 1) % dig
	str := fmt.Sprintf("%d", num)
	return int(str[a] - '0')
}

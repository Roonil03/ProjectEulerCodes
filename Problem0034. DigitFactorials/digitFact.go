package main

import "fmt"

func main() {
	fact := []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}
	//Upper bound calculated with calculator checking with variations of digits with 9, and trying a round number figure
	b := 2000000
	sum := 0
	for i := 10; i < b; i++ {
		a, num := 0, i
		for num > 0 {
			d := num % 10
			a += fact[d]
			num /= 10
		}
		if a == i {
			sum += i
		}
	}
	fmt.Print("The answer is: ", sum)
}

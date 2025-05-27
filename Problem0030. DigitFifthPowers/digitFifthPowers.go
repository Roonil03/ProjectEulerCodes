package main

import "fmt"

func main() {
	var pow5 [10]int
	for i := range pow5 {
		pow5[i] = i * i * i * i * i
	}
	up := 6 * pow5[9]
	total := 0

	for n := 2; n <= up; n++ {
		sum := 0
		for i := n; i > 0; i /= 10 {
			sum += pow5[i%10]
		}
		if sum == n {
			total += n
		}
	}
	fmt.Print("The answer is: ", total)
}

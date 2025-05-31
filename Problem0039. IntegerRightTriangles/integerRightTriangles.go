package main

import "fmt"

func countSolutions(p int) int {
	count := 0
	for a := 1; a <= p/3; a++ {
		numerator := p*p - 2*a*p
		denominator := 2 * (p - a)
		if denominator > 0 && numerator%denominator == 0 {
			b := numerator / denominator
			if a <= b {
				count++
			}
		}
	}
	return count
}

func main() {
	maxSolutions := 0
	bestP := 0
	for p := 1; p <= 1000; p++ {
		solutions := countSolutions(p)
		if solutions > maxSolutions {
			maxSolutions = solutions
			bestP = p
		}
	}
	fmt.Println(bestP)
}

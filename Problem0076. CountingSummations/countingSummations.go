package main

import "fmt"

func main() {
	tar := 100
	ways := make([]int, tar+1)
	ways[0] = 1
	for n := 1; n < tar; n++ {
		for i := n; i <= tar; i++ {
			ways[i] += ways[i-n]
		}
	}
	res := ways[tar]
	fmt.Println("The answer is: ", res)
}

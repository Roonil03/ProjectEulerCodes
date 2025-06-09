package main

import "fmt"

func main() {
	lim := 999999
	f := make([]int, lim+1)
	for p := 2; p <= lim; p++ {
		if f[p] == 0 {
			for m := p; m <= lim; m += p {
				f[m]++
			}
		}
	}
	for i := 2; i < lim-3; i++ {
		if f[i] == 4 && f[i+1] == 4 && f[i+2] == 4 && f[i+3] == 4 {
			fmt.Println("The answer is: ", i)
			return
		}
	}
}

package main

import (
	"fmt"
)

func main() {
	const N = 50
	var f [N + 1]uint64
	f[0] = 1
	for i := 1; i <= N; i++ {
		f[i] = f[i-1]
		if i >= 2 {
			f[i] += f[i-2]
		}
		if i >= 3 {
			f[i] += f[i-3]
		}
		if i >= 4 {
			f[i] += f[i-4]
		}
	}
	fmt.Println("The answer is:", f[N])
}

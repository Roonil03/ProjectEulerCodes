package main

import "fmt"

func main() {
	for x := 125874; ; /*x < max*/ x += 9 {
		ok := true
		for m := 2; m <= 6 && ok; m++ {
			if !sameDigits(x, m*x) {
				ok = false
			}
		}
		if ok {
			fmt.Println("The answer is: ", x)
		}
	}
}

func fingerprint(n int) (fp [10]byte) {
	for n > 0 {
		fp[n%10]++
		n /= 10
	}
	return
}

func sameDigits(a, b int) bool {
	return fingerprint(a) == fingerprint(b)
}

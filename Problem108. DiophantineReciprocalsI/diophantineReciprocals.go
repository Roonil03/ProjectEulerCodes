package main

import (
	"fmt"
)

func main() {
	P := []int{}
	isp := make([]bool, 501)
	for i := 2; i <= 500; i++ {
		if !isp[i] {
			P = append(P, i)
			for j := i * i; j <= 500; j += i {
				isp[j] = true
			}
		}
	}
	for n := 1; ; n++ {
		t, cnt := n, 1
		for _, p := range P {
			if p*p > t {
				break
			}
			e := 0
			for t%p == 0 {
				t /= p
				e++
			}
			if e > 0 {
				cnt *= 2*e + 1
				if cnt > 2000 {
					break
				}
			}
		}
		if t > 1 {
			cnt *= 3
		}
		if cnt > 2000 {
			fmt.Println("The answer is: ", n)
			return
		}
	}
}

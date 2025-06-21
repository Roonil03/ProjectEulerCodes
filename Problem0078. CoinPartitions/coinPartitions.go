package main

import "fmt"

func main() {
	tar, temp := 1000000, 100000
	part := make([]int, temp)
	part[0] = 1
	var pent []int
	for k := 1; k <= 500; k++ {
		pent1, pent2 := k*(3*k-1)/2, k*(3*k+1)/2
		if pent1 >= temp {
			break
		}
		pent = append(pent, pent1)
		if pent2 < temp {
			pent = append(pent, pent2)
		}
	}
	for n := 1; n < temp; n++ {
		sum := 0
		for i, p := range pent {
			if p > n {
				break
			}
			sign := 1 - (i & 2)
			sum += sign * part[n-p]
		}
		part[n] = ((sum % tar) + tar) % tar
		if part[n] == 0 {
			fmt.Println("The answer is: ", n)
			return
		}
	}
	fmt.Print("Increase temp")
}

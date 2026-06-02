package main

import "fmt"

func main() {
	var m uint64
	for n := uint64(1); n <= 26; n++ {
		r := uint64(1)
		k := n
		if k > 13 {
			k = 26 - n
		}
		for i := uint64(1); i <= k; i++ {
			r = r * (27 - i) / i
		}
		v := r * ((1 << n) - n - 1)
		if v > m {
			m = v
		}
	}
	fmt.Println(m)
}

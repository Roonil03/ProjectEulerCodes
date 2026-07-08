package main

import "fmt"

func main() {
	sq := []int64{4, 9, 25, 49, 121, 169, 289, 361, 529, 841, 961, 1369, 1681, 1849, 2209}
	d := make(map[int64]bool)
	c := make([]int64, 51)
	c[0] = 1
	d[1] = true
	for i := 1; i < 51; i++ {
		c[i] = 1
		for j := i - 1; j > 0; j-- {
			c[j] += c[j-1]
			d[c[j]] = true
		}
	}
	var s int64
	for k := range d {
		f := true
		for _, v := range sq {
			if k%v == 0 {
				f = false
				break
			}
		}
		if f {
			s += k
		}
	}
	fmt.Println(s)
}

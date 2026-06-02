package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	const n = 200000
	v2 := make([]int, n+1)
	v5 := make([]int, n+1)
	for i := 1; i <= n; i++ {
		c2 := 0
		tmp := i
		for tmp%2 == 0 {
			c2++
			tmp /= 2
		}
		v2[i] = v2[i-1] + c2
		c5 := 0
		tmp = i
		for tmp%5 == 0 {
			c5++
			tmp /= 5
		}
		v5[i] = v5[i-1] + c5
	}
	target2 := v2[n] - 12
	target5 := v5[n] - 12
	var res int64
	var wg sync.WaitGroup
	workers := runtime.NumCPU()
	var currentI int32 = 0
	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var temp int64
			for {
				i := int(atomic.AddInt32(&currentI, 1) - 1)
				if i > n/3 {
					break
				}
				v2_i := v2[i]
				v5_i := v5[i]
				maxJ := (n - i) / 2
				for j := i; j <= maxJ; j++ {
					k := n - i - j
					if v2_i+v2[j]+v2[k] <= target2 && v5_i+v5[j]+v5[k] <= target5 {
						if i == j || j == k {
							temp += 3
						} else {
							temp += 6
						}
					}
				}
			}
			atomic.AddInt64(&res, temp)
		}()
	}
	wg.Wait()
	fmt.Println(res)
}

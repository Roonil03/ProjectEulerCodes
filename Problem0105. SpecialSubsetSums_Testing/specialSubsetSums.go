package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	filepath := "Problem0105. SpecialSubsetSums_Testing/input.txt"
	f, _ := os.Open(filepath)
	defer f.Close()
	s := bufio.NewScanner(f)
	total := 0
	for s.Scan() {
		fs := strings.Split(s.Text(), ",")
		n := len(fs)
		a := make([]int, n)
		for i, v := range fs {
			a[i], _ = strconv.Atoi(v)
		}
		sort.Ints(a)
		ok := true
		for k := 1; k <= n/2; k++ {
			sumL, sumS := 0, 0
			for i := 0; i <= k; i++ {
				sumS += a[i]
			}
			for i := n - k; i < n; i++ {
				sumL += a[i]
			}
			if sumS <= sumL {
				ok = false
				break
			}
		}
		if !ok {
			continue
		}
		mm := map[int][]int{}
		lim := 1 << n
		dup := false
		for m := 1; m < lim; m++ {
			sum := 0
			for i := 0; i < n; i++ {
				if m&(1<<i) != 0 {
					sum += a[i]
				}
			}
			for _, pm := range mm[sum] {
				if pm&m == 0 {
					dup = true
					break
				}
			}
			if dup {
				break
			}
			mm[sum] = append(mm[sum], m)
		}
		if dup {
			continue
		}
		for _, v := range a {
			total += v
		}
	}
	fmt.Println("The answer is:", total)
}

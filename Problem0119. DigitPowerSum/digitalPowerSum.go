package main

import (
	"fmt"
	"math/big"
	"sort"
)

func main() {
	const target = 30
	maxBase := 200
	set := make(map[string]*big.Int)
	limit := new(big.Int).Exp(big.NewInt(10), big.NewInt(25), nil)

	for b := 2; b <= maxBase; b++ {
		base := big.NewInt(int64(b))
		v := new(big.Int).Set(base)
		for e := 2; ; e++ {
			v.Mul(v, base)
			if v.Cmp(limit) > 0 {
				break
			}
			if v.Cmp(big.NewInt(10)) < 0 {
				continue
			}
			if digitSum(v) == b {
				set[v.String()] = new(big.Int).Set(v)
			}
		}
	}
	candidates := make([]*big.Int, 0, len(set))
	for _, v := range set {
		candidates = append(candidates, v)
	}
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].Cmp(candidates[j]) < 0
	})
	fmt.Println("The answer is:", candidates[target-1])
}

func digitSum(n *big.Int) int {
	sum := 0
	for _, d := range n.String() {
		sum += int(d - '0')
	}
	return sum
}

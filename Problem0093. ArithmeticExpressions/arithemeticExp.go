package main

import (
	"fmt"
	"sort"
)

type rational struct{ p, q int }

func newRat(p, q int) rational {
	if q < 0 {
		p, q = -p, -q
	}
	a, b := p, q
	if a < 0 {
		a = -a
	}
	for b != 0 {
		a, b = b, a%b
	}
	g := a
	return rational{p / g, q / g}
}

func (r rational) add(s rational) rational { return newRat(r.p*s.q+s.p*r.q, r.q*s.q) }
func (r rational) sub(s rational) rational { return newRat(r.p*s.q-s.p*r.q, r.q*s.q) }
func (r rational) mul(s rational) rational { return newRat(r.p*s.p, r.q*s.q) }
func (r rational) div(s rational) (rational, bool) {
	if s.p == 0 {
		return rational{}, false
	}
	return newRat(r.p*s.q, r.q*s.p), true
}

func generateResults(nums []rational, results map[int]bool) {
	n := len(nums)
	if n == 1 {
		if nums[0].q == 1 && nums[0].p > 0 {
			results[nums[0].p] = true
		}
		return
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			a, b := nums[i], nums[j]
			rem := make([]rational, 0, n-1)
			for k := 0; k < n; k++ {
				if k != i && k != j {
					rem = append(rem, nums[k])
				}
			}
			ops := []rational{a.add(b), a.sub(b), b.sub(a), a.mul(b)}
			for _, r := range ops {
				generateResults(append(rem, r), results)
			}
			if r, ok := a.div(b); ok {
				generateResults(append(rem, r), results)
			}
			if r, ok := b.div(a); ok {
				generateResults(append(rem, r), results)
			}
		}
	}
}

func main() {
	bestLen, bestA, bestB, bestC, bestD := 0, 0, 0, 0, 0
	for a := 1; a <= 6; a++ {
		for b := a + 1; b <= 7; b++ {
			for c := b + 1; c <= 8; c++ {
				for d := c + 1; d <= 9; d++ {
					reachable := make(map[int]bool)
					digits := []int{a, b, c, d}
					idx := []int{0, 1, 2, 3}
					for {
						nums := make([]rational, 4)
						for i, id := range idx {
							nums[i] = rational{digits[id], 1}
						}
						generateResults(nums, reachable)
						k := 2
						for k >= 0 && idx[k] > idx[k+1] {
							k--
						}
						if k < 0 {
							break
						}
						l := 3
						for idx[k] > idx[l] {
							l--
						}
						idx[k], idx[l] = idx[l], idx[k]
						for i, j := k+1, 3; i < j; i, j = i+1, j-1 {
							idx[i], idx[j] = idx[j], idx[i]
						}
					}
					seq := make([]int, 0, len(reachable))
					for v := range reachable {
						seq = append(seq, v)
					}
					sort.Ints(seq)
					n := 1
					for _, v := range seq {
						if v == n {
							n++
						} else if v > n {
							break
						}
					}
					if n-1 > bestLen {
						bestLen, bestA, bestB, bestC, bestD = n-1, a, b, c, d
					}
				}
			}
		}
	}
	fmt.Printf("The answer is: %d%d%d%d\n", bestA, bestB, bestC, bestD)
}

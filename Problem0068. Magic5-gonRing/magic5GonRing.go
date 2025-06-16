package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Ring struct {
	outer [5]int
	inner [5]int
	sum   int
}

func main() {
	res := solve()
	fmt.Println("The answer is:", res)
}

func solve() string {
	res := ""
	for ms := 12; ms <= 19; ms++ {
		tos := 110 - 5*ms
		if tos <= 0 {
			continue
		}
		var vo [][]int
		cand := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		comb := genComb(cand, 4)
		for _, c := range comb {
			or := append(c, 10)
			s := 0
			for _, v := range or {
				s += v
			}
			if s == tos {
				perms := genPerm(or)
				vo = append(vo, perms...)
			}
		}
		for _, out := range vo {
			ic := []int{}
			used := make(map[int]bool)
			for _, v := range out {
				used[v] = true
			}
			for i := 1; i <= 10; i++ {
				if !used[i] {
					ic = append(ic, i)
				}
			}
			r := solveInner(out, ic, ms)
			if r != "" && len(r) == 16 {
				if r > res {
					res = r
				}
			}
		}
	}
	return res
}

func (r *Ring) isValid() bool {
	for i := 0; i < 5; i++ {
		next := (i + 1) % 5
		if r.outer[i]+r.inner[i]+r.inner[next] != r.sum {
			return false
		}
	}
	return true
}

func (r *Ring) toString() string {
	mp := 0
	for i := 1; i < 5; i++ {
		if r.outer[i] < r.outer[mp] {
			mp = i
		}
	}
	var res strings.Builder
	for i := 0; i < 5; i++ {
		pos := (mp + i) % 5
		nxt := (pos + 1) % 5
		res.WriteString(strconv.Itoa(r.outer[pos]))
		res.WriteString(strconv.Itoa(r.inner[pos]))
		res.WriteString(strconv.Itoa(r.inner[nxt]))
	}
	return res.String()
}

func genComb(nums []int, k int) [][]int {
	var res [][]int
	var curr []int
	var bt func(start int)
	bt = func(start int) {
		if len(curr) == k {
			c := make([]int, k)
			copy(c, curr)
			res = append(res, c)
			return
		}
		for i := start; i < len(nums); i++ {
			curr = append(curr, nums[i])
			bt(i + 1)
			curr = curr[:len(curr)-1]
		}
	}
	bt(0)
	return res
}

func genPerm(arr []int) [][]int {
	var res [][]int
	var bt func(curr []int, used []bool)
	bt = func(curr []int, used []bool) {
		if len(curr) == len(arr) {
			p := make([]int, len(curr))
			copy(p, curr)
			res = append(res, p)
			return
		}
		for i := 0; i < len(arr); i++ {
			if !used[i] {
				used[i] = true
				curr = append(curr, arr[i])
				bt(curr, used)
				curr = curr[:len(curr)-1]
				used[i] = false
			}
		}
	}
	bt([]int{}, make([]bool, len(arr)))
	return res
}

func solveInner(out []int, ic []int, ms int) string {
	var res []string
	var bt func(inn []int, used []bool, pos int)
	bt = func(inn []int, used []bool, pos int) {
		if pos == 5 {
			ring := Ring{sum: ms}
			copy(ring.outer[:], out)
			copy(ring.inner[:], inn)
			if ring.isValid() {
				result := ring.toString()
				if len(result) == 16 {
					res = append(res, result)
				}
			}
			return
		}
		for i, cand := range ic {
			if !used[i] {
				used[i] = true
				inn[pos] = cand
				bt(inn, used, pos+1)
				used[i] = false
			}
		}
	}
	inn := make([]int, 5)
	used := make([]bool, len(ic))
	bt(inn, used, 0)
	if len(res) > 0 {
		sort.Strings(res)
		return res[len(res)-1]
	}
	return ""
}

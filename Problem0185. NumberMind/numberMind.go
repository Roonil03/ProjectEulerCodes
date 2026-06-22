package main

import (
	"fmt"
	"strings"
)

type State [22]int8

var (
	tar      [22]int8
	cond     [22]string
	leftMap  map[State][]uint32
	result   string
	valid    [16][10]bool
	maxLeft  [22]int8
	maxRight [22]int8
)

func dfsLeft(id int, val uint32, cur State, poss State) {
	if id == 8 {
		leftMap[cur] = append(leftMap[cur], val)
		return
	}
	for d := byte('0'); d <= byte('9'); d++ {
		if !valid[id][d-'0'] {
			continue
		}
		fg := true
		var nCur, nPoss State
		for i := 0; i < 22; i++ {
			nCur[i], nPoss[i] = cur[i], poss[i]
			req := cond[i][id]
			if d == req {
				nCur[i]++
			}
			if valid[id][req-'0'] {
				nPoss[i]--
			}
			if nCur[i] > tar[i] || nCur[i]+nPoss[i]+maxRight[i] < tar[i] {
				fg = false
				break
			}
		}
		if fg {
			dfsLeft(id+1, val*10+uint32(d-'0'), nCur, nPoss)
		}
	}
}

func dfsRight(id int, val uint32, cur State, poss State) {
	if result != "" {
		return
	}
	if id == 16 {
		var req State
		for i := 0; i < 22; i++ {
			req[i] = tar[i] - cur[i]
		}
		if lefts, ok := leftMap[req]; ok {
			for _, l := range lefts {
				result = fmt.Sprintf("%08d%08d", l, val)
				return
			}
		}
		return
	}
	for d := byte('0'); d <= byte('9'); d++ {
		if !valid[id][d-'0'] {
			continue
		}
		fg := true
		var nCur, nPoss State
		for i := 0; i < 22; i++ {
			nCur[i], nPoss[i] = cur[i], poss[i]
			req := cond[i][id]
			if d == req {
				nCur[i]++
			}
			if valid[id][req-'0'] {
				nPoss[i]--
			}
			if nCur[i] > tar[i] || nCur[i]+nPoss[i]+maxLeft[i] < tar[i] {
				fg = false
				break
			}
		}
		if fg {
			dfsRight(id+1, val*10+uint32(d-'0'), nCur, nPoss)
		}
	}
}

func main() {
	input := `5616185650518293 ;2
3847439647293047 ;1
5855462940810587 ;3
9742855507068353 ;3
4296849643607543 ;3
3174248439465858 ;1
4513559094146117 ;2
7890971548908067 ;3
8157356344118483 ;1
2615250744386899 ;2
8690095851526254 ;3
6375711915077050 ;1
6913859173121360 ;1
6442889055042768 ;2
2321386104303845 ;0
2326509471271448 ;2
5251583379644322 ;2
1748270476758276 ;3
4895722652190306 ;1
3041631117224635 ;3
1841236454324589 ;3
2659862637316867 ;2`

	lines := strings.Split(strings.TrimSpace(input), "\n")
	for i, v := range lines {
		parts := strings.Split(strings.TrimSpace(v), " ;")
		cond[i] = parts[0]
		var s int
		fmt.Sscanf(parts[1], "%d", &s)
		tar[i] = int8(s)
	}
	for i := range 16 {
		for d := range 10 {
			valid[i][d] = true
		}
	}
	for i := range 22 {
		if tar[i] == 0 {
			for c := 0; c < 16; c++ {
				valid[c][cond[i][c]-'0'] = false
			}
		}
	}
	for i := range 22 {
		for c := range 8 {
			if valid[c][cond[i][c]-'0'] {
				maxLeft[i]++
			}
		}
		for c := 8; c < 16; c++ {
			if valid[c][cond[i][c]-'0'] {
				maxRight[i]++
			}
		}
	}
	leftMap = make(map[State][]uint32)
	dfsLeft(0, 0, State{}, maxLeft)
	dfsRight(8, 0, State{}, maxRight)
	fmt.Println(result)
}

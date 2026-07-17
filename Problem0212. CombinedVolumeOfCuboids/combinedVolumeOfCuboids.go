package main

import "fmt"

const MAX = 10405

var count [65536]int
var length [65536]int

func update(node, l, r, ql, qr, val int) {
	if ql <= l && r <= qr {
		count[node] += val
	} else {
		mid := (l + r) >> 1
		if ql < mid {
			update(node<<1, l, mid, ql, qr, val)
		}
		if qr > mid {
			update(node<<1|1, mid, r, ql, qr, val)
		}
	}
	if count[node] > 0 {
		length[node] = r - l
	} else if r-l == 1 {
		length[node] = 0
	} else {
		length[node] = length[node<<1] + length[node<<1|1]
	}
}

type Rect struct {
	x1, x2, y1, y2 int
}

type Event struct {
	y1, y2, val int
}

func main() {
	ss := make([]int64, 300002)
	for i := int64(1); i <= 55; i++ {
		v := (100003 - 200003*i + 300007*i*i*i) % 1000000
		if v < 0 {
			v += 1000000
		}
		ss[i] = v
	}
	for i := 56; i <= 300000; i++ {
		ss[i] = (ss[i-24] + ss[i-55]) % 1000000
	}
	activeAtZ := make([][]Rect, MAX)
	for i := 1; i <= 50000; i++ {
		x0 := int(ss[6*i-5] % 10000)
		y0 := int(ss[6*i-4] % 10000)
		z0 := int(ss[6*i-3] % 10000)
		dx := int(1 + (ss[6*i-2] % 399))
		dy := int(1 + (ss[6*i-1] % 399))
		dz := int(1 + (ss[6*i] % 399))
		r := Rect{x0, x0 + dx, y0, y0 + dy}
		for j := z0; j < z0+dz; j++ {
			activeAtZ[j] = append(activeAtZ[j], r)
		}
	}
	var eventsAtX [MAX][]Event
	res := int64(0)
	for i := range MAX {
		if len(activeAtZ[i]) == 0 {
			continue
		}
		maxXThisZ := 0
		for _, v := range activeAtZ[i] {
			eventsAtX[v.x1] = append(eventsAtX[v.x1], Event{v.y1, v.y2, 1})
			eventsAtX[v.x2] = append(eventsAtX[v.x2], Event{v.y1, v.y2, -1})
			if v.x2 > maxXThisZ {
				maxXThisZ = v.x2
			}
		}
		area := int64(0)
		for j := 0; j <= maxXThisZ; j++ {
			for _, ev := range eventsAtX[j] {
				update(1, 0, MAX, ev.y1, ev.y2, ev.val)
			}
			area += int64(length[1])
			eventsAtX[j] = eventsAtX[j][:0]
		}
		res += area
	}
	fmt.Print(res)
}

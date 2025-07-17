package main

import "fmt"

func main() {
	const n = 2000
	var g [55]int64
	var k int64 = 1
	var m [n][n]int64
	idx := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if k <= 55 {
				val := (100003-200003*k+300007*k*k*k)%1000000 - 500000
				g[(k-1)%55] = val
				m[i][j] = val
				k++
			} else {
				val := (g[(k-24-1)%55]+g[(k-55-1)%55]+1000000)%1000000 - 500000
				g[(k-1)%55] = val
				m[i][j] = val
				k++
			}
			idx++
		}
	}
	var mx int64
	for i := 0; i < n; i++ {
		var c, s int64
		for j := 0; j < n; j++ {
			c += m[i][j]
			if c > s {
				s = c
			}
			if c < 0 {
				c = 0
			}
		}
		if s > mx {
			mx = s
		}
	}
	for j := 0; j < n; j++ {
		var c, s int64
		for i := 0; i < n; i++ {
			c += m[i][j]
			if c > s {
				s = c
			}
			if c < 0 {
				c = 0
			}
		}
		if s > mx {
			mx = s
		}
	}
	for i := 0; i < n; i++ {
		var c, s int64
		for x, y := 0, i; x < n && y < n; x, y = x+1, y+1 {
			c += m[x][y]
			if c > s {
				s = c
			}
			if c < 0 {
				c = 0
			}
		}
		if s > mx {
			mx = s
		}
	}
	for j := 1; j < n; j++ {
		var c, s int64
		for x, y := j, 0; x < n && y < n; x, y = x+1, y+1 {
			c += m[x][y]
			if c > s {
				s = c
			}
			if c < 0 {
				c = 0
			}
		}
		if s > mx {
			mx = s
		}
	}
	for i := 0; i < n; i++ {
		var c, s int64
		for x, y := 0, i; x < n && y >= 0; x, y = x+1, y-1 {
			c += m[x][y]
			if c > s {
				s = c
			}
			if c < 0 {
				c = 0
			}
		}
		if s > mx {
			mx = s
		}
	}
	for j := n - 2; j >= 0; j-- {
		var c, s int64
		for x, y := j, n-1; x < n && y >= 0; x, y = x+1, y-1 {
			c += m[x][y]
			if c > s {
				s = c
			}
			if c < 0 {
				c = 0
			}
		}
		if s > mx {
			mx = s
		}
	}
	fmt.Println("The answer is:", mx)
}

package main

import "fmt"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	const N = 50
	total := 0
	total += N * N * 3
	for x := 1; x <= N; x++ {
		for y := 1; y <= N; y++ {
			d := gcd(x, y)
			dx := x / d
			dy := y / d
			count := 0
			qx, qy := x-dy, y+dx
			for qx >= 0 && qy <= N {
				count++
				qx -= dy
				qy += dx
			}
			qx, qy = x+dy, y-dx
			for qy >= 0 && qx <= N {
				count++
				qx += dy
				qy -= dx
			}
			total += count
		}
	}
	fmt.Println("The answer is:", total)
}

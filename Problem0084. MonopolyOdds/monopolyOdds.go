package main

import (
	"fmt"
	"math"
	"sort"
)

const (
	S       = 40
	N       = 4
	MaxIter = 10000
	Eps     = 1e-12
)

func main() {
	T := buildTransition()
	pi := stationary(T)
	type sq struct {
		i int
		p float64
	}
	arr := make([]sq, S)
	for i := 0; i < S; i++ {
		arr[i] = sq{i, pi[i]}
	}
	sort.Slice(arr, func(a, b int) bool {
		return arr[a].p > arr[b].p
	})
	fmt.Printf("The answer is: %02d%02d%02d\n", arr[0].i, arr[1].i, arr[2].i)
}

func buildTransition() [S][S]float64 {
	dice := make([]float64, 2*N+1)
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			dice[i+j] += 1.0
		}
	}
	for d := range dice {
		dice[d] /= float64(N * N)
	}
	var T [S][S]float64
	for i := 0; i < S; i++ {
		for sum, p := range dice {
			if p == 0 {
				continue
			}
			pos := (i + sum) % S
			if pos == 30 {
				T[i][10] += p
				continue
			}
			switch pos {
			case 2, 17, 33:
				for _, d := range []int{0, 10} {
					T[i][d] += p / 16
				}
				T[i][pos] += p * 14 / 16
			case 7, 22, 36:
				cards := []int{0, 10, 11, 24, 39, 5}
				r := nextR(pos)
				cards = append(cards, r, r, nextU(pos), (pos-3+S)%S)
				for len(cards) < 16 {
					cards = append(cards, pos)
				}
				for _, d := range cards {
					T[i][d] += p / 16
				}
			default:
				T[i][pos] += p
			}
		}
	}
	return T
}

func nextR(k int) int {
	for d := 1; d <= S; d++ {
		t := (k + d) % S
		if t%10 == 5 {
			return t
		}
	}
	return k
}

func nextU(k int) int {
	for d := 1; d <= S; d++ {
		t := (k + d) % S
		if t == 12 || t == 28 {
			return t
		}
	}
	return k
}

func stationary(T [S][S]float64) []float64 {
	pi := make([]float64, S)
	pi[0] = 1.0
	tmp := make([]float64, S)
	for it := 0; it < MaxIter; it++ {
		for i := 0; i < S; i++ {
			tmp[i] = 0
			for j := 0; j < S; j++ {
				tmp[i] += pi[j] * T[j][i]
			}
		}
		diff := 0.0
		for i := 0; i < S; i++ {
			diff += math.Abs(tmp[i] - pi[i])
			pi[i] = tmp[i]
		}
		if diff < Eps {
			break
		}
	}
	return pi
}

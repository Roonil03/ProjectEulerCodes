package main

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/dsp/fourier"
)

const (
	H = 8
	W = 11
	K = 1e12
)

func powλ(p, q int64) complex128 {
	// eigenvalue λ = 2(cos(2πp/H)+cos(2πq/W))
	lam := 2 * (math.Cos(2*math.Pi*float64(p)/H) +
		math.Cos(2*math.Pi*float64(q)/W))
	// fast exponentiation
	z := complex(lam, 0)
	var res complex128 = 1
	exp := int64(K)
	for exp > 0 {
		if exp&1 == 1 {
			res *= z
		}
		z *= z
		exp >>= 1
	}
	return res
}

func main() {
	// Initial grid from image (row-major)
	init := [][]float64{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},                         // row 0
		{0, 'B' - 'A', 0, 0, 0, 0, 0, 0, 0, 0, 'd' - 'a'},         // row 1
		{0, 'A' - 'A', 'E' - 'A', 0, 0, 0, 0, 0, 0, 'C' - 'A', 0}, // row 2 ← added trailing 0
		{0, 'D' - 'A', 0, 0, 0, 0, 0, 0, 0, 0, 0},                 // row 3
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},                         // row 4
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},                         // row 5
		{0, 'e' - 'a', 0, 0, 0, 0, 0, 0, 0, 'c' - 'a', 'a' - 'a'}, // row 6
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},                         // row 7
	}
	// Flatten rows into complex slice for FFT
	data := make([]complex128, H*W)
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			data[i*W+j] = complex(init[i][j], 0)
		}
	}

	// 2D FFT: first dim H then W
	fftH := fourier.NewCmplxFFT(H)
	for j := 0; j < W; j++ {
		col := make([]complex128, H)
		for i := 0; i < H; i++ {
			col[i] = data[i*W+j]
		}
		fftH.Coefficients(col, col)
		for i := 0; i < H; i++ {
			data[i*W+j] = col[i]
		}
	}
	fftW := fourier.NewCmplxFFT(W)
	for i := 0; i < H; i++ {
		row := data[i*W : (i+1)*W]
		fftW.Coefficients(row, row)
	}

	// Multiply each frequency by λ^{K}
	for p := int64(0); p < H; p++ {
		for q := int64(0); q < W; q++ {
			idx := int(p)*W + int(q)
			data[idx] *= powλ(p, q)
		}
	}

	// Inverse 2D FFT
	for i := 0; i < H; i++ {
		row := data[i*W : (i+1)*W]
		fftW.Sequence(row, row)
		for j := 0; j < W; j++ {
			data[i*W+j] /= complex(W, 0)
		}
	}
	for j := 0; j < W; j++ {
		col := make([]complex128, H)
		for i := 0; i < H; i++ {
			col[i] = data[i*W+j]
		}
		fftH.Sequence(col, col)
		for i := 0; i < H; i++ {
			data[i*W+j] = col[i] / complex(H, 0)
		}
	}

	// Round, mod 7, map to letters
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			v := int(math.Round(real(data[i*W+j])))
			c := byte((v%7+7)%7) + 'A'
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}
}

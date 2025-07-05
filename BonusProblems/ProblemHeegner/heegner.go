package main

import (
	"container/heap"
	"fmt"
	"math"
	"math/big"

	"github.com/ALTree/bigfloat"
)

func isSquare(n int) bool {
	r := int(math.Sqrt(float64(n)) + 0.5)
	return r*r == n
}

// distance to nearest integer for a *big.Float
func fracDistance(f *big.Float) *big.Float {
	// floor = trunc toward −Inf for positive f, use Int for trunc toward zero
	flo, _ := f.Int(nil)
	fi := new(big.Float).SetPrec(f.Prec()).SetInt(flo)
	// compute ceil
	ci := new(big.Int).Set(flo)
	if f.Cmp(fi) > 0 {
		ci.Add(ci, big.NewInt(1))
	}
	cf := new(big.Float).SetPrec(f.Prec()).SetInt(ci)
	a := new(big.Float).SetPrec(f.Prec()).Sub(f, fi)
	b := new(big.Float).SetPrec(f.Prec()).Sub(cf, f)
	if a.Cmp(b) < 0 {
		return a
	}
	return b
}

// compute cosh(x) = (e^x + e^(−x)) / 2 using ALTree/bigfloat Exp
func coshHighPrec(x *big.Float) *big.Float {
	pos := bigfloat.Exp(x)
	neg := bigfloat.Exp(new(big.Float).SetPrec(x.Prec()).Neg(x))
	sum := new(big.Float).SetPrec(x.Prec()).Add(pos, neg)
	return sum.Quo(sum, big.NewFloat(2).SetPrec(x.Prec()))
}

// heap item for (distance, n)
type item struct {
	d *big.Float
	n int
}
type minPQ []*item

func (pq minPQ) Len() int            { return len(pq) }
func (pq minPQ) Less(i, j int) bool  { return pq[i].d.Cmp(pq[j].d) < 0 }
func (pq minPQ) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *minPQ) Push(x interface{}) { *pq = append(*pq, x.(*item)) }
func (pq *minPQ) Pop() interface{} {
	old := *pq
	it := old[len(old)-1]
	*pq = old[:len(old)-1]
	return it
}

func main() {
	// high precision
	prec := uint(256)
	pi, _ := new(big.Float).SetPrec(prec).SetString(
		"3.141592653589793238462643383279502884197169399375105820974944")

	// initialize min-heap
	pq := &minPQ{}
	heap.Init(pq)

	for n := -1000; n <= 1000; n++ {
		if n == 0 {
			continue
		}
		absN := n
		if absN < 0 {
			absN = -absN
		}
		if isSquare(absN) {
			continue
		}

		var dist *big.Float
		if n > 0 {
			// cos(π√n)
			c := math.Cos(math.Pi * math.Sqrt(float64(n)))
			cf := new(big.Float).SetPrec(prec).SetFloat64(c)
			dist = fracDistance(cf)
		} else {
			// cosh(π√|n|)
			root := new(big.Float).SetPrec(prec).
				Sqrt(new(big.Float).SetPrec(prec).SetFloat64(float64(absN)))
			arg := new(big.Float).SetPrec(prec).Mul(pi, root)
			ch := coshHighPrec(arg)
			dist = fracDistance(ch)
		}
		heap.Push(pq, &item{dist, n})
	}

	best := heap.Pop(pq).(*item)
	fmt.Println(best.n)
}

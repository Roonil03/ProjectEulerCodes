package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	var res int64 = 0
	bound := int64(1000000000000)
	pk := new(big.Int)
	qk := new(big.Int)
	pp := new(big.Int)
	qp := new(big.Int)
	mnum := new(big.Int)
	mden := new(big.Int)
	R := new(big.Int)
	L := new(big.Int)
	two := big.NewInt(2)
	bigN := new(big.Int)
	temp := new(big.Int)
	for i := 2; i <= 100000; i++ {
		a0 := int64(math.Sqrt(float64(i)))
		if a0*a0 == int64(i) {
			continue
		}
		m, d, a := int64(0), int64(1), a0
		p2, p1 := int64(1), a0
		q2, q1 := int64(0), int64(1)
		k := 0
		for {
			m = d*a - m
			d = (int64(i) - m*m) / d
			a = (a0 + m) / d
			qn := a*q1 + q2
			if qn > bound {
				break
			}
			pn := a*p1 + p2
			p2, p1 = p1, pn
			q2, q1 = q1, qn
			k++
		}
		cm := (bound - q2) / q1
		qp2 := cm*q1 + q2
		pp2 := cm*p1 + p2
		pk.SetInt64(p1)
		qk.SetInt64(q1)
		pp.SetInt64(pp2)
		qp.SetInt64(qp2)
		mnum.Mul(pk, qp)
		temp.Mul(pp, qk)
		mnum.Add(mnum, temp)
		mden.Mul(qk, qp)
		mden.Mul(mden, two)
		R.Mul(mnum, mnum)
		L.Mul(mden, mden)
		bigN.SetInt64(int64(i))
		L.Mul(L, bigN)
		cmp := L.Cmp(R)
		if k%2 == 0 {
			if cmp > 0 {
				res += qp2
			} else {
				res += q1
			}
		} else {
			if cmp > 0 {
				res += q1
			} else {
				res += qp2
			}
		}
	}
	fmt.Println(res)
}

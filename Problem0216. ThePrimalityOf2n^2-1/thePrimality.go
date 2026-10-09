package main

import "fmt"

const lim = 50000000

func pow(a, e, m int64) int64 {
	r := int64(1)
	for e > 0 {
		if e&1 > 0 {
			r = r * a % m
		}
		a = a * a % m
		e >>= 1
	}
	return r
}

func root(p int64) int64 {
	if p&7 == 7 {
		x := pow(2, (p+1)/4, p)
		if x&1 > 0 {
			return (x + p) / 2
		}
		return x / 2
	}
	q := p - 1
	s := 0
	for q&1 == 0 {
		q >>= 1
		s++
	}
	z := int64(3)
	for pow(z, (p-1)/2, p) == 1 {
		z++
	}
	c := pow(z, q, p)
	x := pow(2, (q+1)/2, p)
	t := pow(2, q, p)
	for t != 1 {
		i := 1
		u := t * t % p
		for u != 1 {
			u = u * u % p
			i++
		}
		b := pow(c, 1<<uint(s-i-1), p)
		x = x * b % p
		c = b * b % p
		t = t * c % p
		s = i
	}
	if x&1 > 0 {
		return (x + p) / 2
	}
	return x / 2
}

func main() {
	l := 70710679
	p := make([]bool, l+1)
	for i := 2; i*i <= l; i++ {
		if !p[i] {
			for j := i * i; j <= l; j += i {
				p[j] = true
			}
		}
	}
	p[0], p[1] = true, true
	b := make([]bool, lim+1)
	for i := 3; i <= l; i++ {
		if p[i] || i&7 != 1 && i&7 != 7 {
			continue
		}
		r := int(root(int64(i)))
		for _, x := range []int{r, i - r} {
			for x <= lim {
				b[x] = true
				x += i
			}
		}
	}
	for n := 2; 2*n*n-1 <= l; n++ {
		if !p[2*n*n-1] {
			b[n] = false
		}
	}
	res := 0
	for n := 2; n <= lim; n++ {
		if !b[n] {
			res++
		}
	}
	fmt.Println(res)
}

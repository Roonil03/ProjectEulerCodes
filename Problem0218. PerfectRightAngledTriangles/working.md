Let a primitive right triangle be:

$$a = m^2 - n^2, \quad b = 2mn, \quad c = m^2 + n^2$$

with $\gcd(m, n) = 1$ and opposite parity. Its area is:

$$A = \frac{ab}{2} = mn(m^2 - n^2)$$

Because $c$ is a perfect square, write:

$$m^2 + n^2 = d^2$$

So $(m, n, d)$ is itself a primitive Pythagorean triple.

Now prove that $84 = 6 \cdot 28$ always divides $A$:

* One of $m, n$ is the even leg of a primitive Pythagorean triple, so it is divisible by $4$. Hence $4 \mid mn$.
* Modulo $3$, if neither $m$ nor $n$ were divisible by $3$, then $m^2 + n^2 \equiv 2 \pmod 3$, which cannot be a square. Thus $3 \mid mn$.
* Modulo $7$, the condition $m^2 + n^2 = d^2$ implies that either $7 \mid m$, $7 \mid n$, or $m^2 \equiv n^2 \pmod 7$.

Therefore:

$$7 \mid mn(m^2 - n^2)$$

Thus:

$$84 \mid A$$

Every perfect right-angled triangle is automatically super-perfect, regardless of the upper limit on $c$.
Therefore, the number that are not super-perfect is `0`.
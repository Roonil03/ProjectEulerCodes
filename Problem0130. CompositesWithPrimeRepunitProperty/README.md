# Composites with Prime Repunit Property
Problem 130
### Description:
<p>A number consisting entirely of ones is called a repunit. We shall define $R(k)$ to be a repunit of length $k$; for example, $R(6) = 111111$.</p>
<p>Given that $n$ is a positive integer and $\gcd(n, 10) = 1$, it can be shown that there always exists a value, $k$, for which $R(k)$ is divisible by $n$, and let $A(n)$ be the least such value of $k$; for example, $A(7) = 6$ and $A(41) = 5$.</p>
<p>You are given that for all primes, $p \gt 5$, that $p - 1$ is divisible by $A(p)$. For example, when $p = 41$, $A(41) = 5$, and $40$ is divisible by $5$.</p>
<p>However, there are rare composite values for which this is also true; the first five examples being $91$, $259$, $451$, $481$, and $703$.</p>
<p>Find the sum of the first twenty-five composite values of $n$ for which $\gcd(n, 10) = 1$ and $n - 1$ is divisible by $A(n)$.</p>


### Answer:
```
149253
```

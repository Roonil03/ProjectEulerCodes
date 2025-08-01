# abc-hits
Problem 127
### Description:
<p>The radical of $n$, $\operatorname{rad}(n)$, is the product of distinct prime factors of $n$. For example, $504 = 2^3 \times 3^2 \times 7$, so $\operatorname{rad}(504) = 2 \times 3 \times 7 = 42$.</p>
<p>We shall define the triplet of positive integers $(a, b, c)$ to be an abc-hit if:</p>
<ol><li>$\gcd(a, b) = \gcd(a, c) = \gcd(b, c) = 1$</li>
<li>$a \lt b$</li>
<li>$a + b = c$</li>
<li>$\operatorname{rad}(abc) \lt c$</li>
</ol><p>For example, $(5, 27, 32)$ is an abc-hit, because:</p>
<ol><li>$\gcd(5, 27) = \gcd(5, 32) = \gcd(27, 32) = 1$</li>
<li>$5 \lt 27$</li>
<li>$5 + 27 = 32$</li>
<li>$\operatorname{rad}(4320) = 30 \lt 32$</li>
</ol><p>It turns out that abc-hits are quite rare and there are only thirty-one abc-hits for $c \lt 1000$, with $\sum c = 12523$.</p>
<p>Find $\sum c$ for $c \lt 120000$.</p>

### Answer:
```
18407904
```
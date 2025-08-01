# Hexagonal Tile Differences
Problem 128
### Description:
<p>A hexagonal tile with number $1$ is surrounded by a ring of six hexagonal tiles, starting at "12 o'clock" and numbering the tiles $2$ to $7$ in an anti-clockwise direction.</p>
<p>New rings are added in the same fashion, with the next rings being numbered $8$ to $19$, $20$ to $37$, $38$ to $61$, and so on. The diagram below shows the first three rings.</p>
<div class="center">
<img src="https://projecteuler.net/resources/images/0128.png?1678992052" class="dark_img" alt=""></div>
<p>By finding the difference between tile $n$ and each of its six neighbours we shall define $\operatorname{PD}(n)$ to be the number of those differences which are prime.</p>
<p>For example, working clockwise around tile $8$ the differences are $12, 29, 11, 6, 1$, and $13$. So $\operatorname{PD}(8) = 3$.</p>
<p>In the same way, the differences around tile $17$ are $1, 17, 16, 1, 11$, and $10$, hence $\operatorname{PD}(17) = 2$.</p>
<p>It can be shown that the maximum value of $\operatorname{PD}(n)$ is $3$.</p>
<p>If all of the tiles for which $\operatorname{PD}(n) = 3$ are listed in ascending order to form a sequence, the $10$th tile would be $271$.</p>
<p>Find the $2000$th tile in this sequence.</p>

### Answer:
```
14516824220
```
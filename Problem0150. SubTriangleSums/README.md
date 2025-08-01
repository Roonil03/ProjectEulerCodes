# Sub-triangle Sums
Problem 150
### Description:
In a triangular array of positive and negative integers, we wish to find a sub-triangle such that the sum of the numbers it contains is the smallest possible.

In the example below, it can be easily verified that the marked triangle satisfies this condition having a sum of −42.

<img src="https://projecteuler.net/resources/images/0150.gif?1678992055"><br>

We wish to make such a triangular array with one thousand rows, so we generate 500500 pseudo-random numbers sk in the range ±219, using a type of random number generator (known as a Linear Congruential Generator) as follows:t := 0
for k = 1 up to k = 500500:
    t := (615949*t + 797807) \text{ modulo } 2^{20}
    s_k := t - 2^{19}

Thus: $s_1 = 273519$, $s_2 = -153582$, $s_3 = 450905$ etc

Our triangular array is then formed using the pseudo-random numbers thus:

$$\begin{center}
$s_1$ \\
$s_2$ $s_3$ \\
$s_4$ $s_5$ $s_6$ \\
$s_7$ $s_8$ $s_9$ $s_{10}$ \\
...
\end{center}$$

Sub-triangles can start at any element of the array and extend down as far as we like (taking-in the two elements directly below it from the next row, the three elements directly below from the row after that, and so on).  
The "sum of a sub-triangle" is defined as the sum of all the elements it contains.  
Find the smallest possible sub-triangle sum.  

### Answer:
```
-271248680
```


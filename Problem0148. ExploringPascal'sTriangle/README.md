# Exploring Pascal's Triangle
Problem 148
### Description:
We can easily verify that none of the entries in the first seven rows of Pascal's triangle are divisible by 7:
\[
\begin{array}{c}
1 \\
1 \quad 1 \\
1 \quad 2 \quad 1 \\
1 \quad 3 \quad 3 \quad 1 \\
1 \quad 4 \quad 6 \quad 4 \quad 1 \\
1 \quad 5 \quad 10 \quad 10 \quad 5 \quad 1 \\
1 \quad 6 \quad 15 \quad 20 \quad 15 \quad 6 \quad 1
\end{array}
\]  
However, if we check the first one hundred rows, we will find that only 2361 of the 5050 entries are not divisible by 7.

Find the number of entries which are not divisible by 7 in the first one billion ($10^9$) rows of Pascal's triangle.

### Answer:
```
2129970655314432
```
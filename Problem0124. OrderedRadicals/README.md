# Ordered Radicals
Problem 124
### Description:
The radical of $n$, $\text{rad}(n)$, is the product of the distinct prime factors of $n$. For example, $504 = 2^3 \times 3^2 \times 7$, so $\text{rad}(504) = 2 \times 3 \times 7 = 42$.

If we calculate $\text{rad}(n)$ for $1 \le n \le 10$, then sort them on $\text{rad}(n)$, and sorting on $n$ if the radical values are equal, we get:

\begin{tabular}{|c|c|c|c|c|}
\hline
\multicolumn{2}{|c|}{\textbf{Unsorted}} & \multicolumn{3}{|c|}{\textbf{Sorted}} \\
\hline
$n$ & $\text{rad}(n)$ & $n$ & $\text{rad}(n)$ & $k$ \\
\hline
1 & 1 & 1 & 1 & 1 \\
\hline
2 & 2 & 2 & 2 & 2 \\
\hline
3 & 3 & 4 & 2 & 3 \\
\hline
4 & 2 & 8 & 2 & 4 \\
\hline
5 & 5 & 3 & 3 & 5 \\
\hline
6 & 6 & 9 & 3 & 6 \\
\hline
7 & 7 & 5 & 5 & 7 \\
\hline
8 & 2 & 6 & 6 & 8 \\
\hline
9 & 3 & 7 & 7 & 9 \\
\hline
10 & 10 & 10 & 10 & 10 \\
\hline
\end{tabular}

Let $E(k)$ be the $k$-th element in the sorted $n$ column; for example, $E(4) = 8$ and $E(6) = 9$.

If $\text{rad}(n)$ is sorted for $1 \le n \le 100000$, find $E(10000)$.

### Answer:
```
21417
```

# Diophantine Equation
Problem 66
### Description:
Consider quadratic Diophantine equations of the form:
$$x^2 - Dy^2 = 1$$
For example, when D = 13
, the minimal solution in x
 is 649<sup>2</sup> - 13 x 180<sup>2</sup> = 1
.

It can be assumed that there are no solutions in positive integers when D
 is square.

By finding minimal solutions in x
 for D = {2, 3, 5, 6, 7}
, we obtain the following:
$$\begin{align}
3^2 - 2 \times 2^2 &= 1\\
2^2 - 3 \times 1^2 &= 1\\
{\color{red}{\mathbf 9}}^2 - 5 \times 4^2 &= 1\\
5^2 - 6 \times 2^2 &= 1\\
8^2 - 7 \times 3^2 &= 1
\end{align}$$
 
Hence, by considering minimal solutions in x
 for D <= 7
, the largest x
 is obtained when D = 5
.

Find the value of D <= 1000
 in minimal solutions of x
 for which the largest value of x
 is obtained.

### Answer:
```
661
```

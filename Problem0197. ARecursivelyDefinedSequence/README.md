# A Recursively Defined Sequence
Problem 197
### Description:
Given the function:
$$f(x) = \lfloor 2^{30.403243784 - x^2} \rfloor \times 10^{-9}$$
where $\lfloor \cdot \rfloor$ denotes the floor function.

The sequence $u_n$ is defined by:
* $u_0 = -1$
* $u_{n+1} = f(u_n)$

Find $u_n + u_{n+1}$ for $n = 10^{12}$. 

*Note: Provide the answer with 9 digits after the decimal point.*

### Answer:
```
1.710637717
```
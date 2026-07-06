# Subsets with a Unique Sum
Problem 201
### Description:
For any set $A$ of numbers, let $\text{sum}(A)$ be the sum of the elements of $A$. 

Consider the set $B = \{1, 3, 6, 8, 10, 11\}$. There are 20 subsets of $B$ containing three elements, and their sums are:

* $\text{sum}(\{1, 3, 6\}) = 10$
* $\text{sum}(\{1, 3, 8\}) = 12$
* $\text{sum}(\{1, 3, 10\}) = 14$
* $\text{sum}(\{1, 3, 11\}) = 15$
* $\text{sum}(\{1, 6, 8\}) = 15$
* $\text{sum}(\{1, 6, 10\}) = 17$
* $\text{sum}(\{1, 6, 11\}) = 18$
* $\text{sum}(\{1, 8, 10\}) = 19$
* $\text{sum}(\{1, 8, 11\}) = 20$
* $\text{sum}(\{1, 10, 11\}) = 22$
* $\text{sum}(\{3, 6, 8\}) = 17$
* $\text{sum}(\{3, 6, 10\}) = 19$
* $\text{sum}(\{3, 6, 11\}) = 20$
* $\text{sum}(\{3, 8, 10\}) = 21$
* $\text{sum}(\{3, 8, 11\}) = 22$
* $\text{sum}(\{3, 10, 11\}) = 24$
* $\text{sum}(\{6, 8, 10\}) = 24$
* $\text{sum}(\{6, 8, 11\}) = 25$
* $\text{sum}(\{6, 10, 11\}) = 27$
* $\text{sum}(\{8, 10, 11\}) = 29$

Some of these sums occur more than once, while others are unique. For a set $A$, let $U(A, k)$ be the set of **unique sums** of $k$-element subsets of $A$. In our example, we find $U(B, 3) = \{10, 12, 14, 18, 21, 25, 27, 29\}$ and $\text{sum}(U(B, 3)) = 156$.

Now consider the 100-element set $S = \{1^2, 2^2, \dots, 100^2\}$.
$S$ has $100891344545564193334812497256$ 50-element subsets.

**Determine the sum of all integers which are the sum of exactly one of the 50-element subsets of $S$.** *In other words, find $\text{sum}(U(S, 50))$.*

### Answer:
```
115039000
```
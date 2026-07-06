# Ambiguous Numbers
Problem 198
### Description:
A **best approximation** to a real number $x$ for the denominator bound $d$ is a rational number $\frac{r}{s}$ (in reduced form) with $s \leq d$, such that any rational number $\frac{p}{q}$ which is closer to $x$ than $\frac{r}{s}$ has $q > d$.

Usually, the best approximation to a real number is uniquely determined for all denominator bounds. However, there are exceptions; for example, $\frac{9}{40}$ has two best approximations, $\frac{1}{4}$ and $\frac{1}{5}$, for the denominator bound $6$. 

We call a real number $x$ **ambiguous** if there is at least one denominator bound for which $x$ possesses two best approximations. Clearly, an ambiguous number is necessarily rational.

How many ambiguous numbers $x = \frac{p}{q}$, where $0 < x < \frac{1}{100}$, are there whose denominator $q$ does not exceed $10^8$?

### Answer:
```
52374425
```
# Special Subset Sums: Meta-testing
Problem 106
### Description:
<p>Let $S(A)$ represent the sum of elements in set $A$ of size $n$. We shall call it a special sum set if for any two non-empty disjoint subsets, $B$ and $C$, the following properties are true:</p>
<ol><li>$S(B) \ne S(C)$; that is, sums of subsets cannot be equal.</li>
<li>If $B$ contains more elements than $C$ then $S(B) \gt S(C)$.</li>
</ol><p>For this problem we shall assume that a given set contains $n$ strictly increasing elements and it already satisfies the second rule.</p>
<p>Surprisingly, out of the $25$ possible subset pairs that can be obtained from a set for which $n = 4$, only $1$ of these pairs need to be tested for equality (first rule). Similarly, when $n = 7$, only $70$ out of the $966$ subset pairs need to be tested.</p>
<p>For $n = 12$, how many of the $261625$ subset pairs that can be obtained need to be tested for equality?</p>

### Answer:
```
21384
```

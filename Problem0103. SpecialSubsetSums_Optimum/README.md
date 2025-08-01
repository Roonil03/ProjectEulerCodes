# Special Subset Sums: Optimum
Problem 103
### Description:
<p>Let $$S(A)$$ represent the sum of elements in set $$A$$ of size $$n$$. We shall call it a special sum set if for any two non-empty disjoint subsets, $$B$$ and $$C$$, the following properties are true:</p>
<ol><li>$$S(B) \ne S(C)$$; that is, sums of subsets cannot be equal.</li>
<li>If $$B$$ contains more elements than $$C$$ then $$S(B) \gt S(C)$$.</li>
</ol><p>If $$S(A)$$ is minimised for a given $$n$$, we shall call it an optimum special sum set. The first five optimum special sum sets are given below.</p>
<ul style="list-style-type:none;">
<li>$$n = 1$$: $$\{1\}$$</li>
<li>$$n = 2$$: $$\{1, 2\}$$</li>
<li>$$n = 3$$: $$\{2, 3, 4\}$$</li>
<li>$$n = 4$$: $$\{3, 5, 6, 7\}$$</li>
<li>$$n = 5$$: $$\{6, 9, 11, 12, 13\}$$</li></ul>
<p>It <i>seems</i> that for a given optimum set, $$A = \{a_1, a_2, \dots, a_n\}$$, the next optimum set is of the form $$B = \{b, a_1 + b, a_2 + b, \dots, a_n + b\}$$, where $$b$$ is the "middle" element on the previous row.</p>
<p>By applying this "rule" we would expect the optimum set for $$n = 6$$ to be $$A = \{11, 17, 20, 22, 23, 24\}$$, with $$S(A) = 117$$. However, this is not the optimum set, as we have merely applied an algorithm to provide a near optimum set. The optimum set for $$n = 6$$ is $$A = \{11, 18, 19, 20, 22, 25\}$$, with $$S(A) = 115$$ and corresponding set string: 111819202225.</p>
<p>Given that $$A$$ is an optimum special sum set for $$n = 7$$, find its set string.</p>

### Answer:
```
20313839404245
```

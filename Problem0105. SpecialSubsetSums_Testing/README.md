# Special Subset Sums: Testing
Problem 105
### Description:
<p>Let $$S(A)$$ represent the sum of elements in set $$A$$ of size $$n$$. We shall call it a special sum set if for any two non-empty disjoint subsets, $$B$$ and $$C$$, the following properties are true:</p>
<ol><li>$$S(B) \ne S(C)$$; that is, sums of subsets cannot be equal.</li>
<li>If $$B$$ contains more elements than $$C$$ then $$S(B) \gt S(C)$$.</li>
</ol><p>For example, $$\{81, 88, 75, 42, 87, 84, 86, 65\}$$ is not a special sum set because $$65 + 87 + 88 = 75 + 81 + 84$$, whereas $$\{157, 150, 164, 119, 79, 159, 161, 139, 158\}$$ satisfies both rules for all possible subset pair combinations and $$S(A) = 1286$$.</p>
<p>Using sets.txt (right click and "Save Link/Target As..."), a 4K text file with one-hundred sets containing seven to twelve elements (the two examples given above are the first two sets in the file), identify all the special sum sets, $$A_1, A_2, \dots, A_k$$, and find the value of $$S(A_1) + S(A_2) + \cdots + S(A_k)$$.</p>

### Answer:
```
73702
```

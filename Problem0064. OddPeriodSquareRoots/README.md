# Odd Period Square Roots
Problem 64
### Description:
All square roots are periodic when written as continued fractions and can be written in the form:

$$\sqrt{N} = a_0 + \cfrac{1}{a_1 + \cfrac{1}{a_2 + \cfrac{1}{a_3 + \cdots}}}$$

 
 
For example, let us consider 

$$\quad \quad \sqrt{23}=4+\sqrt{23}-4=4+\frac 1 {\frac 1 {\sqrt{23}-4}}=4+\frac 1  {1+\frac{\sqrt{23}-3}7}$$
 
 
 
If we continue we would get the following expansion:

$$\displaystyle \quad \quad \sqrt{23}=4+\frac 1 {1+\frac 1 {3+ \frac 1 {1+\frac 1 {8+ \dots}}}}$$
 
 
The process can be summarised as follows:
$$a_0 = 4, \quad \frac{1}{\sqrt{23}-4} = \frac{\sqrt{23}+4}{7} = 1 + \frac{\sqrt{23}-3}{7}$$

$$a_1 = 1, \quad \frac{7}{\sqrt{23}-3} = \frac{7(\sqrt{23}+3)}{14} = 3 + \frac{\sqrt{23}-3}{2}$$

$$a_2 = 3, \quad \frac{2}{\sqrt{23}-3} = \frac{2(\sqrt{23}+3)}{14} = 1 + \frac{\sqrt{23}-4}{7}$$

$$a_3 = 1, \quad \frac{7}{\sqrt{23}-4} = \frac{7(\sqrt{23}+4)}{7} = 8 + \sqrt{23} - 4$$

$$a_4 = 8, \quad \frac{1}{\sqrt{23}-4} = \frac{\sqrt{23}+4}{7} = 1 + \frac{\sqrt{23}-3}{7}$$

$$a_5 = 1, \quad \frac{7}{\sqrt{23}-3} = \frac{7(\sqrt{23}+3)}{14} = 3 + \frac{\sqrt{23}-3}{2}$$

$$a_6 = 3, \quad \frac{2}{\sqrt{23}-3} = \frac{2(\sqrt{23}+3)}{14} = 1 + \frac{\sqrt{23}-4}{7}$$

$$a_7 = 1, \quad \frac{7}{\sqrt{23}-4} = \frac{7(\sqrt{23}+4)}{7} = 8 + \sqrt{23} - 4$$


It can be seen that the sequence is repeating. For conciseness, we use the notation √23 = `[4;(1,3,1,8)]`
, to indicate that the block (1,3,1,8) repeats indefinitely.

The first ten continued fraction representations of (irrational) square roots are:
$$\sqrt{2} = [1; (2)], \text{ period} = 1$$

$$\sqrt{3} = [1; (1, 2)], \text{ period} = 2$$

$$\sqrt{5} = [2; (4)], \text{ period} = 1$$

$$\sqrt{6} = [2; (2, 4)], \text{ period} = 2$$

$$\sqrt{7} = [2; (1, 1, 1, 4)], \text{ period} = 4$$

$$\sqrt{8} = [2; (1, 4)], \text{ period} = 2$$

$$\sqrt{10} = [3; (6)], \text{ period} = 1$$

$$\sqrt{11} = [3; (3, 6)], \text{ period} = 2$$

$$\sqrt{12} = [3; (2, 6)], \text{ period} = 2$$

$$\sqrt{13} = [3; (1, 1, 1, 1, 6)], \text{ period} = 5$$


Exactly four continued fractions, for N <= 13
, have an odd period.

How many continued fractions for N <= 10,000
 have an odd period?
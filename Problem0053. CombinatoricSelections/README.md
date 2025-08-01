# Combinatoric Selections
Problem 53
### Description:
There are exactly ten ways of selecting three from five, 12345:

123, 124, 125, 134, 135, 145, 234, 235, 245, and 345

In combinatorics, we use the notation, <sup>5</sup>C</sub>3</sub> = 10
.

In general, <sup>n</sup>C<sub>r</sub> = n! / (r! (n-r)!)
 
, where r <= n
, n! = n x (n-1) x ... x 3 x 2 x 1
, and 0! = 1
.

It is not until n = 23
, that a value exceeds one-million: <sup>23</sup>C<sub>10</sub> = 1144066
.

How many, not necessarily distinct, values of <sup>n</sup>C<sub>r</sub>
 for 1 <= n <= 100
, are greater than one-million?

### Answer:
```
4075
```
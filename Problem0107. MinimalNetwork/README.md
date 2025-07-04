# Minimal Network
Problem 107
### Description:
The following undirected network consists of seven vertices and twelve edges with a total weight of 243.  
<img src="https://projecteuler.net/resources/images/0107_1.png?1678992052"><br>
The same network can be represented by the matrix below.  
|   | A  | B  | C  | D  | E  | F  | G  |
|---|----|----|----|----|----|----|----|
| A | -  | 16 | 12 | 21 | -  | -  | -  |
| B | 16 | -  | 17 | 20 | -  | -  | -  |
| C | 12 | -  | -  | 28 | -  | 31 | -  |
| D | 21 | 17 | 28 | -  | 18 | 19 | 23 |
| E | -  | 20 | -  | 18 | -  | -  | 11 |
| F | -  | -  | 31 | 19 | -  | -  | 27 |
| G | -  | -  | -  | 23 | 11 | 27 | -  |

However, it is possible to optimise the network by removing some edges and still ensure that all points on the network remain connected. The network which achieves the maximum saving is shown below. It has a weight of 93, representing a saving of 243 − 93 = 150 from the original network.

<img src="https://projecteuler.net/resources/images/0107_2.png?1678992052"><br>

Using network.txt (right click and 'Save Link/Target As...'), a 6K text file containing a network with forty vertices, and given in matrix form, find the maximum saving which can be achieved by removing redundant edges whilst ensuring that the network remains connected.

### Answer:
```
259679
```
Given a list of integers, and a set of conditions (odd, even, greater than 5, multiple of 3, prime, and many more such custom conditions that may be dynamically defined by user), write a program to return only the integers from the given list that match ANY of the conditions.

```
Sample Input:
A list containing 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20
Conditions specified using a set of functions: prime, greater than 15, multiple of 5
Sample Output: 2, 3, 5, 7, 10, 11, 13, 15, 16, 17, 18, 19, 20

Sample Input:
A list containing 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20
Conditions specified using a set of functions: less than 6, multiple of 3
Sample Output: 1, 2, 3, 4, 5, 6, 9, 12, 15, 18
```